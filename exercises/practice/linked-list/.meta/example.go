package linkedlist

import "errors"

var ErrEmptyList = errors.New("list is empty")


// Node is a node in a linked list.
type Node[T comparable] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
}

// NewNode constructs a new Node with the given value & no next/prev links.
func NewNode[T comparable](v T) *Node[T] {
	return &Node[T]{
		Value: v,
		next:  nil,
		prev:  nil,
	}
}

// Next returns a pointer to the next node in the linked list.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Prev returns a pointer to the previous node in the linked list.
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// First returns a pointer to the first node in the linked list (head).
func (n *Node[T]) First() *Node[T] {
	cur := n
	for ; cur.prev != nil; cur = cur.prev {
	}

	return cur
}

// Last returns a pointer to the last node in the linked list (tail).
func (n *Node[T]) Last() *Node[T] {
	cur := n
	for ; cur.next != nil; cur = cur.next {
	}

	return cur
}

// List is a doubly-linked list with Head and Tail.
type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

// NewList constructs a doubly linked list from a sequence of integers.
func NewList[T comparable](elements ...T) *List[T] {
	ll := &List[T]{
		head: nil,
		tail: nil,
	}

	if len(elements) < 1 {
		return ll
	}

	ll.head = NewNode(elements[0])
	ll.tail = ll.head

	if len(elements) == 1 {
		return ll
	}

	cur := ll.head
	for i := 1; i < len(elements); i++ {
		cur.next = NewNode(elements[i])
		cur.next.prev = cur
		cur = cur.next
	}

	ll.tail = cur

	return ll
}

// First returns a pointer to the first node in the linked list (head).
func (ll *List[T]) First() *Node[T] {
	return ll.head
}

// Last returns a pointer to the last node in the linked list (tail).
func (ll *List[T]) Last() *Node[T] {
	return ll.tail
}

// Reverse reverses the given linked list in-place.
func (ll *List[T]) Reverse() {
	if ll.head == nil || ll.head.next == nil {
		return
	}

	node := ll.tail
	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.next
	}
	ll.head, ll.tail = ll.tail, ll.head
}

// Unshift pushes a new value before Head.
func (ll *List[T]) Unshift(v T) {
	n := NewNode(v)

	switch {
	default:
		panic("bad PushFront implementation")
	case ll.head == nil && ll.tail == nil: // empty list
		ll.head = n
		ll.tail = n
	case ll.head != nil && ll.tail != nil: // non-empty list
		n.next = ll.head
		ll.head.prev = n

		ll.head = n
	}
}

// Push pushes a new value after Tail.
func (ll *List[T]) Push(v T) {
	n := NewNode(v)

	switch {
	default:
		panic("bad PushBack implementation")
	case ll.head == nil && ll.tail == nil: // empty list
		ll.head = n
		ll.tail = n
	case ll.head != nil && ll.tail != nil: // non-empty list
		ll.tail.next = n
		n.prev = ll.tail

		ll.tail = n
	}
}

// Shift posp the element at Head. It returns error if the linked list is empty.
func (ll *List[T]) Shift() (T, error) {
	switch {
	default:
		panic("bad PopFront implementation")
	case ll.head == nil && ll.tail == nil: // empty list
		var zero T
		return zero, ErrEmptyList
	case ll.head != nil && ll.tail != nil && ll.head.next == nil: // 1 element
		v := ll.head.Value
		ll.head = nil
		ll.tail = nil

		return v, nil
	case ll.head != nil && ll.tail != nil && ll.head.next != nil: // >1 element
		v := ll.head.Value
		ll.head.next.prev = nil
		ll.head = ll.head.next

		return v, nil
	}
}

// Pop pops the element at Tail. It returns error if the linked list is empty.
func (ll *List[T]) Pop() (T, error) {
	switch {
	default:
		panic("bad PopBack implementation")
	case ll.head == nil && ll.tail == nil: // empty list
		var zero T
		return zero, ErrEmptyList
	case ll.head != nil && ll.tail != nil && ll.tail.prev == nil: // 1 element
		v := ll.tail.Value
		ll.head = nil
		ll.tail = nil

		return v, nil
	case ll.head != nil && ll.tail != nil && ll.tail.prev != nil: // >1 element
		v := ll.tail.Value
		ll.tail.prev.next = nil
		ll.tail = ll.tail.prev

		return v, nil
	}
}

// Count returns the number of nodes in the list.
func (ll *List[T]) Count() int {
	var count int
	for cur := ll.head; cur != nil; cur = cur.Next() {
		count++
	}
	return count
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List[T]) Delete(v T) bool {
	for cur := ll.head; cur != nil; cur = cur.Next() {
		if cur.Value == v {
			if cur.prev == nil {
				ll.head = cur.next
			} else {
				cur.prev.next = cur.next
			}
			if cur.next == nil {
				ll.tail = cur.prev
			} else {
				cur.next.prev = cur.prev
			}
			return true
		}
	}
	return false
}
