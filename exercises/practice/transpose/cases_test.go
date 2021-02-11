package transpose

// Source: exercism/problem-specifications
// Commit: 92bc877 switch to new json-schema
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{
	{
		"empty string",
		[]string{},
		[]string{},
	},
	{
		"two characters in a row",
		[]string{
			"A1",
		},
		[]string{
			"A",
			"1",
		},
	},
	{
		"two characters in a column",
		[]string{
			"A",
			"1",
		},
		[]string{
			"A1",
		},
	},
	{
		"simple",
		[]string{
			"ABC",
			"123",
		},
		[]string{
			"A1",
			"B2",
			"C3",
		},
	},
	{
		"single line",
		[]string{
			"Single line.",
		},
		[]string{
			"S",
			"i",
			"n",
			"g",
			"l",
			"e",
			" ",
			"l",
			"i",
			"n",
			"e",
			".",
		},
	},
	{
		"first line longer than second line",
		[]string{
			"The fourth line.",
			"The fifth line.",
		},
		[]string{
			"TT",
			"hh",
			"ee",
			"  ",
			"ff",
			"oi",
			"uf",
			"rt",
			"th",
			"h ",
			" l",
			"li",
			"in",
			"ne",
			"e.",
			".",
		},
	},
	{
		"second line longer than first line",
		[]string{
			"The first line.",
			"The second line.",
		},
		[]string{
			"TT",
			"hh",
			"ee",
			"  ",
			"fs",
			"ie",
			"rc",
			"so",
			"tn",
			" d",
			"l ",
			"il",
			"ni",
			"en",
			".e",
			" .",
		},
	},
	{
		"mixed line length",
		[]string{
			"The longest line.",
			"A long line.",
			"A longer line.",
			"A line.",
		},
		[]string{
			"TAAA",
			"h   ",
			"elll",
			" ooi",
			"lnnn",
			"ogge",
			"n e.",
			"glr",
			"ei ",
			"snl",
			"tei",
			" .n",
			"l e",
			"i .",
			"n",
			"e",
			".",
		},
	},
	{
		"square",
		[]string{
			"HEART",
			"EMBER",
			"ABUSE",
			"RESIN",
			"TREND",
		},
		[]string{
			"HEART",
			"EMBER",
			"ABUSE",
			"RESIN",
			"TREND",
		},
	},
	{
		"rectangle",
		[]string{
			"FRACTURE",
			"OUTLINED",
			"BLOOMING",
			"SEPTETTE",
		},
		[]string{
			"FOBS",
			"RULE",
			"ATOP",
			"CLOT",
			"TIME",
			"UNIT",
			"RENT",
			"EDGE",
		},
	},
	{
		"triangle",
		[]string{
			"T",
			"EE",
			"AAA",
			"SSSS",
			"EEEEE",
			"RRRRRR",
		},
		[]string{
			"TEASER",
			" EASER",
			"  ASER",
			"   SER",
			"    ER",
			"     R",
		},
	},
}
