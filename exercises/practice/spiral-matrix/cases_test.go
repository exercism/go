package spiralmatrix

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

type testCase struct {
	description string
	input       int
	expected    [][]int
}

var testCases = []testCase{
	{
		description: "empty spiral",
		input:       0,
		expected:    [][]int{},
	},
	{
		description: "trivial spiral",
		input:       1,
		expected: [][]int{
			[]int{1},
		},
	},
	{
		description: "spiral of size 2",
		input:       2,
		expected: [][]int{
			[]int{1, 2},
			[]int{4, 3},
		},
	},
	{
		description: "spiral of size 3",
		input:       3,
		expected: [][]int{
			[]int{1, 2, 3},
			[]int{8, 9, 4},
			[]int{7, 6, 5},
		},
	},
	{
		description: "spiral of size 4",
		input:       4,
		expected: [][]int{
			[]int{1, 2, 3, 4},
			[]int{12, 13, 14, 5},
			[]int{11, 16, 15, 6},
			[]int{10, 9, 8, 7},
		},
	},
	{
		description: "spiral of size 5",
		input:       5,
		expected: [][]int{
			[]int{1, 2, 3, 4, 5},
			[]int{16, 17, 18, 19, 6},
			[]int{15, 24, 25, 20, 7},
			[]int{14, 23, 22, 21, 8},
			[]int{13, 12, 11, 10, 9},
		},
	},
}
