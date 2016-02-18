package leap

import "testing"

// Define a function IsLeapYear(int) bool.
//
// Also define a testVersion with a value that matches
// the targetTestVersion here.

const targetTestVersion = 2

func TestLeapYears(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
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
