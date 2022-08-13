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
	in       string
	expected int64
	errCase  string
}{
	{in: "1", expected: 1, errCase: "none"},
	{in: "10", expected: 0x10, errCase: "none"},
	{in: "2d", expected: 0x2d, errCase: "none"},
	{in: "012", expected: 0x12, errCase: "none"},
	{in: "cfcfcf", expected: 0xcfcfcf, errCase: "none"},
	{in: "CFCFCF", expected: 0xcfcfcf, errCase: "none"},
	{in: "", expected: 0, errCase: "syntax"},
	{in: "peanut", expected: 0, errCase: "syntax"},
	{in: "2cg134", expected: 0, errCase: "syntax"},
	{in: "8000000000000000", expected: 0, errCase: "range"},
	{in: "9223372036854775809", expected: 0, errCase: "range"},
}

func TestParseHex(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.in, func(t *testing.T) {
			actual, err := ParseHex(test.in)
			switch {
			case test.errCase != "none":
				if err == nil {
					t.Fatalf("ParseHex(%q) expected an error, got: %d", test.in, actual)
				}

				if !strings.Contains(strings.ToLower(err.Error()), test.errCase) {
					t.Fatalf("ParseHex(%q) returned error %q but expected error containing %q", test.in, err, test.errCase)
				}
			case err != nil:
				t.Fatalf("ParseHex(%q) returned error %q, want: %d", test.in, err, test.expected)
			case actual != test.expected:
				t.Errorf("ParseHex(%q) = %d, want: %d", test.in, actual, test.expected)
			}
		})
	}
}

func TestHandleErrors(t *testing.T) {
	tests := make([]string, len(testCases))
	for i, test := range testCases {
		tests[i] = test.in
	}
	got := HandleErrors(tests)
	if len(got) != len(tests) {
		t.Fatalf("HandleErrors() returned %d results for %d tests, want: %d", len(got), len(tests), len(tests))
	}
	for i, err := range got {
		if err != testCases[i].errCase {
			t.Errorf("For ParseHex(%q), HandleErrors reports %q, want %q", tests[i], err, testCases[i].errCase)
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
