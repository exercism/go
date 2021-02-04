package complexnumbers

// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type testGroup struct {
	group string
	tests []testCase

	input    inputType
	expected interface{}
}

type testCase struct {
	description string
	input       inputType
	expected    interface{}

	group string
	tests []testCase
}

type inputType struct {
	z  []interface{}
	z1 []interface{}
	z2 []interface{}
}

var testGroups = []testGroup{
	{
		group: "Real part",
		tests: []testCase{
			{
				description: "Real part of a purely real number",
				input: inputType{
					z: []interface{}{1, 0},
				},
				expected: 1,
			},
			{
				description: "Real part of a purely imaginary number",
				input: inputType{
					z: []interface{}{0, 1},
				},
				expected: 0,
			},
			{
				description: "Real part of a number with real and imaginary part",
				input: inputType{
					z: []interface{}{1, 2},
				},
				expected: 1,
			},
		},
	},
	{
		group: "Imaginary part",
		tests: []testCase{
			{
				description: "Imaginary part of a purely real number",
				input: inputType{
					z: []interface{}{1, 0},
				},
				expected: 0,
			},
			{
				description: "Imaginary part of a purely imaginary number",
				input: inputType{
					z: []interface{}{0, 1},
				},
				expected: 1,
			},
			{
				description: "Imaginary part of a number with real and imaginary part",
				input: inputType{
					z: []interface{}{1, 2},
				},
				expected: 2,
			},
		},
	},
	{
		group: "Imaginary unit",
		input: inputType{
			z1: []interface{}{0, 1},
			z2: []interface{}{0, 1},
		},
		expected: []interface{}{-1, 0},
	},
	{
		group: "Arithmetic",
		tests: []testCase{
			{
				group: "Addition",
				tests: []testCase{
					{
						description: "Add purely real numbers",
						input: inputType{
							z1: []interface{}{1, 0},
							z2: []interface{}{2, 0},
						},
						expected: []interface{}{3, 0},
					},
					{
						description: "Add purely imaginary numbers",
						input: inputType{
							z1: []interface{}{0, 1},
							z2: []interface{}{0, 2},
						},
						expected: []interface{}{0, 3},
					},
					{
						description: "Add numbers with real and imaginary part",
						input: inputType{
							z1: []interface{}{1, 2},
							z2: []interface{}{3, 4},
						},
						expected: []interface{}{4, 6},
					},
				},
			},
			{
				group: "Subtraction",
				tests: []testCase{
					{
						description: "Subtract purely real numbers",
						input: inputType{
							z1: []interface{}{1, 0},
							z2: []interface{}{2, 0},
						},
						expected: []interface{}{-1, 0},
					},
					{
						description: "Subtract purely imaginary numbers",
						input: inputType{
							z1: []interface{}{0, 1},
							z2: []interface{}{0, 2},
						},
						expected: []interface{}{0, -1},
					},
					{
						description: "Subtract numbers with real and imaginary part",
						input: inputType{
							z1: []interface{}{1, 2},
							z2: []interface{}{3, 4},
						},
						expected: []interface{}{-2, -2},
					},
				},
			},
			{
				group: "Multiplication",
				tests: []testCase{
					{
						description: "Multiply purely real numbers",
						input: inputType{
							z1: []interface{}{1, 0},
							z2: []interface{}{2, 0},
						},
						expected: []interface{}{2, 0},
					},
					{
						description: "Multiply purely imaginary numbers",
						input: inputType{
							z1: []interface{}{0, 1},
							z2: []interface{}{0, 2},
						},
						expected: []interface{}{-2, 0},
					},
					{
						description: "Multiply numbers with real and imaginary part",
						input: inputType{
							z1: []interface{}{1, 2},
							z2: []interface{}{3, 4},
						},
						expected: []interface{}{-5, 10},
					},
				},
			},
			{
				group: "Division",
				tests: []testCase{
					{
						description: "Divide purely real numbers",
						input: inputType{
							z1: []interface{}{1, 0},
							z2: []interface{}{2, 0},
						},
						expected: []interface{}{0.5, 0},
					},
					{
						description: "Divide purely imaginary numbers",
						input: inputType{
							z1: []interface{}{0, 1},
							z2: []interface{}{0, 2},
						},
						expected: []interface{}{0.5, 0},
					},
					{
						description: "Divide numbers with real and imaginary part",
						input: inputType{
							z1: []interface{}{1, 2},
							z2: []interface{}{3, 4},
						},
						expected: []interface{}{0.44, 0.08},
					},
				},
			},
		},
	},
	{
		group: "Absolute value",
		tests: []testCase{
			{
				description: "Absolute value of a positive purely real number",
				input: inputType{
					z: []interface{}{5, 0},
				},
				expected: 5,
			},
			{
				description: "Absolute value of a negative purely real number",
				input: inputType{
					z: []interface{}{-5, 0},
				},
				expected: 5,
			},
			{
				description: "Absolute value of a purely imaginary number with positive imaginary part",
				input: inputType{
					z: []interface{}{0, 5},
				},
				expected: 5,
			},
			{
				description: "Absolute value of a purely imaginary number with negative imaginary part",
				input: inputType{
					z: []interface{}{0, -5},
				},
				expected: 5,
			},
			{
				description: "Absolute value of a number with real and imaginary part",
				input: inputType{
					z: []interface{}{3, 4},
				},
				expected: 5,
			},
		},
	},
	{
		group: "Complex conjugate",
		tests: []testCase{
			{
				description: "Conjugate a purely real number",
				input: inputType{
					z: []interface{}{5, 0},
				},
				expected: []interface{}{5, 0},
			},
			{
				description: "Conjugate a purely imaginary number",
				input: inputType{
					z: []interface{}{0, 5},
				},
				expected: []interface{}{0, -5},
			},
			{
				description: "Conjugate a number with real and imaginary part",
				input: inputType{
					z: []interface{}{1, 1},
				},
				expected: []interface{}{1, -1},
			},
		},
	},
	{
		group: "Complex exponential function",
		tests: []testCase{
			{
				description: "Euler's identity/formula",
				input: inputType{
					z: []interface{}{0, "pi"},
				},
				expected: []interface{}{-1, 0},
			},
			{
				description: "Exponential of 0",
				input: inputType{
					z: []interface{}{0, 0},
				},
				expected: []interface{}{1, 0},
			},
			{
				description: "Exponential of a purely real number",
				input: inputType{
					z: []interface{}{1, 0},
				},
				expected: []interface{}{"e", 0},
			},
			{
				description: "Exponential of a number with real and imaginary part",
				input: inputType{
					z: []interface{}{"ln(2)", "pi"},
				},
				expected: []interface{}{-2, 0},
			},
		},
	},
}

// !NOTE: Error during source formatting: Line:Column 134:2: missing ',' before newline in composite literal (and 3 more errors)
