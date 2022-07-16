package collatzconjecture

import (
	"testing"
)

func TestCollatzConjecture(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := CollatzConjecture(testCase.input)
			if testCase.expectError {
				if err == nil {
					t.Errorf("CollatzConjecture(%v) expected an error, got %v", testCase.input, actual)
				}
			} else {
				if err != nil {
					t.Errorf("CollatzConjecture(%v) returns unexpected error %v", testCase.input, err)
				} else if actual != testCase.expected {
					t.Errorf("CollatzConjecture(%v) expected %v, got %v", testCase.input, testCase.expected, actual)
				}
			}
		})
	}
}

func BenchmarkCollatzConjecture(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			CollatzConjecture(testCase.input)
		}
	}
}
