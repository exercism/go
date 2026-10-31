package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

// NewNode constructs a new Node with the given value & no next/prev links.
func NewList[T comparable](elements ...T) *Node[T] {
	panic("Please implement the NewList function")
}

// Next returns a pointer to the next node in the linked list.
func (n *Node[T]) Next() *Node[T] {
	panic("Please implement the Next function")
}

// Prev returns a pointer to the previous node in the linked list.
func (n *Node[T]) Prev() *Node[T] {
	panic("Please implement the Prev function")
}

// Unshift pushes a new value before Head.
func (ll *List[T]) Unshift(v T) {
}

// Push pushes a new value after Tail.
func (ll *List[T]) Push(v T) {
	panic("Please implement the Push function")
}

// Shift posp the element at Head. It returns error if the linked list is empty.
func (ll *List[T]) Shift() (T, error) {
	panic("Please implement the Shift function")
}

// Pop pops the element at Tail. It returns error if the linked list is empty.
func (ll *List[T]) Pop() (T, error) {
	panic("Please implement the Pop function")
}

// Reverse reverses the given linked list in-place.
func (ll *List[T]) Reverse() {
	panic("Please implement the Reverse function")
}

// First returns a pointer to the first node in the linked list (head).
func (n *Node[T]) First() *Node[T] {
	panic("Please implement the First function")
}

// Last returns a pointer to the last node in the linked list (tail).
func (n *Node[T]) Last() *Node[T] {
	panic("Please implement the Last function")
}

// Count returns the number of nodes in the list.
func (ll *List[T]) Count() int {
	panic("Please implement the Count function")
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List[T]) Delete(v T) bool {
	panic("Please implement the Delete function")
}
