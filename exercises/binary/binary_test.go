package binary

import (
	"testing"
)

// You must implement the function,
//
//    func ParseBinary(string) (int, error)
//
// It is standard for Go functions to return error values to report error conditions.
// The test cases below are all valid binary numbers however.  For this exercise you
// may simply return nil for the error value in all cases.
//
// For bonus points though, what errors might be possible when parsing a number?
// Can you add code to detect error conditions and return appropriate error values?

var testCases = []struct {
	binary   string
	expected int
}{
	{"1", 1},
	{"10", 2},
	{"11", 3},
	{"100", 4},
	{"1001", 9},
	{"11010", 26},
	{"10001101000", 1128},
	{"0", 0},
}

func TestParseBinary(t *testing.T) {
	for _, tt := range testCases {
		actual, err := ParseBinary(tt.binary)
		// We don't expect errors for any of the test cases.
		if err != nil {
			t.Fatalf("ParseBinary(%v) returned error %q.  Error not expected.",
				tt.binary, err)
		}
		// Well, we don't expect wrong answers either.
		if actual != tt.expected {
			t.Fatalf("ParseBinary(%v): actual %d, expected %v",
				tt.binary, actual, tt.expected)
		}
	}
}

// Benchmark combined time for all tests
func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			ParseBinary(tt.binary)
		}
	}
}
