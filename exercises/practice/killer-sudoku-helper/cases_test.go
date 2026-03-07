package killersudokuhelper

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7cd0ab1 Add exercise: killer-sudoku-helper (#2007)

var testCases = []struct {
	description string
	sum         int
	size        int
	exclude     []int
	expected    [][]int
}{
	{
		description: "1",
		sum:         1,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{1}},
	},
	{
		description: "2",
		sum:         2,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{2}},
	},
	{
		description: "3",
		sum:         3,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{3}},
	},
	{
		description: "4",
		sum:         4,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{4}},
	},
	{
		description: "5",
		sum:         5,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{5}},
	},
	{
		description: "6",
		sum:         6,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{6}},
	},
	{
		description: "7",
		sum:         7,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{7}},
	},
	{
		description: "8",
		sum:         8,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{8}},
	},
	{
		description: "9",
		sum:         9,
		size:        1,
		exclude:     []int{},
		expected:    [][]int{[]int{9}},
	},
	{
		description: "Cage with sum 45 contains all digits 1:9",
		sum:         45,
		size:        9,
		exclude:     []int{},
		expected:    [][]int{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	},
	{
		description: "Cage with only 1 possible combination",
		sum:         7,
		size:        3,
		exclude:     []int{},
		expected:    [][]int{[]int{1, 2, 4}},
	},
	{
		description: "Cage with several combinations",
		sum:         10,
		size:        2,
		exclude:     []int{},
		expected:    [][]int{[]int{1, 9}, []int{2, 8}, []int{3, 7}, []int{4, 6}},
	},
	{
		description: "Cage with several combinations that is restricted",
		sum:         10,
		size:        2,
		exclude:     []int{1, 4},
		expected:    [][]int{[]int{2, 8}, []int{3, 7}},
	},
}
