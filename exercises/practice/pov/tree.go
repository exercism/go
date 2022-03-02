package pov

// Define the Tree type here

// New creates and returns a new Tree with the given root value and children
// You can assume that there will be no duplicate values in test trees.
func New(value string, children ...*Tree) *Tree {
	panic("Please implement this function")
}

// Value returns the value at the root of a tree
func (tr *Tree) Value() string {
	panic("Please implement this function")
}

// Children returns a slice containing the children of a tree
// There is no need to sort the elements in the result slice,
// they can be in any order
func (tr *Tree) Children() []*Tree {
	panic("Please implement this function")
}

// String describes a tree in a compact S-expression format
// This helps to make test outputs more readable
// Feel free to adapt this method as you see fit
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}
