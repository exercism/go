package reverse

// Source: exercism/problem-specifications
// Commit: 2f77985 reverse-string: apply "input" policy
// Problem Specifications Version: 1.1.0

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
}
