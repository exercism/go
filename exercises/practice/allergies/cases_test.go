package allergies

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 988be7b allergies: format (#1978)

type allergicToInput struct {
	allergen string
	score    uint
}

type allergicToTestCase struct {
	description string
	input       allergicToInput
	expected    bool
}

var allergicToTests = []allergicToTestCase{
	{
		description: "not allergic to anything: check eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to eggs: check eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    1,
		},
		expected: true,
	},
	{
		description: "allergic to eggs and something else: check eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    3,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not eggs: check eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    2,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to peanuts: check peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    2,
		},
		expected: true,
	},
	{
		description: "allergic to peanuts and something else: check peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    7,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not peanuts: check peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    5,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to shellfish: check shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    4,
		},
		expected: true,
	},
	{
		description: "allergic to shellfish and something else: check shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    14,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not shellfish: check shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    10,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to strawberries: check strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    8,
		},
		expected: true,
	},
	{
		description: "allergic to strawberries and something else: check strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    28,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not strawberries: check strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    20,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to tomatoes: check tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    16,
		},
		expected: true,
	},
	{
		description: "allergic to tomatoes and something else: check tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    56,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not tomatoes: check tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    40,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to chocolate: check chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    32,
		},
		expected: true,
	},
	{
		description: "allergic to chocolate and something else: check chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    112,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not chocolate: check chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    80,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to pollen: check pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    64,
		},
		expected: true,
	},
	{
		description: "allergic to pollen and something else: check pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    224,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not pollen: check pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    160,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything: check cats",
		input: allergicToInput{
			allergen: "cats",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to cats: check cats",
		input: allergicToInput{
			allergen: "cats",
			score:    128,
		},
		expected: true,
	},
	{
		description: "allergic to cats and something else: check cats",
		input: allergicToInput{
			allergen: "cats",
			score:    192,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not cats: check cats",
		input: allergicToInput{
			allergen: "cats",
			score:    64,
		},
		expected: false,
	},
	{
		description: "allergic to everything: check cats",
		input: allergicToInput{
			allergen: "cats",
			score:    255,
		},
		expected: true,
	},
}

type listTestCase struct {
	description string
	score       uint
	expected    []string
}

var listTests = []listTestCase{
	{
		description: "no allergies",
		score:       0,
		expected:    []string{},
	},
	{
		description: "just eggs",
		score:       1,
		expected:    []string{"eggs"},
	},
	{
		description: "just peanuts",
		score:       2,
		expected:    []string{"peanuts"},
	},
	{
		description: "just strawberries",
		score:       8,
		expected:    []string{"strawberries"},
	},
	{
		description: "eggs and peanuts",
		score:       3,
		expected:    []string{"eggs", "peanuts"},
	},
	{
		description: "more than eggs but not peanuts",
		score:       5,
		expected:    []string{"eggs", "shellfish"},
	},
	{
		description: "lots of stuff",
		score:       248,
		expected:    []string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"},
	},
	{
		description: "everything",
		score:       255,
		expected:    []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"},
	},
	{
		description: "no allergen score parts",
		score:       509,
		expected:    []string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"},
	},
	{
		description: "no allergen score parts without highest valid score",
		score:       257,
		expected:    []string{"eggs"},
	},
}
