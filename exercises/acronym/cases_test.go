package acronym

// Source: exercism/problem-specifications
// Commit: 2e46654b Acronym: New test case for apostrophe (#1377)
// Problem Specifications Version: 1.6.0

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
		input:    "GNU Image Manipulation Program",
		expected: "GIMP",
	},
	{
		input:    "Complementary metal-oxide semiconductor",
		expected: "CMOS",
	},
	{
		input:    "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me",
		expected: "ROTFLSHTMDCOALM",
	},
	{
		input:    "Something - I made up from thin air",
		expected: "SIMUFTA",
	},
	{
		input:    "Halley's Comet",
		expected: "HC",
	},
}
