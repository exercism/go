package resistorcolorduo

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 25c2ae5 build(deps): bump fast-uri from 3.1.0 to 3.1.2 (#2653)

type valueTestCase struct {
	description string
	input       []string
	expected    int
}

var valueTestCases = []valueTestCase{
	{
		description: "Brown and black",
		input:       []string{"brown", "black"},
		expected:    10,
	},
	{
		description: "Blue and grey",
		input:       []string{"blue", "grey"},
		expected:    68,
	},
	{
		description: "Yellow and violet",
		input:       []string{"yellow", "violet"},
		expected:    47,
	},
	{
		description: "White and red",
		input:       []string{"white", "red"},
		expected:    92,
	},
	{
		description: "Orange and orange",
		input:       []string{"orange", "orange"},
		expected:    33,
	},
	{
		description: "Ignore additional colors",
		input:       []string{"green", "brown", "orange"},
		expected:    51,
	},
	{
		description: "Black and brown, one-digit",
		input:       []string{"black", "brown"},
		expected:    1,
	},
}
