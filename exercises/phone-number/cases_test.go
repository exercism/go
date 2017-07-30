package phonenumber

// Source: exercism/x-common
// Commit: 39cba0d phone-number: Remove test using malformed input. (#772)
// x-common version: 1.2.0

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
