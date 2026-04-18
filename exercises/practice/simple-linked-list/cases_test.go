package simplelinkedlist

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 76ee5b7 [simple-linked-list] Add missing initialValues field (#2645)

type Operation struct {
	operation    string
	value        int
	expectedInt  int
	expectedInts []int
	expectedErr  string
}

type testCase struct {
	description   string
	initialValues []int
	operations    []Operation
}

var testCases = []testCase{
	{
		description:   "Empty list has length of zero",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Size", expectedInt: 0},
		},
	},
	{
		description:   "Singleton list has length of one",
		initialValues: []int{1},
		operations: []Operation{
			Operation{operation: "Size", expectedInt: 1},
		},
	},
	{
		description:   "Non-empty list has correct length",
		initialValues: []int{1, 2, 3},
		operations: []Operation{
			Operation{operation: "Size", expectedInt: 3},
		},
	},
	{
		description:   "Pop from empty list is an error",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Pop", expectedErr: "list is empty"},
		},
	},
	{
		description:   "Can pop from singleton list",
		initialValues: []int{1},
		operations: []Operation{
			Operation{operation: "Pop", expectedInt: 1},
		},
	},
	{
		description:   "Can pop from non-empty list",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Pop", expectedInt: 2},
		},
	},
	{
		description:   "Can pop multiple items",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Pop", expectedInt: 2},
			Operation{operation: "Pop", expectedInt: 1},
		},
	},
	{
		description:   "Pop updates the count",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Size", expectedInt: 2},
			Operation{operation: "Pop", expectedInt: 2},
			Operation{operation: "Size", expectedInt: 1},
			Operation{operation: "Pop", expectedInt: 1},
			Operation{operation: "Size", expectedInt: 0},
		},
	},
	{
		description:   "Can push to an empty list",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Push", value: 1},
		},
	},
	{
		description:   "Can push to a non-empty list",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Push", value: 3},
		},
	},
	{
		description:   "Push updates count",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Push", value: 3},
			Operation{operation: "Size", expectedInt: 3},
		},
	},
	{
		description:   "Push and pop",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Push", value: 1},
			Operation{operation: "Push", value: 2},
			Operation{operation: "Pop", expectedInt: 2},
			Operation{operation: "Push", value: 3},
			Operation{operation: "Size", expectedInt: 2},
			Operation{operation: "Pop", expectedInt: 3},
			Operation{operation: "Pop", expectedInt: 1},
			Operation{operation: "Size", expectedInt: 0},
		},
	},
	{
		description:   "Peek on empty list is an error",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Peek", expectedErr: "list is empty"},
		},
	},
	{
		description:   "Can peek on singleton list",
		initialValues: []int{1},
		operations: []Operation{
			Operation{operation: "Peek", expectedInt: 1},
		},
	},
	{
		description:   "Can peek on non-empty list",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Peek", expectedInt: 2},
		},
	},
	{
		description:   "Peek does not change the count",
		initialValues: []int{1, 2},
		operations: []Operation{
			Operation{operation: "Peek", expectedInt: 2},
			Operation{operation: "Size", expectedInt: 2},
		},
	},
	{
		description:   "Can peek after a pop and push",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Push", value: 1},
			Operation{operation: "Push", value: 2},
			Operation{operation: "Peek", expectedInt: 2},
			Operation{operation: "Pop", expectedInt: 2},
			Operation{operation: "Peek", expectedInt: 1},
			Operation{operation: "Push", value: 3},
			Operation{operation: "Peek", expectedInt: 3},
		},
	},
	{
		description:   "Empty linked list to list is empty",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Array", expectedInts: []int{}},
		},
	},
	{
		description:   "To list with multiple values",
		initialValues: []int{1, 2, 3},
		operations: []Operation{
			Operation{operation: "Array", expectedInts: []int{1, 2, 3}},
		},
	},
	{
		description:   "To list after a pop",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Push", value: 1},
			Operation{operation: "Push", value: 2},
			Operation{operation: "Push", value: 3},
			Operation{operation: "Pop", expectedInt: 3},
			Operation{operation: "Push", value: 4},
			Operation{operation: "Array", expectedInts: []int{1, 2, 4}},
		},
	},
	{
		description:   "Reversed empty list has same values",
		initialValues: []int{},
		operations: []Operation{
			Operation{operation: "Reverse"},
			Operation{operation: "Array", expectedInts: []int{}},
		},
	},
	{
		description:   "Reversed singleton list is same list",
		initialValues: []int{1},
		operations: []Operation{
			Operation{operation: "Reverse"},
			Operation{operation: "Array", expectedInts: []int{1}},
		},
	},
	{
		description:   "Reversed non-empty list is reversed",
		initialValues: []int{1, 2, 3},
		operations: []Operation{
			Operation{operation: "Reverse"},
			Operation{operation: "Size", expectedInt: 3},
			Operation{operation: "Pop", expectedInt: 1},
			Operation{operation: "Pop", expectedInt: 2},
			Operation{operation: "Pop", expectedInt: 3},
		},
	},
	{
		description:   "Double reverse",
		initialValues: []int{1, 2, 3},
		operations: []Operation{
			Operation{operation: "Reverse"},
			Operation{operation: "Reverse"},
			Operation{operation: "Pop", expectedInt: 3},
			Operation{operation: "Pop", expectedInt: 2},
			Operation{operation: "Pop", expectedInt: 1},
		},
	},
}
