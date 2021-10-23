package prime

// Source: exercism/problem-specifications
// Commit: 4a3ba76 nth-prime: Apply new "input" policy
// Problem Specifications Version: 2.1.0

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
