package change

// Source: exercism/x-common
// Commit: 3d8b5b3 change: Fix canonical-data.json formatting

var testCases = []struct {
	description string
	coins       []int
	target      int
	expected    interface{}
}{
	{
		"single coin change",
		[]int{
			1,
			5,
			10,
			25,
			100,
		},
		25,
		[]interface{}{25},
	},
	{
		"multiple coin change",
		[]int{
			1,
			5,
			10,
			25,
			100,
		},
		15,
		[]interface{}{5, 10},
	},
	{
		"change with Lilliputian Coins",
		[]int{
			1,
			4,
			15,
			20,
			50,
		},
		23,
		[]interface{}{4, 4, 15},
	},
	{
		"change with Lower Elbonia Coins",
		[]int{
			1,
			5,
			10,
			21,
			25,
		},
		63,
		[]interface{}{21, 21, 21},
	},
	{
		"large target values",
		[]int{
			1,
			2,
			5,
			10,
			20,
			50,
			100,
		},
		999,
		[]interface{}{2, 2, 5, 20, 20, 50, 100, 100, 100, 100, 100, 100, 100, 100, 100},
	},
	{
		"possible change without unit coins available",
		[]int{
			2,
			5,
			10,
			20,
			50,
		},
		21,
		[]interface{}{2, 2, 2, 5, 10},
	},
	{
		"no coins make 0 change",
		[]int{
			1,
			5,
			10,
			21,
			25,
		},
		0,
		[]interface{}{},
	},
	{
		"error testing for change smaller than the smallest of coins",
		[]int{
			5,
			10,
		},
		3,
		-1,
	},
	{
		"error if no combination can add up to target",
		[]int{
			5,
			10,
		},
		94,
		-1,
	},
	{
		"cannot find negative change values",
		[]int{
			1,
			2,
			5,
		},
		-5,
		-1,
	},
}
