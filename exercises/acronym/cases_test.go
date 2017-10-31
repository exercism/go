package acronym

// Source: exercism/problem-specifications
// Commit: cae7ae1 acronym: remove test case with mixed-case phrase (#788)
// Problem Specifications Version: 1.1.0

type acronymTest struct {
	input    string
	expected string
}

var stringTestCases = []acronymTest{
	{
		input:    "Portable Network Graphics",
		expected: "PNG",
	},
	{
		input:    "Ruby on Rails",
		expected: "ROR",
	},
	{
		input:    "First In, First Out",
		expected: "FIFO",
	},
	{
		input:    "PHP: Hypertext Preprocessor",
		expected: "PHP",
	},
	{
		input:    "GNU Image Manipulation Program",
		expected: "GIMP",
	},
	{
		input:    "Complementary metal-oxide semiconductor",
		expected: "CMOS",
	},
}
