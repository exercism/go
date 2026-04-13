package ocrnumbers

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 2e820e1 Auto-format portions of some JSON files (#1967)

type testCase struct {
	description string
	input       string
	expected    []string
	errorMsg    string
}

var testCases = []testCase{
	{
		description: "Recognizes 0",
		input:       "\n _ \n| |\n|_|\n   ",
		expected:    []string{"0"},
	},
	{
		description: "Recognizes 1",
		input:       "\n   \n  |\n  |\n   ",
		expected:    []string{"1"},
	},
	{
		description: "Unreadable but correctly sized inputs return ?",
		input:       "\n   \n  _\n  |\n   ",
		expected:    []string{"?"},
	},
	{
		description: "Input with a number of lines that is not a multiple of four raises an error",
		input:       "\n _ \n| |\n   ",
		errorMsg:    "number of input lines is not a multiple of four",
	},
	{
		description: "Input with a number of columns that is not a multiple of three raises an error",
		input:       "\n    \n   |\n   |\n    ",
		errorMsg:    "number of input columns is not a multiple of three",
	},
	{
		description: "Recognizes 110101100",
		input:       "\n       _     _        _  _ \n  |  || |  || |  |  || || |\n  |  ||_|  ||_|  |  ||_||_|\n                           ",
		expected:    []string{"110101100"},
	},
	{
		description: "Garbled numbers in a string are replaced with ?",
		input:       "\n       _     _           _ \n  |  || |  || |     || || |\n  |  | _|  ||_|  |  ||_||_|\n                           ",
		expected:    []string{"11?10?1?0"},
	},
	{
		description: "Recognizes 2",
		input:       "\n _ \n _|\n|_ \n   ",
		expected:    []string{"2"},
	},
	{
		description: "Recognizes 3",
		input:       "\n _ \n _|\n _|\n   ",
		expected:    []string{"3"},
	},
	{
		description: "Recognizes 4",
		input:       "\n   \n|_|\n  |\n   ",
		expected:    []string{"4"},
	},
	{
		description: "Recognizes 5",
		input:       "\n _ \n|_ \n _|\n   ",
		expected:    []string{"5"},
	},
	{
		description: "Recognizes 6",
		input:       "\n _ \n|_ \n|_|\n   ",
		expected:    []string{"6"},
	},
	{
		description: "Recognizes 7",
		input:       "\n _ \n  |\n  |\n   ",
		expected:    []string{"7"},
	},
	{
		description: "Recognizes 8",
		input:       "\n _ \n|_|\n|_|\n   ",
		expected:    []string{"8"},
	},
	{
		description: "Recognizes 9",
		input:       "\n _ \n|_|\n _|\n   ",
		expected:    []string{"9"},
	},
	{
		description: "Recognizes string of decimal numbers",
		input:       "\n    _  _     _  _  _  _  _  _ \n  | _| _||_||_ |_   ||_||_|| |\n  ||_  _|  | _||_|  ||_| _||_|\n                              ",
		expected:    []string{"1234567890"},
	},
	{
		description: "Numbers separated by empty lines are recognized. Lines are joined by commas.",
		input:       "\n    _  _ \n  | _| _|\n  ||_  _|\n         \n    _  _ \n|_||_ |_ \n  | _||_|\n         \n _  _  _ \n  ||_||_|\n  ||_| _|\n         ",
		expected:    []string{"123", "456", "789"},
	},
}
