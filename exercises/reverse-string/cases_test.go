package reverse

// Source: exercism/problem-specifications
// Commit: ae82c90 More consistent reverse-string canonical data
// Problem Specifications Version: 1.0.1

var testCases = []struct {
	description string
	input       string
	expected    string
}{
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
