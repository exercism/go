package knapsack

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 6c1d8e2 Knapsack: Fix empty list of items (#2350)

type maximumValueCaseInput struct {
	MaximumWeight int
	Items         []Item
}

type maximumValueTest struct {
	description string
	input       maximumValueCaseInput
	expected    int
}

var maximumValueTests = []maximumValueTest{
	{
		description: "no items",
		input: maximumValueCaseInput{
			MaximumWeight: 100,
			Items:         []Item{},
		},
		expected: 0,
	},

	{
		description: "one item, too heavy",
		input: maximumValueCaseInput{
			MaximumWeight: 10,
			Items: []Item{
				{
					Weight: 100,
					Value:  1,
				},
			},
		},
		expected: 0,
	},

	{
		description: "five items (cannot be greedy by weight)",
		input: maximumValueCaseInput{
			MaximumWeight: 10,
			Items: []Item{
				{
					Weight: 2,
					Value:  5,
				},

				{
					Weight: 2,
					Value:  5,
				},

				{
					Weight: 2,
					Value:  5,
				},

				{
					Weight: 2,
					Value:  5,
				},

				{
					Weight: 10,
					Value:  21,
				},
			},
		},
		expected: 21,
	},

	{
		description: "five items (cannot be greedy by value)",
		input: maximumValueCaseInput{
			MaximumWeight: 10,
			Items: []Item{
				{
					Weight: 2,
					Value:  20,
				},

				{
					Weight: 2,
					Value:  20,
				},

				{
					Weight: 2,
					Value:  20,
				},

				{
					Weight: 2,
					Value:  20,
				},

				{
					Weight: 10,
					Value:  50,
				},
			},
		},
		expected: 80,
	},

	{
		description: "example knapsack",
		input: maximumValueCaseInput{
			MaximumWeight: 10,
			Items: []Item{
				{
					Weight: 5,
					Value:  10,
				},

				{
					Weight: 4,
					Value:  40,
				},

				{
					Weight: 6,
					Value:  30,
				},

				{
					Weight: 4,
					Value:  50,
				},
			},
		},
		expected: 90,
	},

	{
		description: "8 items",
		input: maximumValueCaseInput{
			MaximumWeight: 104,
			Items: []Item{
				{
					Weight: 25,
					Value:  350,
				},

				{
					Weight: 35,
					Value:  400,
				},

				{
					Weight: 45,
					Value:  450,
				},

				{
					Weight: 5,
					Value:  20,
				},

				{
					Weight: 25,
					Value:  70,
				},

				{
					Weight: 3,
					Value:  8,
				},

				{
					Weight: 2,
					Value:  5,
				},

				{
					Weight: 2,
					Value:  5,
				},
			},
		},
		expected: 900,
	},

	{
		description: "15 items",
		input: maximumValueCaseInput{
			MaximumWeight: 750,
			Items: []Item{
				{
					Weight: 70,
					Value:  135,
				},

				{
					Weight: 73,
					Value:  139,
				},

				{
					Weight: 77,
					Value:  149,
				},

				{
					Weight: 80,
					Value:  150,
				},

				{
					Weight: 82,
					Value:  156,
				},

				{
					Weight: 87,
					Value:  163,
				},

				{
					Weight: 90,
					Value:  173,
				},

				{
					Weight: 94,
					Value:  184,
				},

				{
					Weight: 98,
					Value:  192,
				},

				{
					Weight: 106,
					Value:  201,
				},

				{
					Weight: 110,
					Value:  210,
				},

				{
					Weight: 113,
					Value:  214,
				},

				{
					Weight: 115,
					Value:  221,
				},

				{
					Weight: 118,
					Value:  229,
				},

				{
					Weight: 120,
					Value:  240,
				},
			},
		},
		expected: 1458,
	},
}
