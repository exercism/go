package linkedlist

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNodeTraversal(t *testing.T) {
	ll := NewList(1, 2, 3, 4, 5)
	node := ll.head.next.next

	if want, got := ll.head, node.First(); want != got {
		t.Errorf("%#v.First() is wrong, expected= %#v, got= %#v", node, want, got)
	}
	if want, got := ll.tail, node.Last(); want != got {
		t.Errorf("%#v.Last() is wrong, expected= %#v, got= %#v", node, want, got)
	}
	if want, got := ll.tail.prev, node.Next(); want != got {
		t.Errorf("%#v.Next() is wrong, expected= %#v, got= %#v", node, want, got)
	}
	if want, got := ll.head.next, node.Prev(); want != got {
		t.Errorf("%#v.Prev() is wrong, expected= %#v, got= %#v", node, want, got)
	}
}

func TestListTraversal(t *testing.T) {
	ll := NewList(1, 2, 3, 4, 5)

	if want, got := ll.head, ll.First(); want != got {
		t.Errorf("%#v.First() is wrong, expected= %#v, got= %#v", ll, want, got)
	}
	if want, got := ll.tail, ll.Last(); want != got {
		t.Errorf("%#v.First() is wrong, expected= %#v, got= %#v", ll, want, got)
	}
}

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
	elem, count, idx := ll.head, 0, 0
	for ; elem != nil && idx < len(expected); elem, count, idx = elem.next, count+1, idx+1 {
		if elem.Val != expected[idx] {
			t.Errorf("wrong value from %d-th element, expected= %v, got= %v", idx, expected[idx], elem.Val)
		}
	}
	if !(elem == nil && idx == len(expected)) {
		t.Errorf("expected %d elements, got= %d", len(expected), count)
	}

	// if elements are the same, we also need to examine the links (next & prev)
	switch {
	case ll.head == nil && ll.tail == nil: // empty list
		return
	case ll.head != nil && ll.tail != nil && ll.head.next == nil: // 1 element
		valid := ll.head == ll.tail &&
			ll.head.next == nil &&
			ll.head.prev == nil &&
			ll.tail.next == nil &&
			ll.tail.prev == nil

		if !valid {
			t.Errorf("expected to only have 1 element and no links, got= %v", ll.debugString())
		}
	}

	// >1 element
	if ll.head.prev != nil {
		t.Errorf("expected Head.prev == nil, got= %v", ll.head.prev)
	}

	prev := ll.head
	cur := ll.head.next
	for idx := 0; cur != nil; idx++ {
		if !(prev.next == cur && cur.prev == prev) {
			t.Errorf("%d-th element's links is wrong", idx)
		}

		prev = cur
		cur = cur.next
	}

	if ll.tail.next != nil {
		t.Errorf("expected Tail.next == nil, got= %v", ll.head.prev)
	}
}

// debugString prints the linked list with both node's Val, next & prev pointers.
func (ll *List) debugString() string {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("Head= %p; ", ll.head))

	for cur := ll.head; cur != nil; cur = cur.next {
		buf.WriteString(fmt.Sprintf("[prev= %p, Val= %p (%v), next= %p] <-> ", cur.prev, cur, cur.Val, cur.next))
	}

	buf.WriteString(fmt.Sprintf("; Tail= %p; ", ll.tail))
	buf.WriteByte('}')

	return buf.String()
}
