package prime

// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

import "errors"

var tests = []struct {
	description string
	n           int
	p           int
	err         error
}{
	{
		"first prime",
		1,
		2,
		nil,
	},
	{
		"second prime",
		2,
		3,
		nil,
	},
	{
		"sixth prime",
		6,
		13,
		nil,
	},
	{
		"big prime",
		10001,
		104743,
		nil,
	},
	{
		"there is no zeroth prime",
		0,
		0,
		errors.New("input must be greater than 1"),
	},
}
