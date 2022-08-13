package scrabble

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type scrabbleTest struct {
	description string
	input       string
	expected    int
}

var scrabbleScoreTests = []scrabbleTest{
	{
		description: "lowercase letter",
		input:       "a",
		expected:    1,
	},
	{
		description: "uppercase letter",
		input:       "A",
		expected:    1,
	},
	{
		description: "valuable letter",
		input:       "f",
		expected:    4,
	},
	{
		description: "short word",
		input:       "at",
		expected:    2,
	},
	{
		description: "short, valuable word",
		input:       "zoo",
		expected:    12,
	},
	{
		description: "medium word",
		input:       "street",
		expected:    6,
	},
	{
		description: "medium, valuable word",
		input:       "quirky",
		expected:    22,
	},
	{
		description: "long, mixed-case word",
		input:       "OxyphenButazone",
		expected:    41,
	},
	{
		description: "english-like word",
		input:       "pinata",
		expected:    8,
	},
	{
		description: "empty input",
		input:       "",
		expected:    0,
	},
	{
		description: "entire alphabet available",
		input:       "abcdefghijklmnopqrstuvwxyz",
		expected:    87,
	},
}
