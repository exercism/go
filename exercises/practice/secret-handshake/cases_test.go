package secret

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 5fc501b Remove unneeded nesting (#1798)

type secretHandshakeTest struct {
	description string
	input       uint
	expected    []string
}

var tests = []secretHandshakeTest{
	{
		description: "wink for 1",
		input:       1,
		expected:    []string{"wink"},
	},
	{
		description: "double blink for 10",
		input:       2,
		expected:    []string{"double blink"},
	},
	{
		description: "close your eyes for 100",
		input:       4,
		expected:    []string{"close your eyes"},
	},
	{
		description: "jump for 1000",
		input:       8,
		expected:    []string{"jump"},
	},
	{
		description: "combine two actions",
		input:       3,
		expected:    []string{"wink", "double blink"},
	},
	{
		description: "reverse two actions",
		input:       19,
		expected:    []string{"double blink", "wink"},
	},
	{
		description: "reversing one action gives the same action",
		input:       24,
		expected:    []string{"jump"},
	},
	{
		description: "reversing no actions still gives no actions",
		input:       16,
		expected:    []string{},
	},
	{
		description: "all possible actions",
		input:       15,
		expected:    []string{"wink", "double blink", "close your eyes", "jump"},
	},
	{
		description: "reverse all possible actions",
		input:       31,
		expected:    []string{"jump", "close your eyes", "double blink", "wink"},
	},
	{
		description: "do nothing for zero",
		input:       0,
		expected:    []string{},
	},
}
