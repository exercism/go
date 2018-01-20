package piglatin

// Source: exercism/problem-specifications
// Commit: d77de78 pig-latin: Apply new "input" policy
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "word beginning with a",
		input:       "apple",
		expected:    "appleay",
	},
	{
		description: "word beginning with e",
		input:       "ear",
		expected:    "earay",
	},
	{
		description: "word beginning with i",
		input:       "igloo",
		expected:    "iglooay",
	},
	{
		description: "word beginning with o",
		input:       "object",
		expected:    "objectay",
	},
	{
		description: "word beginning with u",
		input:       "under",
		expected:    "underay",
	},
	{
		description: "word beginning with a vowel and followed by a qu",
		input:       "equal",
		expected:    "equalay",
	},
	{
		description: "word beginning with p",
		input:       "pig",
		expected:    "igpay",
	},
	{
		description: "word beginning with k",
		input:       "koala",
		expected:    "oalakay",
	},
	{
		description: "word beginning with x",
		input:       "xenon",
		expected:    "enonxay",
	},
	{
		description: "word beginning with q without a following u",
		input:       "qat",
		expected:    "atqay",
	},
	{
		description: "word beginning with ch",
		input:       "chair",
		expected:    "airchay",
	},
	{
		description: "word beginning with qu",
		input:       "queen",
		expected:    "eenquay",
	},
	{
		description: "word beginning with qu and a preceding consonant",
		input:       "square",
		expected:    "aresquay",
	},
	{
		description: "word beginning with th",
		input:       "therapy",
		expected:    "erapythay",
	},
	{
		description: "word beginning with thr",
		input:       "thrush",
		expected:    "ushthray",
	},
	{
		description: "word beginning with sch",
		input:       "school",
		expected:    "oolschay",
	},
	{
		description: "word beginning with yt",
		input:       "yttria",
		expected:    "yttriaay",
	},
	{
		description: "word beginning with xr",
		input:       "xray",
		expected:    "xrayay",
	},
	{
		description: "y is treated like a consonant at the beginning of a word",
		input:       "yellow",
		expected:    "ellowyay",
	},
	{
		description: "y is treated like a vowel at the end of a consonant cluster",
		input:       "rhythm",
		expected:    "ythmrhay",
	},
	{
		description: "y as second letter in two letter word",
		input:       "my",
		expected:    "ymay",
	},
	{
		description: "a whole phrase",
		input:       "quick fast run",
		expected:    "ickquay astfay unray",
	},
}
