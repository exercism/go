package isbn

import (
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestSteps(t *testing.T) {
	for _, test := range testCases {
		observed := IsVaildISBN(test.input)
		if observed != test.expected {
			t.Fatalf("IsValidISBN(%s) = %s, want %s (%s)",
				test.input, observed, test.expected, test.description)
		}
	}
}
