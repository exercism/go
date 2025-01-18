package wordy

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

type wordyTest struct {
	description string
	question    string
	expectError bool
	expected    int
}

var tests = []wordyTest{
	{
		description: "just a number",
		question:    "What is 5?",
		expectError: false,
		expected:    5,
	},
	{
		description: "addition",
		question:    "What is 1 plus 1?",
		expectError: false,
		expected:    2,
	},
	{
		description: "more addition",
		question:    "What is 53 plus 2?",
		expectError: false,
		expected:    55,
	},
	{
		description: "addition with negative numbers",
		question:    "What is -1 plus -10?",
		expectError: false,
		expected:    -11,
	},
	{
		description: "large addition",
		question:    "What is 123 plus 45678?",
		expectError: false,
		expected:    45801,
	},
	{
		description: "subtraction",
		question:    "What is 4 minus -12?",
		expectError: false,
		expected:    16,
	},
	{
		description: "multiplication",
		question:    "What is -3 multiplied by 25?",
		expectError: false,
		expected:    -75,
	},
	{
		description: "division",
		question:    "What is 33 divided by -3?",
		expectError: false,
		expected:    -11,
	},
	{
		description: "multiple additions",
		question:    "What is 1 plus 1 plus 1?",
		expectError: false,
		expected:    3,
	},
	{
		description: "addition and subtraction",
		question:    "What is 1 plus 5 minus -2?",
		expectError: false,
		expected:    8,
	},
	{
		description: "multiple subtraction",
		question:    "What is 20 minus 4 minus 13?",
		expectError: false,
		expected:    3,
	},
	{
		description: "subtraction then addition",
		question:    "What is 17 minus 6 plus 3?",
		expectError: false,
		expected:    14,
	},
	{
		description: "multiple multiplication",
		question:    "What is 2 multiplied by -2 multiplied by 3?",
		expectError: false,
		expected:    -12,
	},
	{
		description: "addition and multiplication",
		question:    "What is -3 plus 7 multiplied by -2?",
		expectError: false,
		expected:    -8,
	},
	{
		description: "multiple division",
		question:    "What is -12 divided by 2 divided by -3?",
		expectError: false,
		expected:    2,
	},
	{
		description: "unknown operation",
		question:    "What is 52 cubed?",
		expectError: true,
		expected:    0,
	},
	{
		description: "Non math question",
		question:    "Who is the President of the United States?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject problem missing an operand",
		question:    "What is 1 plus?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject problem with no operands or operators",
		question:    "What is?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject two operations in a row",
		question:    "What is 1 plus plus 2?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject two numbers in a row",
		question:    "What is 1 plus 2 1?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject postfix notation",
		question:    "What is 1 2 plus?",
		expectError: true,
		expected:    0,
	},
	{
		description: "reject prefix notation",
		question:    "What is plus 1 2?",
		expectError: true,
		expected:    0,
	},
}
