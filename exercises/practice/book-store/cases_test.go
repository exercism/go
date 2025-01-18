package bookstore

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: ece572c [book-store] Made test descriptions consistent (#2016)

var testCases = []struct {
	description string
	basket      []int
	expected    int
}{
	{
		description: "Only a single book",
		basket:      []int{1},
		expected:    800,
	},
	{
		description: "Two of the same book",
		basket:      []int{2, 2},
		expected:    1600,
	},
	{
		description: "Empty basket",
		basket:      []int{},
		expected:    0,
	},
	{
		description: "Two different books",
		basket:      []int{1, 2},
		expected:    1520,
	},
	{
		description: "Three different books",
		basket:      []int{1, 2, 3},
		expected:    2160,
	},
	{
		description: "Four different books",
		basket:      []int{1, 2, 3, 4},
		expected:    2560,
	},
	{
		description: "Five different books",
		basket:      []int{1, 2, 3, 4, 5},
		expected:    3000,
	},
	{
		description: "Two groups of four is cheaper than group of five plus group of three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5},
		expected:    5120,
	},
	{
		description: "Two groups of four is cheaper than groups of five and three",
		basket:      []int{1, 1, 2, 3, 4, 4, 5, 5},
		expected:    5120,
	},
	{
		description: "Group of four plus group of two is cheaper than two groups of three",
		basket:      []int{1, 1, 2, 2, 3, 4},
		expected:    4080,
	},
	{
		description: "Two each of first four books and one copy each of rest",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		expected:    5560,
	},
	{
		description: "Two copies of each book",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		expected:    6000,
	},
	{
		description: "Three copies of first book and two each of remaining",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1},
		expected:    6800,
	},
	{
		description: "Three each of first two books and two each of remaining books",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 2},
		expected:    7520,
	},
	{
		description: "Four groups of four are cheaper than two groups each of five and three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 5},
		expected:    10240,
	},
	{
		description: "Check that groups of four are created properly even when there are more groups of three than groups of five",
		basket:      []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 5, 5},
		expected:    14560,
	},
	{
		description: "One group of one and four is cheaper than one group of two and three",
		basket:      []int{1, 1, 2, 3, 4},
		expected:    3360,
	},
	{
		description: "One group of one and two plus three groups of four is cheaper than one group of each size",
		basket:      []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5},
		expected:    10000,
	},
}
