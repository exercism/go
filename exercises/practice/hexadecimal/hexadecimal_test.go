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
	"strings"
	"testing"
)

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
		out, err := ParseHex(test.in)
		if test.errCase != "none" {
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Errorf("ParseHex(%q): expected an error, but error is nil",
					test.in)
				continue
			}

			if !strings.Contains(strings.ToLower(err.Error()), test.errCase) {
				t.Errorf(
					"ParseHex(%q) returned error %q. Expected error containing '%s'.",
					test.in, err, test.errCase)
			}
		} else {
			if out != test.out {
				t.Errorf("ParseHex(%q) = %d. Expected %d.",
					test.in, out, test.out)
			}

			// we dont expect error
			if err != nil {
				t.Errorf("ParseHex(%q) returned error %q.  Error not expected.",
					test.in, err)
			}
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ParseHex(test.in)
		}
	}
}
