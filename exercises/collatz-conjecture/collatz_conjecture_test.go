package collatzconjecture

import (
	"testing"
)

const targetTestVersion = 1

var testCases = []struct {
	description string
	input       int
	expectError bool
	expected    int
}{
	{
		description: "zero steps for one",
		input:       1,
		expectError: false,
		expected:    0,
	},
	{
		description: "divide if even",
		input:       16,
		expectError: false,
		expected:    4,
	},
	{
		description: "even and old steps",
		input:       12,
		expectError: false,
		expected:    9,
	},
	{
		description: "large number of even and odd steps",
		input:       1000000,
		expectError: false,
		expected:    152,
	},
	{
		description: "zero is an error",
		input:       0,
		expectError: true,
	},
	{
		description: "negative value is an error",
		input:       -15,
		expectError: true,
	},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestCollatzConjecture(t *testing.T) {
	for _, testCase := range testCases {
		steps, err := CollatzConjecture(testCase.input)
		if testCase.expectError {
			if err == nil {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected an error, got %v",
					testCase.description, testCase.input, steps)
			}
		} else {
			if steps != testCase.expected {
				t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected %v, got %v",
					testCase.description, testCase.input, testCase.expected, steps)
			}
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkCollatzConjecture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			CollatzConjecture(testCase.input)
		}
	}
}
