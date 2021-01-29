package wordy

// Source: exercism/problem-specifications
// Commit: 51c9b0b wordy: Test question `What is?` because it's too short
// Problem Specifications Version: 1.5.0

type wordyTest struct {
	description string
	question    string
	ok          bool
	answer      int
}

var tests = []wordyTest{
	{
		"just a number",
		"What is 5?",
		true,
		5,
	},
	{
		"addition",
		"What is 1 plus 1?",
		true,
		2,
	},
	{
		"more addition",
		"What is 53 plus 2?",
		true,
		55,
	},
	{
		"addition with negative numbers",
		"What is -1 plus -10?",
		true,
		-11,
	},
	{
		"large addition",
		"What is 123 plus 45678?",
		true,
		45801,
	},
	{
		"subtraction",
		"What is 4 minus -12?",
		true,
		16,
	},
	{
		"multiplication",
		"What is -3 multiplied by 25?",
		true,
		-75,
	},
	{
		"division",
		"What is 33 divided by -3?",
		true,
		-11,
	},
	{
		"multiple additions",
		"What is 1 plus 1 plus 1?",
		true,
		3,
	},
	{
		"addition and subtraction",
		"What is 1 plus 5 minus -2?",
		true,
		8,
	},
	{
		"multiple subtraction",
		"What is 20 minus 4 minus 13?",
		true,
		3,
	},
	{
		"subtraction then addition",
		"What is 17 minus 6 plus 3?",
		true,
		14,
	},
	{
		"multiple multiplication",
		"What is 2 multiplied by -2 multiplied by 3?",
		true,
		-12,
	},
	{
		"addition and multiplication",
		"What is -3 plus 7 multiplied by -2?",
		true,
		-8,
	},
	{
		"multiple division",
		"What is -12 divided by 2 divided by -3?",
		true,
		2,
	},
	{
		"unknown operation",
		"What is 52 cubed?",
		false,
		0,
	},
	{
		"Non math question",
		"Who is the President of the United States?",
		false,
		0,
	},
	{
		"reject problem missing an operand",
		"What is 1 plus?",
		false,
		0,
	},
	{
		"reject problem with no operands or operators",
		"What is?",
		false,
		0,
	},
	{
		"reject two operations in a row",
		"What is 1 plus plus 2?",
		false,
		0,
	},
	{
		"reject two numbers in a row",
		"What is 1 plus 2 1?",
		false,
		0,
	},
	{
		"reject postfix notation",
		"What is 1 2 plus?",
		false,
		0,
	},
	{
		"reject prefix notation",
		"What is plus 1 2?",
		false,
		0,
	},
}
