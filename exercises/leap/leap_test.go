package leap

import "testing"

// targetTestVersion is used to ensure that a solution is only evaluated
// against the correct version of the test suite.
const targetTestVersion = 3

// TestTestVersion compares the targetTestVersion defined above with
// a testVersion value.
// This is a common convention throughout the Go Exercism track.
// To make a test like this test pass, define 'testVersion' in your solution.
// We've done this for you in the ./leap.go file provided.
func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestLeapYears(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("IsLeapYear(%d) = %t, want %t (%s)",
				test.year, observed, test.expected, test.description)
		}
	}
}

// Benchmark 400 year interval to get fair weighting of different years.
func Benchmark400(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := 1600; y < 2000; y++ {
			IsLeapYear(y)
		}
	}
}
