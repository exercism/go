package satellite

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4e18617 More complex satellite tests (#2612)

var testCases = []struct {
	description string
	preorder    []string
	inorder     []string
	expected    *Node
	err         string
}{
	{
		description: "Empty tree",
		preorder:    []string{},
		inorder:     []string{},
		expected:    nil,
	},
	{
		description: "Tree with one item",
		preorder:    []string{"a"},
		inorder:     []string{"a"},
		expected:    &Node{Value: "a", Left: nil, Right: nil},
	},
	{
		description: "Tree with many items",
		preorder:    []string{"a", "i", "x", "f", "r"},
		inorder:     []string{"i", "a", "f", "x", "r"},
		expected:    &Node{Value: "a", Left: &Node{Value: "i", Left: nil, Right: nil}, Right: &Node{Value: "x", Left: &Node{Value: "f", Left: nil, Right: nil}, Right: &Node{Value: "r", Left: nil, Right: nil}}},
	},
	{
		description: "Reject traversals of different length",
		preorder:    []string{"a", "b"},
		inorder:     []string{"b", "a", "r"},
		err:         "traversals must have the same length",
	},
	{
		description: "Reject inconsistent traversals of same length",
		preorder:    []string{"x", "y", "z"},
		inorder:     []string{"a", "b", "c"},
		err:         "traversals must have the same elements",
	},
	{
		description: "Reject traversals with repeated items",
		preorder:    []string{"a", "b", "a"},
		inorder:     []string{"b", "a", "a"},
		err:         "traversals must contain unique items",
	},
	{
		description: "A degenerate binary tree",
		preorder:    []string{"a", "b", "c", "d"},
		inorder:     []string{"d", "c", "b", "a"},
		expected:    &Node{Value: "a", Left: &Node{Value: "b", Left: &Node{Value: "c", Left: &Node{Value: "d", Left: nil, Right: nil}, Right: nil}, Right: nil}, Right: nil},
	},
	{
		description: "Another degenerate binary tree",
		preorder:    []string{"a", "b", "c", "d"},
		inorder:     []string{"a", "b", "c", "d"},
		expected:    &Node{Value: "a", Left: nil, Right: &Node{Value: "b", Left: nil, Right: &Node{Value: "c", Left: nil, Right: &Node{Value: "d", Left: nil, Right: nil}}}},
	},
	{
		description: "Tree with many more items",
		preorder:    []string{"a", "b", "d", "g", "h", "c", "e", "f", "i"},
		inorder:     []string{"g", "d", "h", "b", "a", "e", "c", "i", "f"},
		expected:    &Node{Value: "a", Left: &Node{Value: "b", Left: &Node{Value: "d", Left: &Node{Value: "g", Left: nil, Right: nil}, Right: &Node{Value: "h", Left: nil, Right: nil}}, Right: nil}, Right: &Node{Value: "c", Left: &Node{Value: "e", Left: nil, Right: nil}, Right: &Node{Value: "f", Left: &Node{Value: "i", Left: nil, Right: nil}, Right: nil}}},
	},
}
