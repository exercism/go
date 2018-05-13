package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	var testCases = []struct {
		name     string
		in       []int
		expected string
	}{
		{
			name:     "from []int (5 elements)",
			in:       []int{1, 2, 3, 4, 5},
			expected: "{1 <-> 2 <-> 3 <-> 4 <-> 5 <-> }",
		},
		{
			name:     "from []int (2 elements)",
			in:       []int{1, 2},
			expected: "{1 <-> 2 <-> }",
		},
		{
			name:     "from empty []int",
			in:       []int{},
			expected: "{}",
		},
		{
			name:     "from 1-element []int",
			in:       []int{999},
			expected: "{999 <-> }",
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
		in       []int
		expected string
	}{
		{
			name:     "from []int (5 elements)",
			in:       []int{1, 2, 3, 4, 5},
			expected: "{5 <-> 4 <-> 3 <-> 2 <-> 1 <-> }",
		},
		{
			name:     "from []int (2 elements)",
			in:       []int{1, 2},
			expected: "{2 <-> 1 <-> }",
		},
		{
			name:     "from empty []int",
			in:       []int{},
			expected: "{}",
		},
		{
			name:     "from 1-element []int",
			in:       []int{999},
			expected: "{999 <-> }",
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

func pushFront(arg int) checkedAction {
	return func(t *testing.T, ll *DoublyLinkedList) {
		ll.PushFront(arg)
	}
}

func pushBack(arg int) checkedAction {
	return func(t *testing.T, ll *DoublyLinkedList) {
		ll.PushBack(arg)
	}
}

func popFront(expected int, expectedErr error) checkedAction {
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

func popBack(expected int, expectedErr error) checkedAction {
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
		in       []int
		actions  []checkedAction
		expected string
	}{
		{
			name: "PushFront only",
			in:   []int{},
			actions: []checkedAction{
				pushFront(4),
				pushFront(3),
				pushFront(2),
				pushFront(1),
			},
			expected: "{1 <-> 2 <-> 3 <-> 4 <-> }",
		},
		{
			name: "PushBack only",
			in:   []int{},
			actions: []checkedAction{
				pushBack(1),
				pushBack(2),
				pushBack(3),
				pushBack(4),
			},
			expected: "{1 <-> 2 <-> 3 <-> 4 <-> }",
		},
		{
			name: "PopFront only, pop some elements",
			in:   []int{1, 2, 3, 4},
			actions: []checkedAction{
				popFront(1, nil),
				popFront(2, nil),
			},
			expected: "{3 <-> 4 <-> }",
		},
		{
			name: "PopFront only, pop till empty",
			in:   []int{1, 2, 3, 4},
			actions: []checkedAction{
				popFront(1, nil),
				popFront(2, nil),
				popFront(3, nil),
				popFront(4, nil),
				popFront(0, ErrEmptyList),
			},
			expected: "{}",
		},
		{
			name: "PopBack only, pop some elements",
			in:   []int{1, 2, 3, 4},
			actions: []checkedAction{
				popBack(4, nil),
				popBack(3, nil),
			},
			expected: "{1 <-> 2 <-> }",
		},
		{
			name: "PopBack only, pop till empty",
			in:   []int{1, 2, 3, 4},
			actions: []checkedAction{
				popBack(4, nil),
				popBack(3, nil),
				popBack(2, nil),
				popBack(1, nil),
				popBack(0, ErrEmptyList),
			},
			expected: "{}",
		},
		{
			name: "mixed actions",
			in:   []int{2, 3},
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
			expected: "{9 <-> 8 <-> 7 <-> 6 <-> }",
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
func checkDoublyLinkedList(t *testing.T, ll *DoublyLinkedList, expected string) {
	// quick check by looking at the element's value
	if ll.String() != expected {
		t.Errorf("expected= %s, got= %v", expected, ll)
	}

	// if they are the same, we also need to examine the links (next & prev)
	switch {
	// TODO(exklamationmark): add more patterns. the current ones are not exhaustive -> might have edge cases
	case ll.Head == nil && ll.Tail == nil: // empty list
		return // NOP
	case ll.Head != nil && ll.Tail != nil && ll.Head.next == nil: // 1 element
		valid := ll.Head == ll.Tail &&
			ll.Head.next == nil &&
			ll.Head.prev == nil &&
			ll.Tail.next == nil &&
			ll.Tail.prev == nil

		if !valid {
			t.Errorf("expected to only have 1 element and no links, got= %v", ll.DebugString())
		}
	}

	// >1 element

	if ll.Head.prev != nil {
		t.Errorf("expected Head.prev == nil, got= %v", ll.Head.prev)
	}

	prev := ll.Head
	cur := ll.Head.next
	for idx := 0; cur != nil; idx++ {
		if !(prev.next == cur && cur.prev == prev) {
			t.Errorf("%d-th element's links is wrong", idx)
		}

		prev = cur
		cur = cur.next
	}

	if ll.Tail.next != nil {
		t.Errorf("expected Tail.next == nil, got= %v", ll.Head.prev)
	}
}
