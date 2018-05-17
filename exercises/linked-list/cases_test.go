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

// checkedAction calls a function of the linked list and (possibly) checks the result
type checkedAction func(*testing.T, *List)

func pushFront(arg interface{}) checkedAction {
	return func(t *testing.T, ll *List) {
		ll.PushFront(arg)
	}
}

func pushBack(arg interface{}) checkedAction {
	return func(t *testing.T, ll *List) {
		ll.PushBack(arg)
	}
}

func popFront(expected interface{}, expectedErr error) checkedAction {
	return func(t *testing.T, ll *List) {
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
	return func(t *testing.T, ll *List) {
		v, err := ll.PopBack()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}
