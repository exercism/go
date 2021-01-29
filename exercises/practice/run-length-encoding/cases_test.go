package encode

// Source: exercism/problem-specifications
// Commit: 1b7900e run-length-encoding: apply "input" policy
// Problem Specifications Version: 1.1.0

// run-length encode a string
var encodeTests = []struct {
	input       string
	expected    string
	description string
}{
	{"", "", "empty string"},
	{"XYZ", "XYZ", "single characters only are encoded without count"},
	{"AABBBCCCC", "2A3B4C", "string with no single characters"},
	{"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "12WB12W3B24WB", "single characters mixed with repeated characters"},
	{"  hsqq qww  ", "2 hs2q q2w2 ", "multiple whitespace mixed in string"},
	{"aabbbcccc", "2a3b4c", "lowercase characters"},
}

// run-length decode a string
var decodeTests = []struct {
	input       string
	expected    string
	description string
}{
	{"", "", "empty string"},
	{"XYZ", "XYZ", "single characters only"},
	{"2A3B4C", "AABBBCCCC", "string with no single characters"},
	{"12WB12W3B24WB", "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "single characters with repeated characters"},
	{"2 hs2q q2w2 ", "  hsqq qww  ", "multiple whitespace mixed in string"},
	{"2a3b4c", "aabbbcccc", "lower case string"},
}

// encode and then decode
var encodeDecodeTests = []struct {
	input       string
	expected    string
	description string
}{
	{"zzz ZZ  zZ", "zzz ZZ  zZ", "encode followed by decode gives original string"},
}
