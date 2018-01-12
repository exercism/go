package change

// Source: exercism/problem-specifications
// Commit: 044d09a change: Apply new "input" policy
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description    string
	coins          []int
	target         int
	valid          bool  // true => no error, false => error expected
	expectedChange []int // when .valid == true, the expected change coins
}{
	{
		"single coin change",
		[]int{1, 5, 10, 25, 100},
		25,
		true,
		[]int{25},
	},
	{
		"multiple coin change",
		[]int{1, 5, 10, 25, 100},
		15,
		true,
		[]int{5, 10},
	},
	{
		"change with Lilliputian Coins",
		[]int{1, 4, 15, 20, 50},
		23,
		true,
		[]int{4, 4, 15},
	},
	{
		"change with Lower Elbonia Coins",
		[]int{1, 5, 10, 21, 25},
		63,
		true,
		[]int{21, 21, 21},
	},
	{
		"large target values",
		[]int{1, 2, 5, 10, 20, 50, 100},
		999,
		true,
		[]int{2, 2, 5, 20, 20, 50, 100, 100, 100, 100, 100, 100, 100, 100, 100},
	},
	{
		"possible change without unit coins available",
		[]int{2, 5, 10, 20, 50},
		21,
		true,
		[]int{2, 2, 2, 5, 10},
	},
	{
		"another possible change without unit coins available",
		[]int{4, 5},
		27,
		true,
		[]int{4, 4, 4, 5, 5, 5},
	},
	{
		"no coins make 0 change",
		[]int{1, 5, 10, 21, 25},
		0,
		true,
		[]int{},
	},
	{
		"error testing for change smaller than the smallest of coins",
		[]int{5, 10},
		3,
		false,
		[]int(nil),
	},
	{
		"error if no combination can add up to target",
		[]int{5, 10},
		94,
		false,
		[]int(nil),
	},
	{
		"cannot find negative change values",
		[]int{1, 2, 5},
		-5,
		false,
		[]int(nil),
	},
}
