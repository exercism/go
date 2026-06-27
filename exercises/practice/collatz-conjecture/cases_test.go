package collatzconjecture

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 25c2ae5 build(deps): bump fast-uri from 3.1.0 to 3.1.2 (#2653)

type testCase struct {
	description string
	input       int
	expectError bool
	expected    int
}

var testCases = []testCase{
	{
		description: "zero steps for one",
		input:       1,
		expectError: false,
		expected:    0,
	},
	{
		description: "divide if even",
		input:       16,
		expectError: false,
		expected:    4,
	},
	{
		description: "even and odd steps",
		input:       12,
		expectError: false,
		expected:    9,
	},
	{
		description: "large number of even and odd steps",
		input:       1000000,
		expectError: false,
		expected:    152,
	},
	{
		description: "zero is an error",
		input:       0,
		expectError: true,
		expected:    0,
	},
	{
		description: "negative value is an error",
		input:       -15,
		expectError: true,
		expected:    0,
	},
}
