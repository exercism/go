package reversestring

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 25c2ae5 build(deps): bump fast-uri from 3.1.0 to 3.1.2 (#2653)

type reverseTestCase struct {
	description string
	input       string
	expected    string
}

var testCases = []reverseTestCase{
	{
		description: "an empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "a word",
		input:       "robot",
		expected:    "tobor",
	},
	{
		description: "a capitalized word",
		input:       "Ramen",
		expected:    "nemaR",
	},
	{
		description: "a sentence with punctuation",
		input:       "I'm hungry!",
		expected:    "!yrgnuh m'I",
	},
	{
		description: "a palindrome",
		input:       "racecar",
		expected:    "racecar",
	},
	{
		description: "an even-sized word",
		input:       "drawer",
		expected:    "reward",
	},
	{
		description: "wide characters",
		input:       "子猫",
		expected:    "猫子",
	},
}
