package change

// Source: exercism/x-common
// Commit: cda8f98 Create new exercises structure

var testCases = []struct {
	description string
	coins       []int
	target      int
	expected    []int
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
		[]int{
			25,
		},
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
		[]int{
			5,
			10,
		},
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
		[]int{
			4,
			4,
			15,
		},
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
		[]int{
			21,
			21,
			21,
		},
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
		[]int{
			2,
			2,
			5,
			20,
			20,
			50,
			100,
			100,
			100,
			100,
			100,
			100,
			100,
			100,
			100,
		},
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
		[]int{},
	},
	{
		"error testing for change smaller than the smallest of coins",
		[]int{
			5,
			10,
		},
		3,
		[]int{
			-1,
		},
	},
	{
		"cannot find negative change values",
		[]int{
			1,
			2,
			5,
		},
		-5,
		[]int{
			-1,
		},
	},
}
