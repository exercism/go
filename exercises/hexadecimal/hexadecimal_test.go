// Your solution must include the following definitions:
//
// func ParseHex(string) (int64, error)
// func HandleErrors([]string) []string
//
// HandleErrors takes a list of inputs for ParseHex and returns a matching list
// of error cases.  It must call ParseHex on each input, handle the error result,
// and put one of three strings, "none", "syntax", or "range" in the result list
// according to the error.

package hexadecimal

import (
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

var testCases = []struct {
	in      string
	out     int64
	errCase string
}{
	{"1", 1, "none"},
	{"10", 0x10, "none"},
	{"2d", 0x2d, "none"},
	{"012", 0x12, "none"},
	{"cfcfcf", 0xcfcfcf, "none"},
	{"CFCFCF", 0xcfcfcf, "none"},
	{"", 0, "syntax"},
	{"peanut", 0, "syntax"},
	{"2cg134", 0, "syntax"},
	{"8000000000000000", 0, "range"},
	{"9223372036854775809", 0, "range"},
}

func TestParseHex(t *testing.T) {
	for _, test := range testCases {
		switch out, err := ParseHex(test.in); {
		case err != nil:
			if test.errCase == "none" {
				t.Errorf("ParseHex(%q) returned error %q.  Error not expected.",
					test.in, err)
			}
		case test.errCase != "none":
			t.Errorf("ParseHex(%q) = %d, %v. Expected error.",
				test.in, out, err)
		case out != test.out:
			t.Errorf("ParseHex(%q) = %d. Expected %d.",
				test.in, out, test.out)
		}
	}
}

func TestHandleErrors(t *testing.T) {
	tests := make([]string, len(testCases))
	for i, test := range testCases {
		tests[i] = test.in
	}
	er := HandleErrors(tests)
	if len(er) != len(tests) {
		t.Fatalf("For %d tests, HandleErrors returned %d results, want %d",
			len(tests), len(er), len(tests))
	}
	for i, e := range er {
		if e != testCases[i].errCase {
			t.Errorf("For ParseHex(%q), HandleErrors reports %q, want %q",
				tests[i], e, testCases[i].errCase)
		}
	}
}

func BenchmarkParseHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ParseHex(test.in)
		}
	}
}
