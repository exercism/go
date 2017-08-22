package hamming

import "testing"

// targetTestVersion is used to ensure that a solution is only evaluated
// against the correct version of the test suite.
const targetTestVersion = 6

// TestTestVersion compares the targetTestVersion defined above with
// a testVersion value.
// This is a common convention throughout the Go Exercism track.  To make a
// test like this test pass, define the 'testVersion' const in your solution,
// and give it the same value as `targetTestVersion`.  We've done this for you
// in the ./hamming.go file provided.
func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.want < 0 {
			// check if err is of error type
			var _ error = err

			// we expect an error
			if err == nil {
				t.Fatalf("Distance(%q, %q). error is nil.",
					tc.s1, tc.s2)
			}
		} else {
			if got != tc.want {
				t.Fatalf("Distance(%q, %q) = %d, want %d.",
					tc.s1, tc.s2, got, tc.want)
			}

			// we do not expect an error
			if err != nil {
				t.Fatalf("Distance(%q, %q) returned error: %v when expecting none.",
					tc.s1, tc.s2, err)
			}
		}
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench the combined time it takes
	// to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Distance(tc.s1, tc.s2)
		}
	}
}
