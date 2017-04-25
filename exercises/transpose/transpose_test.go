package transpose

import (
	"reflect"
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestTranspose(t *testing.T) {
	for _, test := range testCases {
		actual := Transpose(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			// check for zero length slices
			if len(actual) == 0 || len(test.expected) == 0 {
				t.Fatalf("\n\tTranspose(%q): %s\n\n\tExpected: %q\n\tGot: %q",
					test.input, test.description, test.expected, actual)
			}
			// let's make the error more specific and find the row it's on
			min := min(len(test.expected), len(actual))
			for i := 0; i < min; i++ {
				if test.expected[i] != actual[i] {
					t.Fatalf("\n\tTranspose(%q): %s\n\n\tExpected: %q\n\tGot: %q\n\n\tRow %d Expected: %q Got: %q",
						test.input, test.description, test.expected, actual, i, test.expected[i], actual[i])
				}
			}
		}
	}
}

// helper function
// https://stackoverflow.com/questions/27516387/what-is-the-correct-way-to-find-the-min-between-two-integers-in-go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BenchmarkTranspose(b *testing.B) {
	for _, test := range testCases {
		for i := 0; i < b.N; i++ {
			Transpose(test.input)
		}
	}
}
