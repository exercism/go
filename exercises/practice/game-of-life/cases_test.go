package gameoflife

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 381bb40 Add Game of Life exercise (#2281)

var testCases = []struct {
	description string
	input       [][]int
	expected    [][]int
}{
	{
		description: "empty matrix",
		input:       [][]int{},
		expected:    [][]int{},
	},
	{
		description: "live cells with zero live neighbors die",
		input: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
		expected: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	},
	{
		description: "live cells with only one live neighbor die",
		input: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 1, 0},
		},
		expected: [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	},
	{
		description: "live cells with two live neighbors stay alive",
		input: [][]int{
			{1, 0, 1},
			{1, 0, 1},
			{1, 0, 1},
		},
		expected: [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
	},
	{
		description: "live cells with three live neighbors stay alive",
		input: [][]int{
			{0, 1, 0},
			{1, 0, 0},
			{1, 1, 0},
		},
		expected: [][]int{
			{0, 0, 0},
			{1, 0, 0},
			{1, 1, 0},
		},
	},
	{
		description: "dead cells with three live neighbors become alive",
		input: [][]int{
			{1, 1, 0},
			{0, 0, 0},
			{1, 0, 0},
		},
		expected: [][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
		},
	},
	{
		description: "live cells with four or more neighbors die",
		input: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		expected: [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		},
	},
	{
		description: "bigger matrix",
		input: [][]int{
			{1, 1, 0, 1, 1, 0, 0, 0},
			{1, 0, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 1, 1, 0},
			{1, 0, 0, 0, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 1, 1, 1},
			{0, 0, 1, 0, 1, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 1, 1},
		},
		expected: [][]int{
			{1, 1, 0, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 0},
			{1, 0, 1, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 0, 0, 1, 0, 0, 1},
			{1, 1, 0, 1, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 1},
		},
	},
}
