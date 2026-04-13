package rotationalcipher

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: b820099 Allow prettier to format more files (#1966)

type testCase struct {
	description string
	input       string
	shift       int
	expected    string
}

var testCases = []testCase{
	{
		description: "rotate a by 0, same output as input",
		input:       "a",
		shift:       0,
		expected:    "a",
	},
	{
		description: "rotate a by 1",
		input:       "a",
		shift:       1,
		expected:    "b",
	},
	{
		description: "rotate a by 26, same output as input",
		input:       "a",
		shift:       26,
		expected:    "a",
	},
	{
		description: "rotate m by 13",
		input:       "m",
		shift:       13,
		expected:    "z",
	},
	{
		description: "rotate n by 13 with wrap around alphabet",
		input:       "n",
		shift:       13,
		expected:    "a",
	},
	{
		description: "rotate capital letters",
		input:       "OMG",
		shift:       5,
		expected:    "TRL",
	},
	{
		description: "rotate spaces",
		input:       "O M G",
		shift:       5,
		expected:    "T R L",
	},
	{
		description: "rotate numbers",
		input:       "Testing 1 2 3 testing",
		shift:       4,
		expected:    "Xiwxmrk 1 2 3 xiwxmrk",
	},
	{
		description: "rotate punctuation",
		input:       "Let's eat, Grandma!",
		shift:       21,
		expected:    "Gzo'n zvo, Bmviyhv!",
	},
	{
		description: "rotate all letters",
		input:       "The quick brown fox jumps over the lazy dog.",
		shift:       13,
		expected:    "Gur dhvpx oebja sbk whzcf bire gur ynml qbt.",
	},
}
