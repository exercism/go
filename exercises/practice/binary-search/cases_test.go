package binarysearch

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description string
	inputList   []int
	inputKey    int
	expectedKey int
	err         string
}{
	{
		description: "finds a value in an array with one element",
		inputList:   []int{6},
		inputKey:    6,
		expectedKey: 0,
		err:         "",
	},
	{
		description: "finds a value in the middle of an array",
		inputList:   []int{1, 3, 4, 6, 8, 9, 11},
		inputKey:    6,
		expectedKey: 3,
		err:         "",
	},
	{
		description: "finds a value at the beginning of an array",
		inputList:   []int{1, 3, 4, 6, 8, 9, 11},
		inputKey:    1,
		expectedKey: 0,
		err:         "",
	},
	{
		description: "finds a value at the end of an array",
		inputList:   []int{1, 3, 4, 6, 8, 9, 11},
		inputKey:    11,
		expectedKey: 6,
		err:         "",
	},
	{
		description: "finds a value in an array of odd length",
		inputList:   []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 634},
		inputKey:    144,
		expectedKey: 9,
		err:         "",
	},
	{
		description: "finds a value in an array of even length",
		inputList:   []int{1, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377},
		inputKey:    21,
		expectedKey: 5,
		err:         "",
	},
	{
		description: "identifies that a value is not included in the array",
		inputList:   []int{1, 3, 4, 6, 8, 9, 11},
		inputKey:    7,
		expectedKey: -1,
		err:         "value not in array",
	},
	{
		description: "a value smaller than the array's smallest value is not found",
		inputList:   []int{1, 3, 4, 6, 8, 9, 11},
		inputKey:    0,
		expectedKey: -1,
		err:         "value not in array",
	},
	{
		description: "a value larger than the array's largest value is not found",
		inputList:   []int{1, 3, 4, 6, 8, 9, 11},
		inputKey:    13,
		expectedKey: -1,
		err:         "value not in array",
	},
	{
		description: "nothing is found in an empty array",
		inputList:   []int{},
		inputKey:    1,
		expectedKey: -1,
		err:         "value not in array",
	},
	{
		description: "nothing is found when the left and right bounds cross",
		inputList:   []int{1, 2},
		inputKey:    0,
		expectedKey: -1,
		err:         "value not in array",
	},
}
