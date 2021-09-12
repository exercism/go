package say

// Source: exercism/problem-specifications
// Commit: a0cee46 say: add mention of 'error' to comment
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	input       int64
	expected    string
	expectError bool
}{
	{
		description: "zero",
		input:       0,
		expected:    "zero",
	},
	{
		description: "one",
		input:       1,
		expected:    "one",
	},
	{
		description: "fourteen",
		input:       14,
		expected:    "fourteen",
	},
	{
		description: "twenty",
		input:       20,
		expected:    "twenty",
	},
	{
		description: "twenty-two",
		input:       22,
		expected:    "twenty-two",
	},
	{
		description: "one hundred",
		input:       100,
		expected:    "one hundred",
	},
	{
		description: "one hundred twenty-three",
		input:       123,
		expected:    "one hundred twenty-three",
	},
	{
		description: "one thousand",
		input:       1000,
		expected:    "one thousand",
	},
	{
		description: "one thousand two hundred thirty-four",
		input:       1234,
		expected:    "one thousand two hundred thirty-four",
	},
	{
		description: "one million",
		input:       1000000,
		expected:    "one million",
	},
	{
		description: "one million two thousand three hundred forty-five",
		input:       1002345,
		expected:    "one million two thousand three hundred forty-five",
	},
	{
		description: "one billion",
		input:       1000000000,
		expected:    "one billion",
	},
	{
		description: "a big number",
		input:       987654321123,
		expected:    "nine hundred eighty-seven billion six hundred fifty-four million three hundred twenty-one thousand one hundred twenty-three",
	},
	{
		description: "numbers below zero are out of range",
		input:       -1,
		expectError: true,
	},
	{
		description: "numbers above 999,999,999,999 are out of range",
		input:       1000000000000,
		expectError: true,
	},
}
