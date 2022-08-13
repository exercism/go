package isbn

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: f830544 isbn-verifier - add test case (#1998)

var testCases = []struct {
	description string
	isbn        string
	expected    bool
}{
	{
		description: "valid isbn",
		isbn:        "3-598-21508-8",
		expected:    true,
	},
	{
		description: "invalid isbn check digit",
		isbn:        "3-598-21508-9",
		expected:    false,
	},
	{
		description: "valid isbn with a check digit of 10",
		isbn:        "3-598-21507-X",
		expected:    true,
	},
	{
		description: "check digit is a character other than X",
		isbn:        "3-598-21507-A",
		expected:    false,
	},
	{
		description: "invalid check digit in isbn is not treated as zero",
		isbn:        "4-598-21507-B",
		expected:    false,
	},
	{
		description: "invalid character in isbn is not treated as zero",
		isbn:        "3-598-P1581-X",
		expected:    false,
	},
	{
		description: "X is only valid as a check digit",
		isbn:        "3-598-2X507-9",
		expected:    false,
	},
	{
		description: "valid isbn without separating dashes",
		isbn:        "3598215088",
		expected:    true,
	},
	{
		description: "isbn without separating dashes and X as check digit",
		isbn:        "359821507X",
		expected:    true,
	},
	{
		description: "isbn without check digit and dashes",
		isbn:        "359821507",
		expected:    false,
	},
	{
		description: "too long isbn and no dashes",
		isbn:        "3598215078X",
		expected:    false,
	},
	{
		description: "too short isbn",
		isbn:        "00",
		expected:    false,
	},
	{
		description: "isbn without check digit",
		isbn:        "3-598-21507",
		expected:    false,
	},
	{
		description: "check digit of X should not be used for 0",
		isbn:        "3-598-21515-X",
		expected:    false,
	},
	{
		description: "empty isbn",
		isbn:        "",
		expected:    false,
	},
	{
		description: "input is 9 characters",
		isbn:        "134456729",
		expected:    false,
	},
	{
		description: "invalid characters are not ignored after checking length",
		isbn:        "3132P34035",
		expected:    false,
	},
	{
		description: "invalid characters are not ignored before checking length",
		isbn:        "3598P215088",
		expected:    false,
	},
	{
		description: "input is too long but contains a valid isbn",
		isbn:        "98245726788",
		expected:    false,
	},
}
