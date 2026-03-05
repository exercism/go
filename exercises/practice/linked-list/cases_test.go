// This file contains tests from the shared problem specifications repo.
package linkedlist

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 738e1e2 Rename canonical-schema to canonical-data.schema.json (#1908)

type Operation struct {
	operation string
	value     int
	expected  int
}

var testCases = []struct {
	description string
	operations  []Operation
}{
	{
		description: "pop gets element from the list",
		operations: []Operation{
			{operation: "Push", value: 7},
			{operation: "Pop", expected: 7},
		},
	},
	{
		description: "push/pop respectively add/remove at the end of the list",
		operations: []Operation{
			{operation: "Push", value: 11},
			{operation: "Push", value: 13},
			{operation: "Pop", expected: 13},
			{operation: "Pop", expected: 11},
		},
	},
	{
		description: "shift gets an element from the list",
		operations: []Operation{
			{operation: "Push", value: 17},
			{operation: "Shift", expected: 17},
		},
	},
	{
		description: "shift gets first element from the list",
		operations: []Operation{
			{operation: "Push", value: 23},
			{operation: "Push", value: 5},
			{operation: "Shift", expected: 23},
			{operation: "Shift", expected: 5},
		},
	},
	{
		description: "unshift adds element at start of the list",
		operations: []Operation{
			{operation: "Unshift", value: 23},
			{operation: "Unshift", value: 5},
			{operation: "Shift", expected: 5},
			{operation: "Shift", expected: 23},
		},
	},
	{
		description: "pop, push, shift, and unshift can be used in any order",
		operations: []Operation{
			{operation: "Push", value: 1},
			{operation: "Push", value: 2},
			{operation: "Pop", expected: 2},
			{operation: "Push", value: 3},
			{operation: "Shift", expected: 1},
			{operation: "Unshift", value: 4},
			{operation: "Push", value: 5},
			{operation: "Shift", expected: 4},
			{operation: "Pop", expected: 5},
			{operation: "Shift", expected: 3},
		},
	},
	{
		description: "count an empty list",
		operations: []Operation{
			{operation: "Count", expected: 0},
		},
	},
	{
		description: "count a list with items",
		operations: []Operation{
			{operation: "Push", value: 37},
			{operation: "Push", value: 1},
			{operation: "Count", expected: 2},
		},
	},
	{
		description: "count is correct after mutation",
		operations: []Operation{
			{operation: "Push", value: 31},
			{operation: "Count", expected: 1},
			{operation: "Unshift", value: 43},
			{operation: "Count", expected: 2},
			{operation: "Shift", expected: 0},
			{operation: "Count", expected: 1},
			{operation: "Pop", expected: 0},
			{operation: "Count", expected: 0},
		},
	},
	{
		description: "popping to empty doesn't break the list",
		operations: []Operation{
			{operation: "Push", value: 41},
			{operation: "Push", value: 59},
			{operation: "Pop", expected: 0},
			{operation: "Pop", expected: 0},
			{operation: "Push", value: 47},
			{operation: "Count", expected: 1},
			{operation: "Pop", expected: 47},
		},
	},
	{
		description: "shifting to empty doesn't break the list",
		operations: []Operation{
			{operation: "Push", value: 41},
			{operation: "Push", value: 59},
			{operation: "Shift", expected: 0},
			{operation: "Shift", expected: 0},
			{operation: "Push", value: 47},
			{operation: "Count", expected: 1},
			{operation: "Shift", expected: 47},
		},
	},
	{
		description: "deletes the only element",
		operations: []Operation{
			{operation: "Push", value: 61},
			{operation: "Delete", value: 61},
			{operation: "Count", expected: 0},
		},
	},
	{
		description: "deletes the element with the specified value from the list",
		operations: []Operation{
			{operation: "Push", value: 71},
			{operation: "Push", value: 83},
			{operation: "Push", value: 79},
			{operation: "Delete", value: 83},
			{operation: "Count", expected: 2},
			{operation: "Pop", expected: 79},
			{operation: "Shift", expected: 71},
		},
	},
	{
		description: "deletes the element with the specified value from the list, re-assigns tail",
		operations: []Operation{
			{operation: "Push", value: 71},
			{operation: "Push", value: 83},
			{operation: "Push", value: 79},
			{operation: "Delete", value: 83},
			{operation: "Count", expected: 2},
			{operation: "Pop", expected: 79},
			{operation: "Pop", expected: 71},
		},
	},
	{
		description: "deletes the element with the specified value from the list, re-assigns head",
		operations: []Operation{
			{operation: "Push", value: 71},
			{operation: "Push", value: 83},
			{operation: "Push", value: 79},
			{operation: "Delete", value: 83},
			{operation: "Count", expected: 2},
			{operation: "Shift", expected: 71},
			{operation: "Shift", expected: 79},
		},
	},
	{
		description: "deletes the first of two elements",
		operations: []Operation{
			{operation: "Push", value: 97},
			{operation: "Push", value: 101},
			{operation: "Delete", value: 97},
			{operation: "Count", expected: 1},
			{operation: "Pop", expected: 101},
		},
	},
	{
		description: "deletes the second of two elements",
		operations: []Operation{
			{operation: "Push", value: 97},
			{operation: "Push", value: 101},
			{operation: "Delete", value: 101},
			{operation: "Count", expected: 1},
			{operation: "Pop", expected: 97},
		},
	},
	{
		description: "delete does not modify the list if the element is not found",
		operations: []Operation{
			{operation: "Push", value: 89},
			{operation: "Delete", value: 103},
			{operation: "Count", expected: 1},
		},
	},
	{
		description: "deletes only the first occurrence",
		operations: []Operation{
			{operation: "Push", value: 73},
			{operation: "Push", value: 9},
			{operation: "Push", value: 9},
			{operation: "Push", value: 107},
			{operation: "Delete", value: 9},
			{operation: "Count", expected: 3},
			{operation: "Pop", expected: 107},
			{operation: "Pop", expected: 9},
			{operation: "Pop", expected: 73},
		},
	},
}
