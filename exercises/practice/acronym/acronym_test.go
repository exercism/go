package acronym

import (
	"testing"
)

func TestAcronym(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Abbreviate(tc.input)
			if actual != tc.expected {
				t.Errorf("Abbreviate(%q) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkAcronym(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Abbreviate(test.input)
		}
	}
}
