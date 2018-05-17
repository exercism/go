package linkedlist

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var testCases = []struct {
		name     string
		in       []interface{}
		expected []interface{}
	}{
		{
			name:     "from 5 elements",
			in:       []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "from 2 elements",
			in:       []interface{}{1, 2},
			expected: []interface{}{1, 2},
		},
		{
			name:     "from no element",
			in:       []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "from 1 element",
			in:       []interface{}{999},
			expected: []interface{}{999},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewDoublyLinkedList(tc.in...)

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

func TestReverse(t *testing.T) {
	var testCases = []struct {
		name     string
		in       []interface{}
		expected []interface{}
	}{
		{
			name:     "from 5 elements",
			in:       []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{5, 4, 3, 2, 1},
		},
		{
			name:     "from 2 elements",
			in:       []interface{}{1, 2},
			expected: []interface{}{2, 1},
		},
		{
			name:     "from no element",
			in:       []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "from 1 element",
			in:       []interface{}{999},
			expected: []interface{}{999},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewDoublyLinkedList(tc.in...)
			actual.Reverse()

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

// checkedAction calls a function of the linked list and (possibly) checks the result
type checkedAction func(*testing.T, *DoublyLinkedList)

func pushFront(arg interface{}) checkedAction {
	return func(t *testing.T, ll *DoublyLinkedList) {
		ll.PushFront(arg)
	}
}

func pushBack(arg interface{}) checkedAction {
	return func(t *testing.T, ll *DoublyLinkedList) {
		ll.PushBack(arg)
	}
}

func popFront(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *DoublyLinkedList) {
		v, err := ll.PopFront()
		if err != expectedErr {
			t.Errorf("PopFront() returned wrong, expected no error, got= %v", err)
		}

		if v != expected {
			t.Errorf("PopFront() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func popBack(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *DoublyLinkedList) {
		v, err := ll.PopBack()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func TestOps(t *testing.T) {
	var testCases = []struct {
		name     string
		in       []interface{}
		actions  []checkedAction
		expected []interface{}
	}{
		{
			name: "PushFront only",
			in:   []interface{}{},
			actions: []checkedAction{
				pushFront(4),
				pushFront(3),
				pushFront(2),
				pushFront(1),
			},
			expected: []interface{}{1, 2, 3, 4},
		},
		{
			name: "PushBack only",
			in:   []interface{}{},
			actions: []checkedAction{
				pushBack(1),
				pushBack(2),
				pushBack(3),
				pushBack(4),
			},
			expected: []interface{}{1, 2, 3, 4},
		},
		{
			name: "PopFront only, pop some elements",
			in:   []interface{}{1, 2, 3, 4},
			actions: []checkedAction{
				popFront(1, nil),
				popFront(2, nil),
			},
			expected: []interface{}{3, 4},
		},
		{
			name: "PopFront only, pop till empty",
			in:   []interface{}{1, 2, 3, 4},
			actions: []checkedAction{
				popFront(1, nil),
				popFront(2, nil),
				popFront(3, nil),
				popFront(4, nil),
				popFront(0, ErrEmptyList),
			},
			expected: []interface{}{},
		},
		{
			name: "PopBack only, pop some elements",
			in:   []interface{}{1, 2, 3, 4},
			actions: []checkedAction{
				popBack(4, nil),
				popBack(3, nil),
			},
			expected: []interface{}{1, 2},
		},
		{
			name: "PopBack only, pop till empty",
			in:   []interface{}{1, 2, 3, 4},
			actions: []checkedAction{
				popBack(4, nil),
				popBack(3, nil),
				popBack(2, nil),
				popBack(1, nil),
				popBack(0, ErrEmptyList),
			},
			expected: []interface{}{},
		},
		{
			name: "mixed actions",
			in:   []interface{}{2, 3},
			actions: []checkedAction{
				pushFront(1),
				pushBack(4),
				popFront(1, nil),
				popFront(2, nil),
				popBack(4, nil),
				popBack(3, nil),
				popBack(0, ErrEmptyList),
				popFront(0, ErrEmptyList),
				pushFront(8),
				pushBack(7),
				pushFront(9),
				pushBack(6),
			},
			expected: []interface{}{9, 8, 7, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewDoublyLinkedList(tc.in...)
			for _, ac := range tc.actions {
				ac(t, actual)
			}

			checkDoublyLinkedList(t, actual, tc.expected)
		})
	}
}

// checkDoublyLinkedList checks that the linked list is constructed correctly.
func checkDoublyLinkedList(t *testing.T, ll *DoublyLinkedList, expected []interface{}) {
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

	// if elements are the same, we also need to examine the links (next & prev)
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

// debugString prints the linked list with both node's value, Next & Prev pointers.
func (ll *DoublyLinkedList) debugString() string {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("Head= %p; ", ll.Head))

	for cur := ll.Head; cur != nil; cur = cur.Next {
		buf.WriteString(fmt.Sprintf("[Prev= %p, Val= %p (%v), Next= %p] <-> ", cur.Prev, cur, cur.Val, cur.Next))
	}

	buf.WriteString(fmt.Sprintf("; Tail= %p; ", ll.Tail))
	buf.WriteByte('}')

	return buf.String()
}
