package twofer

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	input       string
	expected    string
}{
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
