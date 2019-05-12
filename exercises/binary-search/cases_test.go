package binarysearch

// Source: exercism/problem-specifications
// Commit: bfb218f binary-search: test description tweak
// Problem Specifications Version: 1.3.0

var testCases = []struct {
	description string
	slice       []int
	key         int
	x           int
	err         string
}{
	{
		description: "finds a value in an array with one element",
		slice:       []int{6},
		key:         6,
		x:           0,
		err:         "",
	},
	{
		description: "finds a value in the middle of an array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         6,
		x:           3,
		err:         "",
	},
	{
		description: "finds a value at the beginning of an array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         1,
		x:           0,
		err:         "",
	},
	{
		description: "finds a value at the end of an array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         11,
		x:           6,
		err:         "",
	},
	{
		description: "finds a value in an array of odd length",
		slice:       []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 634},
		key:         144,
		x:           9,
		err:         "",
	},
	{
		description: "finds a value in an array of even length",
		slice:       []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377},
		key:         21,
		x:           5,
		err:         "",
	},
	{
		description: "identifies that a value is not included in the array",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         7,
		x:           -1,
		err:         "value not in array",
	},
	{
		description: "a value smaller than the array's smallest value is not found",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         0,
		x:           -1,
		err:         "value not in array",
	},
	{
		description: "a value larger than the array's largest value is not found",
		slice:       []int{1, 3, 4, 6, 8, 9, 11},
		key:         13,
		x:           -1,
		err:         "value not in array",
	},
	{
		description: "nothing is found in an empty array",
		slice:       []int{},
		key:         1,
		x:           -1,
		err:         "value not in array",
	},
	{
		description: "nothing is found when the left and right bounds cross",
		slice:       []int{1, 2},
		key:         0,
		x:           -1,
		err:         "value not in array",
	},
}
