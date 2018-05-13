package linkedlist

import (
	"bytes"
	"errors"
	"fmt"
)

// ListNode is a node in a linked list.
type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
}

// NewListNode constructs a new ListNode with the given value & no next/prev links.
func NewListNode(v int) *ListNode {
	return &ListNode{
		val:  v,
		next: nil,
		prev: nil,
	}
}

// DoublyLinkedList is a doubly-linked list with Head and Tail.
type DoublyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

// NewDoublyLinkedList constructs a doubly linked list from a sequence of integers.
func NewDoublyLinkedList(vs ...int) *DoublyLinkedList {
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
		cur.next = NewListNode(vs[i])
		cur.next.prev = cur
		cur = cur.next
	}

	ll.Tail = cur

	return ll
}

// Reverse reverses the given linked list in-place.
func (ll *DoublyLinkedList) Reverse() {
	if ll.Head == nil || ll.Head.next == nil {
		return
	}

	// construct singly-linked list from the back
	dummy := NewListNode(-1)
	cur := dummy
	n := ll.Tail
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
	ll.Head, ll.Tail = ll.Tail, ll.Head
	ll.Head.prev = nil
}

func (ll *DoublyLinkedList) String() string {
	buf := bytes.NewBuffer([]byte{'{'})

	for cur := ll.Head; cur != nil; cur = cur.next {
		buf.WriteString(fmt.Sprintf("%v <-> ", cur.val))
	}

	buf.WriteByte('}')

	return buf.String()
}

// DebugString prints the linked list with both node's value, next & prev pointers.
func (ll *DoublyLinkedList) DebugString() string {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("Head= %p; ", ll.Head))

	for cur := ll.Head; cur != nil; cur = cur.next {
		buf.WriteString(fmt.Sprintf("[prev= %p, val= %p (%v), next= %p] <-> ", cur.prev, cur, cur.val, cur.next))
	}

	buf.WriteString(fmt.Sprintf("; Tail= %p; ", ll.Tail))
	buf.WriteByte('}')

	return buf.String()
}

// PushFront pushes a new value before Head.
func (ll *DoublyLinkedList) PushFront(v int) {
	n := NewListNode(v)

	switch {
	default:
		panic("bad PushFront implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		ll.Head = n
		ll.Tail = n
	case ll.Head != nil && ll.Tail != nil: // non-empty list
		n.next = ll.Head
		ll.Head.prev = n

		ll.Head = n
	}
}

// PushBack pushes a new value after Tail.
func (ll *DoublyLinkedList) PushBack(v int) {
	n := NewListNode(v)

	switch {
	default:
		panic("bad PushBack implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		ll.Head = n
		ll.Tail = n
	case ll.Head != nil && ll.Tail != nil: // non-empty list
		ll.Tail.next = n
		n.prev = ll.Tail

		ll.Tail = n
	}
}

var (
	ErrEmptyList = errors.New("list is empty")
)

// PopFront posp the element at Head. It returns error if the linked list is empty.
func (ll *DoublyLinkedList) PopFront() (int, error) {
	switch {
	default:
		panic("bad PopFront implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		return 0, ErrEmptyList
	case ll.Head != nil && ll.Tail != nil && ll.Head.next == nil: // 1 element
		v := ll.Head.val
		ll.Head = nil
		ll.Tail = nil

		return v, nil
	case ll.Head != nil && ll.Tail != nil && ll.Head.next != nil: // >1 element
		v := ll.Head.val
		ll.Head.next.prev = nil
		ll.Head = ll.Head.next

		return v, nil
	}
}

// PopBack pops the element at Tail. It returns error if the linked list is empty.
func (ll *DoublyLinkedList) PopBack() (int, error) {
	switch {
	default:
		panic("bad PopBack implementation")
	case ll.Head == nil && ll.Tail == nil: // empty list
		return 0, ErrEmptyList
	case ll.Head != nil && ll.Tail != nil && ll.Tail.prev == nil: // 1 element
		v := ll.Tail.val
		ll.Head = nil
		ll.Tail = nil

		return v, nil
	case ll.Head != nil && ll.Tail != nil && ll.Tail.prev != nil: // >1 element
		v := ll.Tail.val
		ll.Tail.prev.next = nil
		ll.Tail = ll.Tail.prev

		return v, nil
	}
}
