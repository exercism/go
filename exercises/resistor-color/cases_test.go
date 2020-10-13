package resistorcolor

// Source: exercism/problem-specifications
// Commit: cacf1f1 Acronym: add underscore test case (#1436)
// Problem Specifications Version: 1.7.0

type resistorColorTest struct {
	input    string
	expected string
}

var stringTestCases = []resistorColorTest{
	{
		input:    "black",
		expected: 0,
	},
	{
		input:    "brown",
		expected: 1,
	},
	{
		input:    "red",
		expected: 2,
	},
	{
		input:    "orange",
		expected: 3,
	},
	{
		input:    "yellow",
		expected: 4,
	},
	{
		input:    "green",
		expected: 5,
	},
	{
		input:    "blue",
		expected: 6,
	},
	{
		input:    "violet",
		expected: 7,
	},
	{
		input:    "grey",
		expected: 8,
	},
	{
		input:    "white",
		expected: 9,
	},
}

var colorsExpected = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
