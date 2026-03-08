package satellite

import (
	"errors"
	"slices"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func TreeFromTraversals(preorder, inorder []string) (*Node, error) {
	if len(preorder) != len(inorder) {
		return nil, errors.New("traversals must have the same length")
	}
	a := slices.Clone(preorder)
	slices.Sort(a)
	b := slices.Clone(inorder)
	slices.Sort(b)
	if !slices.Equal(a, b) {
		return nil, errors.New("traversals must have the same elements")
	}
	m := make(map[string]bool)
	for _, i := range a {
		if m[i] {
			return nil, errors.New("traversals must contain unique items")
		}
		m[i] = true
	}

	if len(preorder) == 0 {
		return nil, nil
	}

	// The first element of the preorder is always the root.
	root := preorder[0]

	// The inorder can be partitioned on the root to give the left and right inorder.
	idx := slices.Index(inorder, root)
	inorderLeft, inorderRight := inorder[:idx], inorder[idx+1:]
	// With N = len(left subtree), preorder left can be found by taking the next N
	// elements after the root. The rest are preorder right.
	preorderLeft := preorder[1 : len(inorderLeft)+1]
	preorderRight := preorder[len(inorderLeft)+1:]

	left, _ := TreeFromTraversals(preorderLeft, inorderLeft)
	right, _ := TreeFromTraversals(preorderRight, inorderRight)
	return &Node{Value: root, Left: left, Right: right}, nil
}
