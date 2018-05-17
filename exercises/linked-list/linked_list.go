package linkedlist

import (
	"errors"
)

// ListNode is a node in a linked list.
type ListNode struct {
	Val  interface{}
	Next *ListNode
	Prev *ListNode
}

// NewListNode constructs a new ListNode with the given value & no next/prev links.
func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
		Prev: nil,
	}
}

// DoublyLinkedList is a doubly-linked list with Head and Tail.
type DoublyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// NewDoublyLinkedList constructs a doubly linked list from a sequence of integers.
func NewDoublyLinkedList(vs ...interface{}) *DoublyLinkedList {
	ll := &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}

	if len(vs) < 1 {
		return ll
	}

	ll.Head = NewListNode(vs[0])
	ll.Tail = ll.Head

	if len(vs) == 1 {
		return ll
	}

	cur := ll.Head
	for i := 1; i < len(vs); i++ {
		cur.Next = NewListNode(vs[i])
		cur.Next.Prev = cur
		cur = cur.Next
	}

	ll.Tail = cur

	return ll
}

// Reverse reverses the given linked list in-place.
func (ll *DoublyLinkedList) Reverse() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	// construct singly-linked list from the back
	dummy := NewListNode(-1)
	cur := dummy
	n := ll.Tail
	for n != nil {
		cur.Next = n

		cur = cur.Next
		n = n.Prev
	}
	cur.Next = nil // cur will be the new ll.Tail -> set .Next = nil

	// add prev -> doubly-linked list
	prev := dummy.Next
	n = dummy.Next.Next
	for n != nil {
		n.Prev = prev

		n = n.Next
		prev = prev.Next
	}

	// update Head & Tail
	ll.Head, ll.Tail = ll.Tail, ll.Head
	ll.Head.Prev = nil
}

// PushFront pushes a new value before Head.
func (ll *DoublyLinkedList) PushFront(v interface{}) {
	n := NewListNode(v)

	switch {
	default:
		panic("bad PushFront implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		ll.Head = n
		ll.Tail = n
	case ll.Head != nil && ll.Tail != nil: // non-empty list
		n.Next = ll.Head
		ll.Head.Prev = n

		ll.Head = n
	}
}

// PushBack pushes a new value after Tail.
func (ll *DoublyLinkedList) PushBack(v interface{}) {
	n := NewListNode(v)

	switch {
	default:
		panic("bad PushBack implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		ll.Head = n
		ll.Tail = n
	case ll.Head != nil && ll.Tail != nil: // non-empty list
		ll.Tail.Next = n
		n.Prev = ll.Tail

		ll.Tail = n
	}
}

var (
	ErrEmptyList = errors.New("list is empty")
)

// PopFront posp the element at Head. It returns error if the linked list is empty.
func (ll *DoublyLinkedList) PopFront() (interface{}, error) {
	switch {
	default:
		panic("bad PopFront implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		return 0, ErrEmptyList
	case ll.Head != nil && ll.Tail != nil && ll.Head.Next == nil: // 1 element
		v := ll.Head.Val
		ll.Head = nil
		ll.Tail = nil

		return v, nil
	case ll.Head != nil && ll.Tail != nil && ll.Head.Next != nil: // >1 element
		v := ll.Head.Val
		ll.Head.Next.Prev = nil
		ll.Head = ll.Head.Next

		return v, nil
	}
}

// PopBack pops the element at Tail. It returns error if the linked list is empty.
func (ll *DoublyLinkedList) PopBack() (interface{}, error) {
	switch {
	default:
		panic("bad PopBack implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		return 0, ErrEmptyList
	case ll.Head != nil && ll.Tail != nil && ll.Tail.Prev == nil: // 1 element
		v := ll.Tail.Val
		ll.Head = nil
		ll.Tail = nil

		return v, nil
	case ll.Head != nil && ll.Tail != nil && ll.Tail.Prev != nil: // >1 element
		v := ll.Tail.Val
		ll.Tail.Prev.Next = nil
		ll.Tail = ll.Tail.Prev

		return v, nil
	}
}
