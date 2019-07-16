package linkedlist

import (
	"errors"
)

// Node is a node in a linked list.
type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

// NewNode constructs a new Node with the given value & no next/prev links.
func NewNode(v interface{}) *Node {
	return &Node{
		Val:  v,
		next: nil,
		prev: nil,
	}
}

// Next returns a pointer to the next node in the linked list.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns a pointer to the previous node in the linked list.
func (n *Node) Prev() *Node {
	return n.prev
}

// First returns a pointer to the first node in the linked list (head).
func (n *Node) First() *Node {
	cur := n
	for ; cur.prev != nil; cur = cur.prev {
	}

	return cur
}

// Last returns a pointer to the last node in the linked list (tail).
func (n *Node) Last() *Node {
	cur := n
	for ; cur.next != nil; cur = cur.next {
	}

	return cur
}

// List is a doubly-linked list with Head and Tail.
type List struct {
	head *Node
	tail *Node
}

// NewList constructs a doubly linked list from a sequence of integers.
func NewList(vs ...interface{}) *List {
	ll := &List{
		head: nil,
		tail: nil,
	}

	if len(vs) < 1 {
		return ll
	}

	ll.head = NewNode(vs[0])
	ll.tail = ll.head

	if len(vs) == 1 {
		return ll
	}

	cur := ll.head
	for i := 1; i < len(vs); i++ {
		cur.next = NewNode(vs[i])
		cur.next.prev = cur
		cur = cur.next
	}

	ll.tail = cur

	return ll
}

// First returns a pointer to the first node in the linked list (head).
func (ll *List) First() *Node {
	return ll.head
}

// Last returns a pointer to the last node in the linked list (tail).
func (ll *List) Last() *Node {
	return ll.tail
}

// Reverse reverses the given linked list in-place.
func (ll *List) Reverse() {
	if ll.head == nil || ll.head.next == nil {
		return
	}

	// construct singly-linked list from the back
	dummy := NewNode(-1)
	cur := dummy
	n := ll.tail
	for n != nil {
		cur.next = n

		cur = cur.next
		n = n.prev
	}
	cur.next = nil // cur will be the new ll.Tail -> set .next = nil

	// add prev -> doubly-linked list
	prev := dummy.next
	n = dummy.next.next
	for n != nil {
		n.prev = prev

		n = n.next
		prev = prev.next
	}

	// update Head & Tail
	ll.head, ll.tail = ll.tail, ll.head
	ll.head.prev = nil
}

// PushFront pushes a new value before Head.
func (ll *List) PushFront(v interface{}) {
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

// PushBack pushes a new value after Tail.
func (ll *List) PushBack(v interface{}) {
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

var (
	ErrEmptyList = errors.New("list is empty")
)

// PopFront posp the element at Head. It returns error if the linked list is empty.
func (ll *List) PopFront() (interface{}, error) {
	switch {
	default:
		panic("bad PopFront implementation")
	case ll.head == nil && ll.tail == nil: // empty list
		return nil, ErrEmptyList
	case ll.head != nil && ll.tail != nil && ll.head.next == nil: // 1 element
		v := ll.head.Val
		ll.head = nil
		ll.tail = nil

		return v, nil
	case ll.head != nil && ll.tail != nil && ll.head.next != nil: // >1 element
		v := ll.head.Val
		ll.head.next.prev = nil
		ll.head = ll.head.next

		return v, nil
	}
}

// PopBack pops the element at Tail. It returns error if the linked list is empty.
func (ll *List) PopBack() (interface{}, error) {
	switch {
	default:
		panic("bad PopBack implementation")
	case ll.head == nil && ll.tail == nil: // empty list
		return nil, ErrEmptyList
	case ll.head != nil && ll.tail != nil && ll.tail.prev == nil: // 1 element
		v := ll.tail.Val
		ll.head = nil
		ll.tail = nil

		return v, nil
	case ll.head != nil && ll.tail != nil && ll.tail.prev != nil: // >1 element
		v := ll.tail.Val
		ll.tail.prev.next = nil
		ll.tail = ll.tail.prev

		return v, nil
	}
}
