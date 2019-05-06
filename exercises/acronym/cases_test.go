package acronym

// Source: exercism/problem-specifications
// Commit: cacf1f1 Acronym: add underscore test case (#1436)
// Problem Specifications Version: 1.7.0

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
	{
		input:    "The Road _Not_ Taken",
		expected: "TRNT",
	},
}
