package linkedlist

// This file contains custom track specific tests.

import "testing"

type newListTestCase[T comparable] struct {
	name     string
	in       []T
	expected []T
}

var newListTestCases = []newListTestCase[int] {
	{
		name:     "from 5 elements",
		in:       []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "from 2 elements",
		in:       []int{1, 2},
		expected: []int{1, 2},
	},
	{
		name:     "from no element",
		in:       []int{},
		expected: []int{},
	},
	{
		name:     "from 1 element",
		in:       []int{999},
		expected: []int{999},
	},
}

type reverseTestCase[T comparable] struct {
	name     string
	in       []T
	expected []T
}

var reverseTestCases = []reverseTestCase[int] {
	{
		name:     "from 5 elements",
		in:       []int{1, 2, 3, 4, 5},
		expected: []int{5, 4, 3, 2, 1},
	},
	{
		name:     "from 2 elements",
		in:       []int{1, 2},
		expected: []int{2, 1},
	},
	{
		name:     "from no element",
		in:       []int{},
		expected: []int{},
	},
	{
		name:     "from 1 element",
		in:       []int{999},
		expected: []int{999},
	},
}

type pushPopTestCase[T comparable] struct {
	name     string
	in       []T
	actions  []checkedAction[T]
	expected []T
}

var pushPopTestCases = []pushPopTestCase[int] {
	{
		name: "PushFront only",
		in:   []int{},
		actions: []checkedAction[int]{
			unshift(4),
			unshift(3),
			unshift(2),
			unshift(1),
		},
		expected: []int{1, 2, 3, 4},
	},
	{
		name: "PushBack only",
		in:   []int{},
		actions: []checkedAction[int]{
			push(1),
			push(2),
			push(3),
			push(4),
		},
		expected: []int{1, 2, 3, 4},
	},
	{
		name: "PopFront only, pop some elements",
		in:   []int{1, 2, 3, 4},
		actions: []checkedAction[int]{
			shift(1, nil),
			shift(2, nil),
		},
		expected: []int{3, 4},
	},
	{
		name: "PopFront only, pop till empty",
		in:   []int{1, 2, 3, 4},
		actions: []checkedAction[int]{
			shift(1, nil),
			shift(2, nil),
			shift(3, nil),
			shift(4, nil),
		},
		expected: []int{},
	},
	{
		name: "PopBack only, pop some elements",
		in:   []int{1, 2, 3, 4},
		actions: []checkedAction[int]{
			pop(4, nil),
			pop(3, nil),
		},
		expected: []int{1, 2},
	},
	{
		name: "PopBack only, pop till empty",
		in:   []int{1, 2, 3, 4},
		actions: []checkedAction[int]{
			pop(4, nil),
			pop(3, nil),
			pop(2, nil),
			pop(1, nil),
		},
		expected: []int{},
	},
	{
		name: "mixed actions",
		in:   []int{2, 3},
		actions: []checkedAction[int]{
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
		expected: []int{9, 8, 7, 6},
	},
}

var pushPopStringTestCases = []pushPopTestCase[string] {
	{
		name: "Push and pop strings",
		in:   []string{"one"},
		actions: []checkedAction[string]{
			push("two"),
			push("three"),
			pop("three", nil),
		},
		expected: []string{"one", "two"},
	},
}

// checkedAction calls a function of the linked list and (possibly) checks the result
type checkedAction[T comparable] func(*testing.T, *List[T])

func unshift[T comparable](arg T) checkedAction[T] {
	return func(t *testing.T, ll *List[T]) {
		ll.Unshift(arg)
	}
}

func push[T comparable](arg T) checkedAction[T] {
	return func(t *testing.T, ll *List[T]) {
		ll.Push(arg)
	}
}

func shift[T comparable](expected T, expectedErr error) checkedAction[T] {
	return func(t *testing.T, ll *List[T]) {
		v, err := ll.Shift()
		if err != expectedErr {
			t.Errorf("PopFront() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopFront() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func pop[T comparable](expected T, expectedErr error) checkedAction[T] {
	return func(t *testing.T, ll *List[T]) {
		v, err := ll.Pop()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}
