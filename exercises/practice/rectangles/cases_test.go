package rectangles

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 55b098b rectangles: Add test for missing sides (#2050)

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
	{
		description: "rectangles must have four sides",
		input: []string{
			"+-+ +-+",
			"| | | |",
			"+-+-+-+",
			"  | |  ",
			"+-+-+-+",
			"| | | |",
			"+-+ +-+",
		},
		expected: 5,
	},
}
