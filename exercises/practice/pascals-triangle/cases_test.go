package pascalstriangle

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: c534e4d pascals-triangle: format each row as one line (#1996)

type testCase struct {
	description string
	rows        int
	expected    [][]int
}

var testCases = []testCase{
	{
		description: "zero rows",
		rows:        0,
		expected:    [][]int{},
	},
	{
		description: "single row",
		rows:        1,
		expected: [][]int{
			[]int{1},
		},
	},
	{
		description: "two rows",
		rows:        2,
		expected: [][]int{
			[]int{1},
			[]int{1, 1},
		},
	},
	{
		description: "three rows",
		rows:        3,
		expected: [][]int{
			[]int{1},
			[]int{1, 1},
			[]int{1, 2, 1},
		},
	},
	{
		description: "four rows",
		rows:        4,
		expected: [][]int{
			[]int{1},
			[]int{1, 1},
			[]int{1, 2, 1},
			[]int{1, 3, 3, 1},
		},
	},
	{
		description: "five rows",
		rows:        5,
		expected: [][]int{
			[]int{1},
			[]int{1, 1},
			[]int{1, 2, 1},
			[]int{1, 3, 3, 1},
			[]int{1, 4, 6, 4, 1},
		},
	},
	{
		description: "six rows",
		rows:        6,
		expected: [][]int{
			[]int{1},
			[]int{1, 1},
			[]int{1, 2, 1},
			[]int{1, 3, 3, 1},
			[]int{1, 4, 6, 4, 1},
			[]int{1, 5, 10, 10, 5, 1},
		},
	},
	{
		description: "ten rows",
		rows:        10,
		expected: [][]int{
			[]int{1},
			[]int{1, 1},
			[]int{1, 2, 1},
			[]int{1, 3, 3, 1},
			[]int{1, 4, 6, 4, 1},
			[]int{1, 5, 10, 10, 5, 1},
			[]int{1, 6, 15, 20, 15, 6, 1},
			[]int{1, 7, 21, 35, 35, 21, 7, 1},
			[]int{1, 8, 28, 56, 70, 56, 28, 8, 1},
			[]int{1, 9, 36, 84, 126, 126, 84, 36, 9, 1},
		},
	},
}
