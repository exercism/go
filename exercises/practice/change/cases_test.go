package change

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description    string
	coins          []int
	target         int
	valid          bool  // true => no error, false => error expected
	expectedChange []int // when .valid == true, the expected change coins
}{
	{
		description:    "change for 1 cent",
		coins:          []int{1, 5, 10, 25},
		target:         1,
		valid:          true,
		expectedChange: []int{1},
	},
	{
		description:    "single coin change",
		coins:          []int{1, 5, 10, 25, 100},
		target:         25,
		valid:          true,
		expectedChange: []int{25},
	},
	{
		description:    "multiple coin change",
		coins:          []int{1, 5, 10, 25, 100},
		target:         15,
		valid:          true,
		expectedChange: []int{5, 10},
	},
	{
		description:    "change with Lilliputian Coins",
		coins:          []int{1, 4, 15, 20, 50},
		target:         23,
		valid:          true,
		expectedChange: []int{4, 4, 15},
	},
	{
		description:    "change with Lower Elbonia Coins",
		coins:          []int{1, 5, 10, 21, 25},
		target:         63,
		valid:          true,
		expectedChange: []int{21, 21, 21},
	},
	{
		description:    "large target values",
		coins:          []int{1, 2, 5, 10, 20, 50, 100},
		target:         999,
		valid:          true,
		expectedChange: []int{2, 2, 5, 20, 20, 50, 100, 100, 100, 100, 100, 100, 100, 100, 100},
	},
	{
		description:    "possible change without unit coins available",
		coins:          []int{2, 5, 10, 20, 50},
		target:         21,
		valid:          true,
		expectedChange: []int{2, 2, 2, 5, 10},
	},
	{
		description:    "another possible change without unit coins available",
		coins:          []int{4, 5},
		target:         27,
		valid:          true,
		expectedChange: []int{4, 4, 4, 5, 5, 5},
	},
	{
		description:    "no coins make 0 change",
		coins:          []int{1, 5, 10, 21, 25},
		target:         0,
		valid:          true,
		expectedChange: []int{},
	},
	{
		description:    "error testing for change smaller than the smallest of coins",
		coins:          []int{5, 10},
		target:         3,
		valid:          false,
		expectedChange: []int(nil),
	},
	{
		description:    "error if no combination can add up to target",
		coins:          []int{5, 10},
		target:         94,
		valid:          false,
		expectedChange: []int(nil),
	},
	{
		description:    "cannot find negative change values",
		coins:          []int{1, 2, 5},
		target:         -5,
		valid:          false,
		expectedChange: []int(nil),
	},
}
