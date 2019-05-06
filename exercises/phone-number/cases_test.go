package phonenumber

// Source: exercism/problem-specifications
// Commit: 435b472 phone-number: improve error message for letters
// Problem Specifications Version: 1.7.0

// Cleanup user-entered phone numbers
var numberTests = []struct {
	description      string
	input            string
	expectErr        bool
	errorDescription string
	number           string
	areaCode         string
	formatted        string
}{
	{
		description: "cleans the number",
		input:       "(223) 456-7890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
	},
	{
		description: "cleans numbers with dots",
		input:       "223.456.7890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
	},
	{
		description: "cleans numbers with multiple spaces",
		input:       "223 456   7890   ",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
	},
	{
		description:      "invalid when 9 digits",
		input:            "123456789",
		expectErr:        true,
		errorDescription: "incorrect number of digits",
	},
	{
		description:      "invalid when 11 digits does not start with a 1",
		input:            "22234567890",
		expectErr:        true,
		errorDescription: "11 digits must start with 1",
	},
	{
		description: "valid when 11 digits and starting with 1",
		input:       "12234567890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
	},
	{
		description: "valid when 11 digits and starting with 1 even with punctuation",
		input:       "+1 (223) 456-7890",
		number:      "2234567890",
		areaCode:    "223",
		formatted:   "(223) 456-7890",
	},
	{
		description:      "invalid when more than 11 digits",
		input:            "321234567890",
		expectErr:        true,
		errorDescription: "more than 11 digits",
	},
	{
		description:      "invalid with letters",
		input:            "123-abc-7890",
		expectErr:        true,
		errorDescription: "letters not permitted",
	},
	{
		description:      "invalid with punctuations",
		input:            "123-@:!-7890",
		expectErr:        true,
		errorDescription: "punctuations not permitted",
	},
	{
		description:      "invalid if area code starts with 0",
		input:            "(023) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with zero",
	},
	{
		description:      "invalid if area code starts with 1",
		input:            "(123) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with one",
	},
	{
		description:      "invalid if exchange code starts with 0",
		input:            "(223) 056-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with zero",
	},
	{
		description:      "invalid if exchange code starts with 1",
		input:            "(223) 156-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with one",
	},
	{
		description:      "invalid if area code starts with 0 on valid 11-digit number",
		input:            "1 (023) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with zero",
	},
	{
		description:      "invalid if area code starts with 1 on valid 11-digit number",
		input:            "1 (123) 456-7890",
		expectErr:        true,
		errorDescription: "area code cannot start with one",
	},
	{
		description:      "invalid if exchange code starts with 0 on valid 11-digit number",
		input:            "1 (223) 056-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with zero",
	},
	{
		description:      "invalid if exchange code starts with 1 on valid 11-digit number",
		input:            "1 (223) 156-7890",
		expectErr:        true,
		errorDescription: "exchange code cannot start with one",
	},
}
