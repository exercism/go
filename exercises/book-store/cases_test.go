package bookstore

// Source: exercism/problem-specifications
// Commit: a2ed4f7 book-store: Clarify expected monetary denomination
// Problem Specifications Version: 1.3.0

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
		description: "Group of four plus group of two is cheaper than two groups of three",
		basket:      []int{1, 1, 2, 2, 3, 4},
		expected:    4080,
	},
	{
		description: "Two each of first 4 books and 1 copy each of rest",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		expected:    5560,
	},
	{
		description: "Two copies of each book",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		expected:    6000,
	},
	{
		description: "Three copies of first book and 2 each of remaining",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1},
		expected:    6800,
	},
	{
		description: "Three each of first 2 books and 2 each of remaining books",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 2},
		expected:    7520,
	},
	{
		description: "Four groups of four are cheaper than two groups each of five and three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 5},
		expected:    10240,
	},
}
