package dndcharacter

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 02209d7 Reimplement test case in DnD Character (#2338)

type modifierTestInput struct {
	Score int
}

var modifierTests = []struct {
	description string
	input       modifierTestInput
	expected    int
}{

	{
		description: "ability modifier for score 3 is -4",
		input: modifierTestInput{
			Score: 3,
		},
		expected: -4,
	},

	{
		description: "ability modifier for score 4 is -3",
		input: modifierTestInput{
			Score: 4,
		},
		expected: -3,
	},

	{
		description: "ability modifier for score 5 is -3",
		input: modifierTestInput{
			Score: 5,
		},
		expected: -3,
	},

	{
		description: "ability modifier for score 6 is -2",
		input: modifierTestInput{
			Score: 6,
		},
		expected: -2,
	},

	{
		description: "ability modifier for score 7 is -2",
		input: modifierTestInput{
			Score: 7,
		},
		expected: -2,
	},

	{
		description: "ability modifier for score 8 is -1",
		input: modifierTestInput{
			Score: 8,
		},
		expected: -1,
	},

	{
		description: "ability modifier for score 9 is -1",
		input: modifierTestInput{
			Score: 9,
		},
		expected: -1,
	},

	{
		description: "ability modifier for score 10 is 0",
		input: modifierTestInput{
			Score: 10,
		},
		expected: 0,
	},

	{
		description: "ability modifier for score 11 is 0",
		input: modifierTestInput{
			Score: 11,
		},
		expected: 0,
	},

	{
		description: "ability modifier for score 12 is +1",
		input: modifierTestInput{
			Score: 12,
		},
		expected: 1,
	},

	{
		description: "ability modifier for score 13 is +1",
		input: modifierTestInput{
			Score: 13,
		},
		expected: 1,
	},

	{
		description: "ability modifier for score 14 is +2",
		input: modifierTestInput{
			Score: 14,
		},
		expected: 2,
	},

	{
		description: "ability modifier for score 15 is +2",
		input: modifierTestInput{
			Score: 15,
		},
		expected: 2,
	},

	{
		description: "ability modifier for score 16 is +3",
		input: modifierTestInput{
			Score: 16,
		},
		expected: 3,
	},

	{
		description: "ability modifier for score 17 is +3",
		input: modifierTestInput{
			Score: 17,
		},
		expected: 3,
	},

	{
		description: "ability modifier for score 18 is +4",
		input: modifierTestInput{
			Score: 18,
		},
		expected: 4,
	},
}
