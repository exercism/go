package linkedlist

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	for _, tc := range newListTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range reverseTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)
			actual.Reverse()

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

func TestPushPop(t *testing.T) {
	for _, tc := range pushPopTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)
			for _, ac := range tc.actions {
				ac(t, actual)
			}

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

// checkDoublyLinkedList checks that the linked list is constructed correctly.
func checkDoublyLinkedList(t *testing.T, ll *List, expected []interface{}) {
	// check that length and elements are correct (scan once from begin -> end)
	elem, count, idx := ll.Head, 0, 0
	for ; elem != nil && idx < len(expected); elem, count, idx = elem.Next, count+1, idx+1 {
		if elem.Val != expected[idx] {
			t.Errorf("wrong value from %d-th element, expected= %v, got= %v", idx, expected[idx], elem.Val)
		}
	}
	if !(elem == nil && idx == len(expected)) {
		t.Errorf("expected %d elements, got= %d", len(expected), count)
	}

	// if elements are the same, we also need to examine the links (Next & Prev)
	switch {
	case ll.Head == nil && ll.Tail == nil: // empty list
		return
	case ll.Head != nil && ll.Tail != nil && ll.Head.Next == nil: // 1 element
		valid := ll.Head == ll.Tail &&
			ll.Head.Next == nil &&
			ll.Head.Prev == nil &&
			ll.Tail.Next == nil &&
			ll.Tail.Prev == nil

		if !valid {
			t.Errorf("expected to only have 1 element and no links, got= %v", ll.debugString())
		}
	}

	// >1 element
	if ll.Head.Prev != nil {
		t.Errorf("expected Head.prev == nil, got= %v", ll.Head.Prev)
	}

	prev := ll.Head
	cur := ll.Head.Next
	for idx := 0; cur != nil; idx++ {
		if !(prev.Next == cur && cur.Prev == prev) {
			t.Errorf("%d-th element's links is wrong", idx)
		}

		prev = cur
		cur = cur.Next
	}

	if ll.Tail.Next != nil {
		t.Errorf("expected Tail.next == nil, got= %v", ll.Head.Prev)
	}
}

// debugString prints the linked list with both node's Val, Next & Prev pointers.
func (ll *List) debugString() string {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("Head= %p; ", ll.Head))

	for cur := ll.Head; cur != nil; cur = cur.Next {
		buf.WriteString(fmt.Sprintf("[Prev= %p, Val= %p (%v), Next= %p] <-> ", cur.Prev, cur, cur.Val, cur.Next))
	}

	buf.WriteString(fmt.Sprintf("; Tail= %p; ", ll.Tail))
	buf.WriteByte('}')

	return buf.String()
}
