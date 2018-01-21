package wordy

// Source: exercism/problem-specifications
// Commit: df75482 wordy: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type wordyTest struct {
	description string
	question    string
	ok          bool
	answer      int
}

var tests = []wordyTest{
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
}
