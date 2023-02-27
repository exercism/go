package resistorcolor

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

type colorCodeTestCase struct {
	description string
	input       string
	expected    int
}

type colorsTestCase struct {
	description string
	input       string
	expected    []string
}

var colorCodeTestCases = []colorCodeTestCase{
	{
		description: "Black",
		input:       "black",
		expected:    0,
	},
	{
		description: "White",
		input:       "white",
		expected:    9,
	},
	{
		description: "Orange",
		input:       "orange",
		expected:    3,
	},
}

var colorsTestCases = []colorsTestCase{
	{
		description: "Colors",
		expected: []string{
			"black",
			"brown",
			"red",
			"orange",
			"yellow",
			"green",
			"blue",
			"violet",
			"grey",
			"white",
		},
	},
}
