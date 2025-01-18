package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...interface{}) *List {
	panic("Please implement the NewList function")
}

func (n *Node) Next() *Node {
	panic("Please implement the Next function")
}

func (n *Node) Prev() *Node {
	panic("Please implement the Prev function")
}

func (l *List) Unshift(v interface{}) {
	panic("Please implement the Unshift function")
}

func (l *List) Push(v interface{}) {
	panic("Please implement the Push function")
}

func (l *List) Shift() (interface{}, error) {
	panic("Please implement the Shift function")
}

func (l *List) Pop() (interface{}, error) {
	panic("Please implement the Pop function")
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	panic("Please implement the First function")
}

func (l *List) Last() *Node {
	panic("Please implement the Last function")
}
