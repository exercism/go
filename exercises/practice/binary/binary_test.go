package binary

import (
	"testing"
)

// You must implement the function,
//
//    func ParseBinary(string) (int, error)
//
// It is standard for Go functions to return error values to report error conditions.
// The test cases have some inputs that are invalid.
// For invalid inputs, return an error that signals to the user why the error happened.
// The test cases can only check that you return *some* error,
// but it's still good practice to return useful errors.

var testCases = []struct {
	binary   string
	expected int
	ok       bool
}{
	{"1", 1, true},
	{"10", 2, true},
	{"11", 3, true},
	{"100", 4, true},
	{"1001", 9, true},
	{"11010", 26, true},
	{"10001101000", 1128, true},
	{"0", 0, true},
	{"foo101", 0, false},
	{"101bar", 0, false},
	{"101baz010", 0, false},
	{"22", 0, false},
}

func TestParseBinary(t *testing.T) {
	for _, tc := range testCases {
		actual, err := ParseBinary(tc.binary)
		if tc.ok {
			if err != nil {
				var _ error = err
				t.Fatalf("ParseBinary(%q) returned error %q.  Error not expected.",
					tc.binary, err)
			}
			if actual != tc.expected {
				t.Fatalf("ParseBinary(%q): got: %d, want: %v",
					tc.binary, actual, tc.expected)
			}
		} else if err == nil {
			t.Fatalf("ParseBinary(%q) returned %d and no error.  Expected an error.",
				tc.binary, actual)
		}
	}
}

// Benchmark combined time for all tests
func BenchmarkBinary(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			ParseBinary(tc.binary)
		}
	}
}
