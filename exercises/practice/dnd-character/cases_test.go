package dndcharacter

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 02209d7 Reimplement test case in DnD Character (#2338)

var modifierTests = []struct {
	description string
	input       struct{ Score int }
	expected    int
}{

	{
		description: "ability modifier for score 3 is -4",
		input: struct{ Score int }{
			Score: 3,
		},
		expected: -4,
	},

	{
		description: "ability modifier for score 4 is -3",
		input: struct{ Score int }{
			Score: 4,
		},
		expected: -3,
	},

	{
		description: "ability modifier for score 5 is -3",
		input: struct{ Score int }{
			Score: 5,
		},
		expected: -3,
	},

	{
		description: "ability modifier for score 6 is -2",
		input: struct{ Score int }{
			Score: 6,
		},
		expected: -2,
	},

	{
		description: "ability modifier for score 7 is -2",
		input: struct{ Score int }{
			Score: 7,
		},
		expected: -2,
	},

	{
		description: "ability modifier for score 8 is -1",
		input: struct{ Score int }{
			Score: 8,
		},
		expected: -1,
	},

	{
		description: "ability modifier for score 9 is -1",
		input: struct{ Score int }{
			Score: 9,
		},
		expected: -1,
	},

	{
		description: "ability modifier for score 10 is 0",
		input: struct{ Score int }{
			Score: 10,
		},
		expected: 0,
	},

	{
		description: "ability modifier for score 11 is 0",
		input: struct{ Score int }{
			Score: 11,
		},
		expected: 0,
	},

	{
		description: "ability modifier for score 12 is +1",
		input: struct{ Score int }{
			Score: 12,
		},
		expected: 1,
	},

	{
		description: "ability modifier for score 13 is +1",
		input: struct{ Score int }{
			Score: 13,
		},
		expected: 1,
	},

	{
		description: "ability modifier for score 14 is +2",
		input: struct{ Score int }{
			Score: 14,
		},
		expected: 2,
	},

	{
		description: "ability modifier for score 15 is +2",
		input: struct{ Score int }{
			Score: 15,
		},
		expected: 2,
	},

	{
		description: "ability modifier for score 16 is +3",
		input: struct{ Score int }{
			Score: 16,
		},
		expected: 3,
	},

	{
		description: "ability modifier for score 17 is +3",
		input: struct{ Score int }{
			Score: 17,
		},
		expected: 3,
	},

	{
		description: "ability modifier for score 18 is +4",
		input: struct{ Score int }{
			Score: 18,
		},
		expected: 4,
	},
}

var abilityTests = []struct {
	description string
	input       struct{}
	expected    string
}{

	{
		description: "random ability is within range",
		input:       struct{}{},
		expected:    "score >= 3 && score <= 18",
	},
}

var characterTests = []struct {
	description string
	input       struct{}
	expected    struct {
		Strength     string
		Dexterity    string
		Constitution string
		Intelligence string
		Wisdom       string
		Charisma     string
		Hitpoints    string
	}
}{

	{
		description: "random character is valid",
		input:       struct{}{},
		expected: struct {
			Strength     string
			Dexterity    string
			Constitution string
			Intelligence string
			Wisdom       string
			Charisma     string
			Hitpoints    string
		}{
			Strength:     "strength >= 3 && strength <= 18",
			Dexterity:    "dexterity >= 3 && dexterity <= 18",
			Constitution: "constitution >= 3 && constitution <= 18",
			Intelligence: "intelligence >= 3 && intelligence <= 18",
			Wisdom:       "wisdom >= 3 && wisdom <= 18",
			Charisma:     "charisma >= 3 && charisma <= 18",
			Hitpoints:    "hitpoints == 10 + modifier(constitution)",
		},
	},

	{
		description: "each ability is only calculated once",
		input:       struct{}{},
		expected: struct {
			Strength     string
			Dexterity    string
			Constitution string
			Intelligence string
			Wisdom       string
			Charisma     string
			Hitpoints    string
		}{
			Strength:     "strength == strength",
			Dexterity:    "dexterity == dexterity",
			Constitution: "constitution == constitution",
			Intelligence: "intelligence == intelligence",
			Wisdom:       "wisdom == wisdom",
			Charisma:     "charisma == charisma",
			Hitpoints:    "",
		},
	},
}

var strengthTests = []struct {
	description string
	input       struct{}
	expected    string
}{}
