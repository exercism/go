package luhn

// Source: exercism/problem-specifications
// Commit: fc7ad52 luhn: test-case covering the usage of %5==0 (#2056)

var testCases = []struct {
	description string
	input       string
	ok          bool
}{
	{
		"single digit strings can not be valid",
		"1",
		false,
	},
	{
		"a single zero is invalid",
		"0",
		false,
	},
	{
		"a simple valid SIN that remains valid if reversed",
		"059",
		true,
	},
	{
		"a simple valid SIN that becomes invalid if reversed",
		"59",
		true,
	},
	{
		"a valid Canadian SIN",
		"055 444 285",
		true,
	},
	{
		"invalid Canadian SIN",
		"055 444 286",
		false,
	},
	{
		"invalid credit card",
		"8273 1232 7352 0569",
		false,
	},
	{
		"invalid long number with an even remainder",
		"1 2345 6789 1234 5678 9012",
		false,
	},
	{
		"invalid long number with a remainder divisible by 5",
		"1 2345 6789 1234 5678 9013",
		false,
	},
	{
		"valid number with an even number of digits",
		"095 245 88",
		true,
	},
	{
		"valid number with an odd number of spaces",
		"234 567 891 234",
		true,
	},
	{
		"valid strings with a non-digit added at the end become invalid",
		"059a",
		false,
	},
	{
		"valid strings with punctuation included become invalid",
		"055-444-285",
		false,
	},
	{
		"valid strings with symbols included become invalid",
		"055# 444$ 285",
		false,
	},
	{
		"single zero with space is invalid",
		" 0",
		false,
	},
	{
		"more than a single zero is valid",
		"0000 0",
		true,
	},
	{
		"input digit 9 is correctly converted to output digit 9",
		"091",
		true,
	},
	{
		"very long input is valid",
		"9999999999 9999999999 9999999999 9999999999",
		true,
	},
	{
		"valid luhn with an odd number of digits and non zero first digit",
		"109",
		true,
	},
	{
		"using ascii value for non-doubled non-digit isn't allowed",
		"055b 444 285",
		false,
	},
	{
		"using ascii value for doubled non-digit isn't allowed",
		":9",
		false,
	},
	{
		"non-numeric, non-space char in the middle with a sum that's divisible by 10 isn't allowed",
		"59%59",
		false,
	},
}
