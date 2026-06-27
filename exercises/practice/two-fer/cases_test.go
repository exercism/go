package twofer

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 25c2ae5 build(deps): bump fast-uri from 3.1.0 to 3.1.2 (#2653)

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "no name given",
		input:       "",
		expected:    "One for you, one for me.",
	},
	{
		description: "a name given",
		input:       "Alice",
		expected:    "One for Alice, one for me.",
	},
	{
		description: "another name given",
		input:       "Bob",
		expected:    "One for Bob, one for me.",
	},
}
