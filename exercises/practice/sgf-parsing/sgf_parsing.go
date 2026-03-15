package sgfparsing

// Node represents an SGF node with properties and child nodes.
type Node struct {
	Properties map[string][]string
	Children   []*Node
}

// Parse decodes an SGF string and returns the root node of the tree.
func Parse(encoded string) (*Node, error) {
	panic("Please implement the Parse function")
}
