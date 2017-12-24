package binarysearch

// Source: exercism/problem-specifications
// Commit: 6114d03 binary-search: Update json for new input policy
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	slice       []int
	key         int
	x           int
}{
	{
		description: "finds a value in an array with one element",
		slice:       []int{6},
		key:         6,
		x:           0,
	},
	{
		description: "finds a value in the middle of an array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         6,
		x:           3,
	},
	{
		description: "finds a value at the beginning of an array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         1,
		x:           0,
	},
	{
		description: "finds a value at the end of an array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         11,
		x:           6,
	},
	{
		description: "finds a value in an array of odd length",
		slice:       []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 634},
		key:         144,
		x:           9,
	},
	{
		description: "finds a value in an array of even length",
		slice:       []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377},
		key:         21,
		x:           5,
	},
	{
		description: "identifies that a value is not included in the array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         7,
		x:           -1,
	},
	{
		description: "a value smaller than the array's smallest value is not included",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         0,
		x:           -1,
	},
	{
		description: "a value larger than the array's largest value is not included",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         13,
		x:           -1,
	},
	{
		description: "nothing is included in an empty array",
		slice:       []int{},
		key:         1,
		x:           -1,
	},
}
