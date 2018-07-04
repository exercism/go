package isbn

// Source: exercism/problem-specifications
// Commit: 3134243 isbn-verifier: Crafted input to catch more incorrect algorithms (#1255)
// Problem Specifications Version: 2.7.0

var testCases = []struct {
	isbn        string
	expected    bool
	description string
}{
	{"3-598-21508-8", true, "valid isbn number"},
	{"3-598-21508-9", false, "invalid isbn check digit"},
	{"3-598-21507-X", true, "valid isbn number with a check digit of 10"},
	{"3-598-21507-A", false, "check digit is a character other than X"},
	{"3-598-P1581-X", false, "invalid character in isbn"},
	{"3-598-2X507-9", false, "X is only valid as a check digit"},
	{"3598215088", true, "valid isbn without separating dashes"},
	{"359821507X", true, "isbn without separating dashes and X as check digit"},
	{"359821507", false, "isbn without check digit and dashes"},
	{"3598215078X", false, "too long isbn and no dashes"},
	{"00", false, "too short isbn"},
	{"3-598-21507", false, "isbn without check digit"},
	{"3-598-21515-X", false, "check digit of X should not be used for 0"},
	{"", false, "empty isbn"},
	{"134456729", false, "input is 9 characters"},
	{"3132P34035", false, "invalid characters are not ignored"},
	{"98245726788", false, "input is too long but contains a valid isbn"},
}
