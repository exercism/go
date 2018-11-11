package brackets

// Source: exercism/problem-specifications
// Commit: 602c610 bracket-push: Add test to prevent early return true (#1300)
// Problem Specifications Version: 1.4.0

type bracketTest struct {
	input    string
	expected bool
}

var testCases = []bracketTest{
	{
		// paired square brackets
		"[]",
		true,
	},
	{
		// empty string
		"",
		true,
	},
	{
		// unpaired brackets
		"[[",
		false,
	},
	{
		// wrong ordered brackets
		"}{",
		false,
	},
	{
		// wrong closing bracket
		"{]",
		false,
	},
	{
		// paired with whitespace
		"{ }",
		true,
	},
	{
		// partially paired brackets
		"{[])",
		false,
	},
	{
		// simple nested brackets
		"{[]}",
		true,
	},
	{
		// several paired brackets
		"{}[]",
		true,
	},
	{
		// paired and nested brackets
		"([{}({}[])])",
		true,
	},
	{
		// unopened closing brackets
		"{[)][]}",
		false,
	},
	{
		// unpaired and nested brackets
		"([{])",
		false,
	},
	{
		// paired and wrong nested brackets
		"[({]})",
		false,
	},
	{
		// paired and incomplete brackets
		"{}[",
		false,
	},
	{
		// math expression
		"(((185 + 223.85) * 15) - 543)/2",
		true,
	},
	{
		// complex latex expression
		"\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)",
		true,
	},
}
