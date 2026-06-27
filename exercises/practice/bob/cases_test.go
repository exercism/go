package bob

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 0bdd27d [bob] Sort test cases (#2662)

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "asking a question",
		input:       "Does this cryogenic chamber make me look fat?",
		expected:    "Sure.",
	},
	{
		description: "shouting",
		input:       "WATCH OUT!",
		expected:    "Whoa, chill out!",
	},
	{
		description: "forceful question",
		input:       "WHAT'S GOING ON?",
		expected:    "Calm down, I know what I'm doing!",
	},
	{
		description: "silence",
		input:       "",
		expected:    "Fine. Be that way!",
	},
	{
		description: "stating something",
		input:       "Tom-ay-to, tom-aaaah-to.",
		expected:    "Whatever.",
	},
	{
		description: "asking a numeric question",
		input:       "You are, what, like 15?",
		expected:    "Sure.",
	},
	{
		description: "asking gibberish",
		input:       "fffbbcbeab?",
		expected:    "Sure.",
	},
	{
		description: "question with no letters",
		input:       "4?",
		expected:    "Sure.",
	},
	{
		description: "non-letters with question",
		input:       ":) ?",
		expected:    "Sure.",
	},
	{
		description: "prattling on",
		input:       "Wait! Hang on. Are you going to be OK?",
		expected:    "Sure.",
	},
	{
		description: "ending with whitespace",
		input:       "Okay if like my  spacebar  quite a bit?   ",
		expected:    "Sure.",
	},
	{
		description: "multiple line question",
		input:       "\nDoes this cryogenic chamber make\n me look fat?",
		expected:    "Sure.",
	},
	{
		description: "shouting gibberish",
		input:       "FCECDFCAAB",
		expected:    "Whoa, chill out!",
	},
	{
		description: "shouting a statement containing a question mark",
		input:       "DO LIONS EAT PEOPLE? AHHHHH.",
		expected:    "Whoa, chill out!",
	},
	{
		description: "shouting numbers",
		input:       "1, 2, 3 GO!",
		expected:    "Whoa, chill out!",
	},
	{
		description: "shouting with special characters",
		input:       "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
		expected:    "Whoa, chill out!",
	},
	{
		description: "shouting with no exclamation mark",
		input:       "I HATE THE DENTIST",
		expected:    "Whoa, chill out!",
	},
	{
		description: "prolonged silence",
		input:       "          ",
		expected:    "Fine. Be that way!",
	},
	{
		description: "alternate silence",
		input:       "\t\t\t\t\t\t\t\t\t\t",
		expected:    "Fine. Be that way!",
	},
	{
		description: "other whitespace",
		input:       "\n\r \t",
		expected:    "Fine. Be that way!",
	},
	{
		description: "talking forcefully",
		input:       "Hi there!",
		expected:    "Whatever.",
	},
	{
		description: "using acronyms in regular speech",
		input:       "It's OK if you don't want to go work for NASA.",
		expected:    "Whatever.",
	},
	{
		description: "no letters",
		input:       "1, 2, 3",
		expected:    "Whatever.",
	},
	{
		description: "statement containing question mark",
		input:       "Ending with ? means a question.",
		expected:    "Whatever.",
	},
	{
		description: "starting with whitespace",
		input:       "         hmmmmmmm...",
		expected:    "Whatever.",
	},
	{
		description: "non-question ending with whitespace",
		input:       "This is a statement ending with whitespace      ",
		expected:    "Whatever.",
	},
}
