package nthprime

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 25c2ae5 build(deps): bump fast-uri from 3.1.0 to 3.1.2 (#2653)

type testCase struct {
	description string
	input       int
	expected    int
	err         string
}

var tests = []testCase{
	{
		description: "first prime",
		input:       1,
		expected:    2,
		err:         "",
	},
	{
		description: "second prime",
		input:       2,
		expected:    3,
		err:         "",
	},
	{
		description: "sixth prime",
		input:       6,
		expected:    13,
		err:         "",
	},
	{
		description: "big prime",
		input:       10001,
		expected:    104743,
		err:         "",
	},
	{
		description: "there is no zeroth prime",
		input:       0,
		expected:    0,
		err:         "there is no zeroth prime",
	},
}
