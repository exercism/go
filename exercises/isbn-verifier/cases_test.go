package isbn

// Source: exercism/problem-specifications
// Commit: adb99f7 isbn-verifier: Add test case with only 9 digits (#1221)
// Problem Specifications Version: 2.4.0

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
	{"3-598-21507", false, "isbn without check digit"},
	{"3-598-21507-XX", false, "too long isbn"},
	{"3-598-21515-X", false, "check digit of X should not be used for 0"},
	{"", false, "empty isbn"},
	{"134456729", false, "input is 9 characters"},
}
