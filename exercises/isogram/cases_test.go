package isogram

// Source: exercism/problem-specifications
// Commit: 7cea153 isogram: add test case with same first and last character (#1308)
// Problem Specifications Version: 1.6.0

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "empty string",
		input:       "",
		expected:    true,
	},
	{
		description: "isogram with only lower case characters",
		input:       "isogram",
		expected:    true,
	},
	{
		description: "word with one duplicated character",
		input:       "eleven",
		expected:    false,
	},
	{
		description: "word with one duplicated character from the end of the alphabet",
		input:       "zzyzx",
		expected:    false,
	},
	{
		description: "longest reported english isogram",
		input:       "subdermatoglyphic",
		expected:    true,
	},
	{
		description: "word with duplicated character in mixed case",
		input:       "Alphabet",
		expected:    false,
	},
	{
		description: "word with duplicated character in mixed case, lowercase first",
		input:       "alphAbet",
		expected:    false,
	},
	{
		description: "hypothetical isogrammic word with hyphen",
		input:       "thumbscrew-japingly",
		expected:    true,
	},
	{
		description: "isogram with duplicated hyphen",
		input:       "six-year-old",
		expected:    true,
	},
	{
		description: "made-up name that is an isogram",
		input:       "Emily Jung Schwartzkopf",
		expected:    true,
	},
	{
		description: "duplicated character in the middle",
		input:       "accentor",
		expected:    false,
	},
	{
		description: "same first and last characters",
		input:       "angola",
		expected:    false,
	},
}
