package flowerfield

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 60a5626 Add test for over-counting horizontal neighbors in Flower Field (#2603)

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{
	{
		description: "no rows",
		input:       []string{},
		expected:    []string{},
	},

	{
		description: "no columns",
		input: []string{
			"",
		},
		expected: []string{
			"",
		},
	},

	{
		description: "no flowers",
		input: []string{
			"   ",
			"   ",
			"   ",
		},
		expected: []string{
			"   ",
			"   ",
			"   ",
		},
	},

	{
		description: "garden full of flowers",
		input: []string{
			"***",
			"***",
			"***",
		},
		expected: []string{
			"***",
			"***",
			"***",
		},
	},

	{
		description: "flower surrounded by spaces",
		input: []string{
			"   ",
			" * ",
			"   ",
		},
		expected: []string{
			"111",
			"1*1",
			"111",
		},
	},

	{
		description: "space surrounded by flowers",
		input: []string{
			"***",
			"* *",
			"***",
		},
		expected: []string{
			"***",
			"*8*",
			"***",
		},
	},

	{
		description: "horizontal line",
		input: []string{
			" * * ",
		},
		expected: []string{
			"1*2*1",
		},
	},

	{
		description: "horizontal line, flowers at edges",
		input: []string{
			"*   *",
		},
		expected: []string{
			"*1 1*",
		},
	},

	{
		description: "vertical line",
		input: []string{
			" ",
			"*",
			" ",
			"*",
			" ",
		},
		expected: []string{
			"1",
			"*",
			"2",
			"*",
			"1",
		},
	},

	{
		description: "vertical line, flowers at edges",
		input: []string{
			"*",
			" ",
			" ",
			" ",
			"*",
		},
		expected: []string{
			"*",
			"1",
			" ",
			"1",
			"*",
		},
	},

	{
		description: "cross",
		input: []string{
			"  *  ",
			"  *  ",
			"*****",
			"  *  ",
			"  *  ",
		},
		expected: []string{
			" 2*2 ",
			"25*52",
			"*****",
			"25*52",
			" 2*2 ",
		},
	},

	{
		description: "large garden",
		input: []string{
			" *  * ",
			"  *   ",
			"    * ",
			"   * *",
			" *  * ",
			"      ",
		},
		expected: []string{
			"1*22*1",
			"12*322",
			" 123*2",
			"112*4*",
			"1*22*2",
			"111111",
		},
	},

	{
		description: "multiple adjacent flowers",
		input: []string{
			" ** ",
		},
		expected: []string{
			"1**1",
		},
	},
}
