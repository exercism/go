package zipper

type Node struct {
	value int
	left  *Node
	right *Node
}

type Zipper struct {
	// Add fields here.
}

func NewZipper(tree *Node) Zipper {
	panic("Please implement the NewZipper() function")
}

func (z Zipper) Value() int {
	panic("Please implement the Value() function")
}

func (z Zipper) ToTree() *Node {
	panic("Please implement the ToTree() function")
}

func (z Zipper) Left() (Zipper, error) {
	panic("Please implement the Left() function")
}

func (z Zipper) Right() (Zipper, error) {
	panic("Please implement the Right() function")
}

func (z Zipper) Up() (Zipper, error) {
	panic("Please implement the Up() function")
}

func (z Zipper) SetValue(v int) Zipper {
	panic("Please implement the SetValue() function")
}

func (z Zipper) SetLeft(v *Node) Zipper {
	panic("Please implement the SetLeft() function")
}

func (z Zipper) SetRight(v *Node) Zipper {
	panic("Please implement the SetRight() function")
}
