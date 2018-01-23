package rectangles

// Source: exercism/problem-specifications
// Commit: 61bccfa rectangles: apply "input" policy
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	input       []string
	expected    int
}{
	{
		description: "no rows",
		input:       []string{},
		expected:    0,
	},
	{
		description: "no columns",
		input: []string{
			"",
		},
		expected: 0,
	},
	{
		description: "no rectangles",
		input: []string{
			" ",
		},
		expected: 0,
	},
	{
		description: "one rectangle",
		input: []string{
			"+-+",
			"| |",
			"+-+",
		},
		expected: 1,
	},
	{
		description: "two rectangles without shared parts",
		input: []string{
			"  +-+",
			"  | |",
			"+-+-+",
			"| |  ",
			"+-+  ",
		},
		expected: 2,
	},
	{
		description: "five rectangles with shared parts",
		input: []string{
			"  +-+",
			"  | |",
			"+-+-+",
			"| | |",
			"+-+-+",
		},
		expected: 5,
	},
	{
		description: "rectangle of height 1 is counted",
		input: []string{
			"+--+",
			"+--+",
		},
		expected: 1,
	},
	{
		description: "rectangle of width 1 is counted",
		input: []string{
			"++",
			"||",
			"++",
		},
		expected: 1,
	},
	{
		description: "1x1 square is counted",
		input: []string{
			"++",
			"++",
		},
		expected: 1,
	},
	{
		description: "only complete rectangles are counted",
		input: []string{
			"  +-+",
			"    |",
			"+-+-+",
			"| | -",
			"+-+-+",
		},
		expected: 1,
	},
	{
		description: "rectangles can be of different sizes",
		input: []string{
			"+------+----+",
			"|      |    |",
			"+---+--+    |",
			"|   |       |",
			"+---+-------+",
		},
		expected: 3,
	},
	{
		description: "corner is required for a rectangle to be complete",
		input: []string{
			"+------+----+",
			"|      |    |",
			"+------+    |",
			"|   |       |",
			"+---+-------+",
		},
		expected: 2,
	},
	{
		description: "large input with many rectangles",
		input: []string{
			"+---+--+----+",
			"|   +--+----+",
			"+---+--+    |",
			"|   +--+----+",
			"+---+--+--+-+",
			"+---+--+--+-+",
			"+------+  | |",
			"          +-+",
		},
		expected: 60,
	},
}
