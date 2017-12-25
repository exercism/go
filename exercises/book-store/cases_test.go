package bookstore

// Source: exercism/problem-specifications
// Commit: 78006a2 move "targetgrouping" to "comments"
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	basket      []int
	expected    float64
}{
	{
		description: "Only a single book",
		basket:      []int{1},
		expected:    8.00,
	},
	{
		description: "Two of the same book",
		basket:      []int{2, 2},
		expected:    16.00,
	},
	{
		description: "Empty basket",
		basket:      []int{},
		expected:    0.00,
	},
	{
		description: "Two different books",
		basket:      []int{1, 2},
		expected:    15.20,
	},
	{
		description: "Three different books",
		basket:      []int{1, 2, 3},
		expected:    21.60,
	},
	{
		description: "Four different books",
		basket:      []int{1, 2, 3, 4},
		expected:    25.60,
	},
	{
		description: "Five different books",
		basket:      []int{1, 2, 3, 4, 5},
		expected:    30.00,
	},
	{
		description: "Two groups of four is cheaper than group of five plus group of three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5},
		expected:    51.20,
	},
	{
		description: "Group of four plus group of two is cheaper than two groups of three",
		basket:      []int{1, 1, 2, 2, 3, 4},
		expected:    40.80,
	},
	{
		description: "Two each of first 4 books and 1 copy each of rest",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		expected:    55.60,
	},
	{
		description: "Two copies of each book",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		expected:    60.00,
	},
	{
		description: "Three copies of first book and 2 each of remaining",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1},
		expected:    68.00,
	},
	{
		description: "Three each of first 2 books and 2 each of remaining books",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 2},
		expected:    75.20,
	},
	{
		description: "Four groups of four are cheaper than two groups each of five and three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5, 1, 1, 2, 2, 3, 3, 4, 5},
		expected:    102.40,
	},
}
