package acronym

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 5fc501b Remove unneeded nesting (#1798)

type acronymTest struct {
	description string
	input       string
	expected    string
}

var testCases = []acronymTest{
	{
		description: "basic",
		input:       "Portable Network Graphics",
		expected:    "PNG",
	},
	{
		description: "lowercase words",
		input:       "Ruby on Rails",
		expected:    "ROR",
	},
	{
		description: "punctuation",
		input:       "First In, First Out",
		expected:    "FIFO",
	},
	{
		description: "all caps word",
		input:       "GNU Image Manipulation Program",
		expected:    "GIMP",
	},
	{
		description: "punctuation without whitespace",
		input:       "Complementary metal-oxide semiconductor",
		expected:    "CMOS",
	},
	{
		description: "very long abbreviation",
		input:       "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me",
		expected:    "ROTFLSHTMDCOALM",
	},
	{
		description: "consecutive delimiters",
		input:       "Something - I made up from thin air",
		expected:    "SIMUFTA",
	},
	{
		description: "apostrophes",
		input:       "Halley's Comet",
		expected:    "HC",
	},
	{
		description: "underscore emphasis",
		input:       "The Road _Not_ Taken",
		expected:    "TRNT",
	},
}
