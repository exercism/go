package pangram

// Source: exercism/problem-specifications
// Commit: 2c020bc Fix typo
// Problem Specifications Version: 1.4.1

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "sentence empty",
		input:       "",
		expected:    false,
	},
	{
		description: "recognizes a perfect lower case pangram",
		input:       "abcdefghijklmnopqrstuvwxyz",
		expected:    true,
	},
	{
		description: "pangram with only lower case",
		input:       "the quick brown fox jumps over the lazy dog",
		expected:    true,
	},
	{
		description: "missing character 'x'",
		input:       "a quick movement of the enemy will jeopardize five gunboats",
		expected:    false,
	},
	{
		description: "another missing character, e.g. 'h'",
		input:       "five boxing wizards jump quickly at it",
		expected:    false,
	},
	{
		description: "pangram with underscores",
		input:       "the_quick_brown_fox_jumps_over_the_lazy_dog",
		expected:    true,
	},
	{
		description: "pangram with numbers",
		input:       "the 1 quick brown fox jumps over the 2 lazy dogs",
		expected:    true,
	},
	{
		description: "missing letters replaced by numbers",
		input:       "7h3 qu1ck brown fox jumps ov3r 7h3 lazy dog",
		expected:    false,
	},
	{
		description: "pangram with mixed case and punctuation",
		input:       "\"Five quacking Zephyrs jolt my wax bed.\"",
		expected:    true,
	},
	{
		description: "upper and lower case versions of the same character should not be counted separately",
		input:       "the quick brown fox jumps over with lazy FX",
		expected:    false,
	},
}
