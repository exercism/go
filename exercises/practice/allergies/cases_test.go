package allergies

// Source: exercism/problem-specifications
// Commit: 988be7b allergies: format (#1978)

// allergicTo
type allergicToInput struct {
	allergen string
	score    uint
}

var allergicToTests = []struct {
	description string
	input       allergicToInput
	expected    bool
}{
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "eggs",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    1,
		},
		expected: true,
	},
	{
		description: "allergic to eggs and something else",
		input: allergicToInput{
			allergen: "eggs",
			score:    3,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not eggs",
		input: allergicToInput{
			allergen: "eggs",
			score:    2,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "eggs",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "peanuts",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    2,
		},
		expected: true,
	},
	{
		description: "allergic to peanuts and something else",
		input: allergicToInput{
			allergen: "peanuts",
			score:    7,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not peanuts",
		input: allergicToInput{
			allergen: "peanuts",
			score:    5,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "peanuts",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "shellfish",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    4,
		},
		expected: true,
	},
	{
		description: "allergic to shellfish and something else",
		input: allergicToInput{
			allergen: "shellfish",
			score:    14,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not shellfish",
		input: allergicToInput{
			allergen: "shellfish",
			score:    10,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "shellfish",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "strawberries",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    8,
		},
		expected: true,
	},
	{
		description: "allergic to strawberries and something else",
		input: allergicToInput{
			allergen: "strawberries",
			score:    28,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not strawberries",
		input: allergicToInput{
			allergen: "strawberries",
			score:    20,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "strawberries",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    16,
		},
		expected: true,
	},
	{
		description: "allergic to tomatoes and something else",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    56,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not tomatoes",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    40,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "tomatoes",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "chocolate",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    32,
		},
		expected: true,
	},
	{
		description: "allergic to chocolate and something else",
		input: allergicToInput{
			allergen: "chocolate",
			score:    112,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not chocolate",
		input: allergicToInput{
			allergen: "chocolate",
			score:    80,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "chocolate",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "pollen",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    64,
		},
		expected: true,
	},
	{
		description: "allergic to pollen and something else",
		input: allergicToInput{
			allergen: "pollen",
			score:    224,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not pollen",
		input: allergicToInput{
			allergen: "pollen",
			score:    160,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "pollen",
			score:    255,
		},
		expected: true,
	},
	{
		description: "not allergic to anything",
		input: allergicToInput{
			allergen: "cats",
			score:    0,
		},
		expected: false,
	},
	{
		description: "allergic only to cats",
		input: allergicToInput{
			allergen: "cats",
			score:    128,
		},
		expected: true,
	},
	{
		description: "allergic to cats and something else",
		input: allergicToInput{
			allergen: "cats",
			score:    192,
		},
		expected: true,
	},
	{
		description: "allergic to something, but not cats",
		input: allergicToInput{
			allergen: "cats",
			score:    64,
		},
		expected: false,
	},
	{
		description: "allergic to everything",
		input: allergicToInput{
			allergen: "cats",
			score:    255,
		},
		expected: true,
	},
}

// list
var listTests = []struct {
	description string
	score       uint
	expected    []string
}{
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
