package microblog

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 5fc501b Remove unneeded nesting (#1798)

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "English language short",
		input:       "Hi",
		expected:    "Hi",
	},
	{
		description: "English language long",
		input:       "Hello there",
		expected:    "Hello",
	},
	{
		description: "German language short (broth)",
		input:       "brühe",
		expected:    "brühe",
	},
	{
		description: "German language long (bear carpet → beards)",
		input:       "Bärteppich",
		expected:    "Bärte",
	},
	{
		description: "Bulgarian language short (good)",
		input:       "Добър",
		expected:    "Добър",
	},
	{
		description: "Greek language short (health)",
		input:       "υγειά",
		expected:    "υγειά",
	},
	{
		description: "Maths short",
		input:       "a=πr²",
		expected:    "a=πr²",
	},
	{
		description: "Maths long",
		input:       "∅⊊ℕ⊊ℤ⊊ℚ⊊ℝ⊊ℂ",
		expected:    "∅⊊ℕ⊊ℤ",
	},
	{
		description: "English and emoji short",
		input:       "Fly 🛫",
		expected:    "Fly 🛫",
	},
	{
		description: "Emoji short",
		input:       "💇",
		expected:    "💇",
	},
	{
		description: "Emoji long",
		input:       "❄🌡🤧🤒🏥🕰😀",
		expected:    "❄🌡🤧🤒🏥",
	},
	{
		description: "Royal Flush?",
		input:       "🃎🂸🃅🃋🃍🃁🃊",
		expected:    "🃎🂸🃅🃋🃍",
	},
}
