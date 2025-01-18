package transpose

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: daf84d6 bowling, transpose: conform array format to rest of file

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{
	{
		description: "empty string",
		input:       []string{},
		expected:    []string{},
	},
	{
		description: "two characters in a row",
		input:       []string{"A1"},
		expected:    []string{"A", "1"},
	},
	{
		description: "two characters in a column",
		input:       []string{"A", "1"},
		expected:    []string{"A1"},
	},
	{
		description: "simple",
		input:       []string{"ABC", "123"},
		expected:    []string{"A1", "B2", "C3"},
	},
	{
		description: "single line",
		input:       []string{"Single line."},
		expected:    []string{"S", "i", "n", "g", "l", "e", " ", "l", "i", "n", "e", "."},
	},
	{
		description: "first line longer than second line",
		input:       []string{"The fourth line.", "The fifth line."},
		expected:    []string{"TT", "hh", "ee", "  ", "ff", "oi", "uf", "rt", "th", "h ", " l", "li", "in", "ne", "e.", "."},
	},
	{
		description: "second line longer than first line",
		input:       []string{"The first line.", "The second line."},
		expected:    []string{"TT", "hh", "ee", "  ", "fs", "ie", "rc", "so", "tn", " d", "l ", "il", "ni", "en", ".e", " ."},
	},
	{
		description: "mixed line length",
		input:       []string{"The longest line.", "A long line.", "A longer line.", "A line."},
		expected:    []string{"TAAA", "h   ", "elll", " ooi", "lnnn", "ogge", "n e.", "glr", "ei ", "snl", "tei", " .n", "l e", "i .", "n", "e", "."},
	},
	{
		description: "square",
		input:       []string{"HEART", "EMBER", "ABUSE", "RESIN", "TREND"},
		expected:    []string{"HEART", "EMBER", "ABUSE", "RESIN", "TREND"},
	},
	{
		description: "rectangle",
		input:       []string{"FRACTURE", "OUTLINED", "BLOOMING", "SEPTETTE"},
		expected:    []string{"FOBS", "RULE", "ATOP", "CLOT", "TIME", "UNIT", "RENT", "EDGE"},
	},
	{
		description: "triangle",
		input:       []string{"T", "EE", "AAA", "SSSS", "EEEEE", "RRRRRR"},
		expected:    []string{"TEASER", " EASER", "  ASER", "   SER", "    ER", "     R"},
	},
	{
		description: "jagged triangle",
		input:       []string{"11", "2", "3333", "444", "555555", "66666"},
		expected:    []string{"123456", "1 3456", "  3456", "  3 56", "    56", "    5"},
	},
}
