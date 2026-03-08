package satellite

import (
	"fmt"
	"strings"
	"testing"
)

func TreesEqual(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Value == b.Value && TreesEqual(a.Right, b.Right) && TreesEqual(a.Left, b.Left)
}

func FmtTree(n *Node) string {
	if n == nil {
		return "{}"
	}
	return fmt.Sprintf("Node{V: %q, L: %s, R: %s}", n.Value, FmtTree(n.Left), FmtTree(n.Right))
}

func TestTreeFromTraversals(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual, err := TreeFromTraversals(tc.preorder, tc.inorder); tc.err == "" && err != nil {
				t.Fatalf("TreeFromTraversals(%#v, %#v) has unexpected error, %v", tc.preorder, tc.inorder, err)
			} else if tc.err != "" && err == nil {
				t.Fatalf("TreeFromTraversals(%#v, %#v) did not error; want error, %v", tc.preorder, tc.inorder, tc.err)
			} else if tc.err != "" && !strings.HasPrefix(err.Error(), tc.err) {
				t.Fatalf("TreeFromTraversals(%#v, %#v) has error %q; want error with prefix %q", tc.preorder, tc.inorder, err, tc.err)
			} else if !TreesEqual(actual, tc.expected) {
				t.Fatalf("TreeFromTraversals(%#v, %#v)\ngot: %s\nwant: %s", tc.preorder, tc.inorder, FmtTree(actual), FmtTree(tc.expected))
			}
		})
	}
}

func BenchmarkValid(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			TreeFromTraversals(tc.preorder, tc.inorder)
		}
	}
}
