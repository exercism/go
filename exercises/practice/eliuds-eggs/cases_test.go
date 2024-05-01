package eliudseggs

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 1a96480 Add `eliuds-eggs` exercise (#2382)

type testCase struct {
	description string
	input       int
	expected    int
}

var testCases = []testCase{
	{
		description: "0 eggs",
		input:       0,
		expected:    0,
	},
	{
		description: "1 egg",
		input:       16,
		expected:    1,
	},
	{
		description: "4 eggs",
		input:       89,
		expected:    4,
	},
	{
		description: "13 eggs",
		input:       2000000000,
		expected:    13,
	},
}
