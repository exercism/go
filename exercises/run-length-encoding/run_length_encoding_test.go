package encode

import "testing"

var encodeTests = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"XYZ", "XYZ"},
	{"AABBBCCCC", "2A3B4C"},
	{"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB", "12WB12W3B24WB"},
	{"  hsqq qww  ", "2 hs2q q2w2 "},
	{"aabbbcccc", "2a3b4c"},
}

var decodeTests = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"XYZ", "XYZ"},
	{"2A3B4C", "AABBBCCCC"},
	{"12WB12W3B24WB", "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"},
	{"2 hs2q q2w2 ", "  hsqq qww  "},
	{"2a3b4c", "aabbbcccc"},
}

var encodeDecodeTests = []struct {
	input    string
	expected string
}{
	{"zzz ZZ  zZ", "zzz ZZ  zZ"},
}

func TestRunLengthEncode(t *testing.T) {
	for _, test := range encodeTests {
		if actual := RunLengthEncode(test.input); actual != test.expected {
			t.Errorf("RunLengthEncode(%s) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
func TestRunLengthDecode(t *testing.T) {
	for _, test := range decodeTests {
		if actual := RunLengthDecode(test.input); actual != test.expected {
			t.Errorf("RunLengthDecode(%s) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
func TestRunLengthEncodeDecode(t *testing.T) {
	for _, test := range encodeDecodeTests {
		if actual := RunLengthDecode(RunLengthEncode(test.input)); actual != test.expected {
			t.Errorf("RunLengthDecode(RunLengthEncode(%s)) = %q, expected %q.",
				test.input, actual, test.expected)
		}
	}
}
