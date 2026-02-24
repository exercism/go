package romannumerals

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 438990d roman-numerals: add test with longest output (#2414)

type romanNumeralTest struct {
	description string
	input       int
	expected    string
}

var validRomanNumeralTests = []romanNumeralTest{
	{
		description: "1 is I",
		input:       1,
		expected:    "I",
	},
	{
		description: "2 is II",
		input:       2,
		expected:    "II",
	},
	{
		description: "3 is III",
		input:       3,
		expected:    "III",
	},
	{
		description: "4 is IV",
		input:       4,
		expected:    "IV",
	},
	{
		description: "5 is V",
		input:       5,
		expected:    "V",
	},
	{
		description: "6 is VI",
		input:       6,
		expected:    "VI",
	},
	{
		description: "9 is IX",
		input:       9,
		expected:    "IX",
	},
	{
		description: "16 is XVI",
		input:       16,
		expected:    "XVI",
	},
	{
		description: "27 is XXVII",
		input:       27,
		expected:    "XXVII",
	},
	{
		description: "48 is XLVIII",
		input:       48,
		expected:    "XLVIII",
	},
	{
		description: "49 is XLIX",
		input:       49,
		expected:    "XLIX",
	},
	{
		description: "59 is LIX",
		input:       59,
		expected:    "LIX",
	},
	{
		description: "66 is LXVI",
		input:       66,
		expected:    "LXVI",
	},
	{
		description: "93 is XCIII",
		input:       93,
		expected:    "XCIII",
	},
	{
		description: "141 is CXLI",
		input:       141,
		expected:    "CXLI",
	},
	{
		description: "163 is CLXIII",
		input:       163,
		expected:    "CLXIII",
	},
	{
		description: "166 is CLXVI",
		input:       166,
		expected:    "CLXVI",
	},
	{
		description: "402 is CDII",
		input:       402,
		expected:    "CDII",
	},
	{
		description: "575 is DLXXV",
		input:       575,
		expected:    "DLXXV",
	},
	{
		description: "666 is DCLXVI",
		input:       666,
		expected:    "DCLXVI",
	},
	{
		description: "911 is CMXI",
		input:       911,
		expected:    "CMXI",
	},
	{
		description: "1024 is MXXIV",
		input:       1024,
		expected:    "MXXIV",
	},
	{
		description: "1666 is MDCLXVI",
		input:       1666,
		expected:    "MDCLXVI",
	},
	{
		description: "3000 is MMM",
		input:       3000,
		expected:    "MMM",
	},
	{
		description: "3001 is MMMI",
		input:       3001,
		expected:    "MMMI",
	},
	{
		description: "3888 is MMMDCCCLXXXVIII",
		input:       3888,
		expected:    "MMMDCCCLXXXVIII",
	},
	{
		description: "3999 is MMMCMXCIX",
		input:       3999,
		expected:    "MMMCMXCIX",
	},
}
