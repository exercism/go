package prime

// Source: exercism/problem-specifications
// Commit: e82cbcd nth-prime: Use object instead of bool for invalid value (#969)
// Problem Specifications Version: 2.0.0

var tests = []struct {
	description string
	n           int
	p           int
	ok          bool
}{
	{
		"first prime",
		1,
		2,
		true,
	},
	{
		"second prime",
		2,
		3,
		true,
	},
	{
		"sixth prime",
		6,
		13,
		true,
	},
	{
		"big prime",
		10001,
		104743,
		true,
	},
	{
		"there is no zeroth prime",
		0,
		0,
		false,
	},
}
