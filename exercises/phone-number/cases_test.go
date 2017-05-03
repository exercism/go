package phonenumber

// Source: exercism/x-common
// Commit: 51d2475 number: Rewrite description and add tests
// x-common version: 1.1.0

// Cleanup user-entered phone numbers
var numberTests = []struct {
	description string
	input       string
	expectErr   bool
	number      string
	areaCode    string
	formatted   string
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
		description: "invalid when 9 digits",
		input:       "123456789",
		expectErr:   true,
	},
	{
		description: "invalid when 11 digits does not start with a 1",
		input:       "22234567890",
		expectErr:   true,
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
		description: "invalid when more than 11 digits",
		input:       "321234567890",
		expectErr:   true,
	},
	{
		description: "invalid with letters",
		input:       "123-abc-7890",
		expectErr:   true,
	},
	{
		description: "invalid with punctuations",
		input:       "123-@:!-7890",
		expectErr:   true,
	},
	{
		description: "invalid with right number of digits but letters mixed in",
		input:       "1a2b3c4d5e6f7g8h9i0j",
		expectErr:   true,
	},
	{
		description: "invalid if area code does not start with 2-9",
		input:       "(123) 456-7890",
		expectErr:   true,
	},
	{
		description: "invalid if exchange code does not start with 2-9",
		input:       "(223) 056-7890",
		expectErr:   true,
	},
}
