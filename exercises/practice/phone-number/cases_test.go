package phonenumber

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7a8722a Reorder keys (#1960)

type testCase struct {
	description       string
	input             string
	expectErr         bool
	expectedNumber    string
	expectedAreaCode  string
	expectedFormatted string
}

var testCases = []testCase{
	{
		description:       "cleans the number",
		input:             "(223) 456-7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "cleans numbers with dots",
		input:             "223.456.7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "cleans numbers with multiple spaces",
		input:             "223 456   7890   ",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "invalid when 9 digits",
		input:             "123456789",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid when 11 digits does not start with a 1",
		input:             "22234567890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "valid when 11 digits and starting with 1",
		input:             "12234567890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "valid when 11 digits and starting with 1 even with punctuation",
		input:             "+1 (223) 456-7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "invalid when more than 11 digits",
		input:             "321234567890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid with letters",
		input:             "523-abc-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid with punctuations",
		input:             "523-@:!-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 0",
		input:             "(023) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 1",
		input:             "(123) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 0",
		input:             "(223) 056-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 1",
		input:             "(223) 156-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 0 on valid 11-digit number",
		input:             "1 (023) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 1 on valid 11-digit number",
		input:             "1 (123) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 0 on valid 11-digit number",
		input:             "1 (223) 056-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 1 on valid 11-digit number",
		input:             "1 (223) 156-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
}
