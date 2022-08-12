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

type testCase struct {
	binary   string
	expected int
}

var validTestCases = []testCase{
	{binary: "1", expected: 1},
	{binary: "10", expected: 2},
	{binary: "11", expected: 3},
	{binary: "100", expected: 4},
	{binary: "1001", expected: 9},
	{binary: "11010", expected: 26},
	{binary: "10001101000", expected: 1128},
	{binary: "0", expected: 0},
}

func TestParseBinary(t *testing.T) {
	for _, tt := range validTestCases {
		t.Run(tt.binary, func(t *testing.T) {
			actual, err := ParseBinary(tt.binary)
			if err != nil {
				t.Fatalf("ParseBinary(%q) returned error: %v, want: %d", tt.binary, err, tt.expected)
			}
			if actual != tt.expected {
				t.Fatalf("ParseBinary(%q) = %d, want: %d", tt.binary, actual, tt.expected)
			}
		})
	}
}

var invalidTestCases = []testCase{
	{binary: "foo101", expected: 0},
	{binary: "101bar", expected: 0},
	{binary: "101baz010", expected: 0},
	{binary: "22", expected: 0},
}

func TestParseBinaryInvalid(t *testing.T) {
	for _, tt := range invalidTestCases {
		t.Run(tt.binary, func(t *testing.T) {
			actual, err := ParseBinary(tt.binary)
			if err == nil {
				t.Fatalf("ParseBinary(%q) expected error, got: %d", tt.binary, actual)
			}
		})
	}
}

// Benchmark combined time for all tests
func BenchmarkBinary(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range validTestCases {
			ParseBinary(tt.binary)
		}
	}
}
