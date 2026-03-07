package lineup

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 4ed5a3c Add new exercise: line-up (#2595)

var testCases = []struct {
	description string
	name        string
	number      int
	expected    string
}{
	{
		description: "format smallest non-exceptional ordinal numeral 4",
		name:        "Gianna",
		number:      4,
		expected:    "Gianna, you are the 4th customer we serve today. Thank you!",
	},
	{
		description: "format greatest single digit non-exceptional ordinal numeral 9",
		name:        "Maarten",
		number:      9,
		expected:    "Maarten, you are the 9th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 5",
		name:        "Petronila",
		number:      5,
		expected:    "Petronila, you are the 5th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 6",
		name:        "Attakullakulla",
		number:      6,
		expected:    "Attakullakulla, you are the 6th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 7",
		name:        "Kate",
		number:      7,
		expected:    "Kate, you are the 7th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 8",
		name:        "Maximiliano",
		number:      8,
		expected:    "Maximiliano, you are the 8th customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 1",
		name:        "Mary",
		number:      1,
		expected:    "Mary, you are the 1st customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 2",
		name:        "Haruto",
		number:      2,
		expected:    "Haruto, you are the 2nd customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 3",
		name:        "Henriette",
		number:      3,
		expected:    "Henriette, you are the 3rd customer we serve today. Thank you!",
	},
	{
		description: "format smallest two digit non-exceptional ordinal numeral 10",
		name:        "Alvarez",
		number:      10,
		expected:    "Alvarez, you are the 10th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 11",
		name:        "Jacqueline",
		number:      11,
		expected:    "Jacqueline, you are the 11th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 12",
		name:        "Juan",
		number:      12,
		expected:    "Juan, you are the 12th customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 13",
		name:        "Patricia",
		number:      13,
		expected:    "Patricia, you are the 13th customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 21",
		name:        "Washi",
		number:      21,
		expected:    "Washi, you are the 21st customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 62",
		name:        "Nayra",
		number:      62,
		expected:    "Nayra, you are the 62nd customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 100",
		name:        "John",
		number:      100,
		expected:    "John, you are the 100th customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 101",
		name:        "Zeinab",
		number:      101,
		expected:    "Zeinab, you are the 101st customer we serve today. Thank you!",
	},
	{
		description: "format non-exceptional ordinal numeral 112",
		name:        "Knud",
		number:      112,
		expected:    "Knud, you are the 112th customer we serve today. Thank you!",
	},
	{
		description: "format exceptional ordinal numeral 123",
		name:        "Yma",
		number:      123,
		expected:    "Yma, you are the 123rd customer we serve today. Thank you!",
	},
}
