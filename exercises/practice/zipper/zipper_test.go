package zipper

import (
	"fmt"
	"strings"
	"testing"
)

func TreeEqual(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.value == b.value && TreeEqual(a.left, b.left) && TreeEqual(a.right, b.right)
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("Node{value: %d, left: %s, right: %s}", n.value, n.left, n.right)
}

func runTest(tc testCase) error {
	var callStack []string

	z := NewZipper(tc.input)
	callStack = append(callStack, fmt.Sprintf("z := NewZipper(%s)", tc.input))
	var err error
	for i, o := range tc.operations {
		switch o.funcCall {
		case "value":
			if tc.want.wantType != "int" {
				panic("Getting a value but not checking for an int")
			}
			got := z.Value()
			callStack = append(callStack, "z.Value()")
			if got != tc.want.intVal {
				return fmt.Errorf("Calls:\n%s\n\ngot: %d, want %d", strings.Join(callStack, "\n"), got, tc.want.intVal)
			}
		case "to_tree":
			if tc.want.wantType != "tree" {
				panic("Getting a tree but not checking for an tree")
			}
			got := z.ToTree()
			callStack = append(callStack, "z.ToTree()")
			if !TreeEqual(got, tc.want.nodeVal) {
				return fmt.Errorf("Calls:\n%s\n\ngot: %s\nwant: %s", strings.Join(callStack, "\n"), got, tc.want.nodeVal)
			}
		case "up":
			z, err = z.Up()
			callStack = append(callStack, "z = z.Up()")
		case "left":
			z, err = z.Left()
			callStack = append(callStack, "z = z.Left()")
		case "right":
			z, err = z.Right()
			callStack = append(callStack, "z = z.Right()")
		case "set_value":
			z = z.SetValue(o.intArg)
			callStack = append(callStack, fmt.Sprintf("z = z.SetValue(%d)", o.intArg))
		case "set_left":
			z = z.SetLeft(o.nodeArg)
			callStack = append(callStack, fmt.Sprintf("z = z.SetLeft(%s)", o.nodeArg))
		case "set_right":
			z = z.SetRight(o.nodeArg)
			callStack = append(callStack, fmt.Sprintf("z = z.SetRight(%s)", o.nodeArg))
		default:
			panic("Unexpected operation")
		}
		if err != nil && (i != len(tc.operations)-1 || tc.want.wantType != "zipper") {
			return fmt.Errorf("Calls:\n%s\n\nreturned an unexpected error, %v", strings.Join(callStack, "\n"), err)
		}
	}
	if err == nil && tc.want.wantType == "zipper" {
		return fmt.Errorf("Calls:\n%s\n\nwas successful, want an error", strings.Join(callStack, "\n"))
	}
	return nil
}

func TestExpectedValue(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if err := runTest(tc); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}

func TestImmutable(t *testing.T) {
	t.Run("SetValue does not change the original tree", func(t *testing.T) {
		root := &Node{value: 5}
		newRoot := NewZipper(root).SetValue(6).ToTree()
		if newRoot == nil || newRoot.value != 6 {
			t.Fatal("z.SetValue(6).ToTree() did not give a tree with value 6")
		}
		if root.value != 5 {
			t.Fatal("Calling z.SetValue() should not change the original tree")
		}
	})
	t.Run("SetRight does not change the original tree", func(t *testing.T) {
		root := &Node{value: 5}
		newRoot := NewZipper(root).SetRight(&Node{value: 6}).ToTree()
		if newRoot == nil || newRoot.right == nil || newRoot.right.value != 6 {
			t.Fatal("z.SetRight(&Node{value: 6}).ToTree() did not give a tree with new right node")
		}
		if root.right != nil {
			t.Fatal("Calling z.SetRight() should not change the original tree")
		}
	})
}

func TestFocusDoesNotChangeNodes(t *testing.T) {
	t.Run("Changing focus does not clone the tree", func(t *testing.T) {
		root := &Node{value: 5, left: &Node{value: 6}, right: &Node{value: 7}}
		z := NewZipper(root)
		z, _ = z.Right()
		z, _ = z.Up()
		z, _ = z.Left()
		z, _ = z.Up()
		newRoot := z.ToTree()
		if root != newRoot {
			t.Fatal("z.ToTree() should return the original tree if there were no changes to the tree.")
		}
	})
}

func BenchmarkExpectedValue(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			runTest(tc)
		}
	}
}

// Test set_* doesn't change original
