package linkedlist

import "testing"

var newListTestCases = []struct {
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

var reverseTestCases = []struct {
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

var pushPopTestCases = []struct {
	name     string
	in       []interface{}
	actions  []checkedAction
	expected []interface{}
}{
	{
		name: "PushFront only",
		in:   []interface{}{},
		actions: []checkedAction{
			unshift(4),
			unshift(3),
			unshift(2),
			unshift(1),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PushBack only",
		in:   []interface{}{},
		actions: []checkedAction{
			push(1),
			push(2),
			push(3),
			push(4),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PopFront only, pop some elements",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			shift(1, nil),
			shift(2, nil),
		},
		expected: []interface{}{3, 4},
	},
	{
		name: "PopFront only, pop till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			shift(1, nil),
			shift(2, nil),
			shift(3, nil),
			shift(4, nil),
		},
		expected: []interface{}{},
	},
	{
		name: "PopBack only, pop some elements",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			pop(4, nil),
			pop(3, nil),
		},
		expected: []interface{}{1, 2},
	},
	{
		name: "PopBack only, pop till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []checkedAction{
			pop(4, nil),
			pop(3, nil),
			pop(2, nil),
			pop(1, nil),
		},
		expected: []interface{}{},
	},
	{
		name: "mixed actions",
		in:   []interface{}{2, 3},
		actions: []checkedAction{
			unshift(1),
			push(4),
			shift(1, nil),
			shift(2, nil),
			pop(4, nil),
			pop(3, nil),
			unshift(8),
			push(7),
			unshift(9),
			push(6),
		},
		expected: []interface{}{9, 8, 7, 6},
	},
}

// checkedAction calls a function of the linked list and (possibly) checks the result
type checkedAction func(*testing.T, *List)

func unshift(arg interface{}) checkedAction {
	return func(t *testing.T, ll *List) {
		ll.Unshift(arg)
	}
}

func push(arg interface{}) checkedAction {
	return func(t *testing.T, ll *List) {
		ll.Push(arg)
	}
}

func shift(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *List) {
		v, err := ll.Shift()
		if err != expectedErr {
			t.Errorf("PopFront() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopFront() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func pop(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *List) {
		v, err := ll.Pop()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}
