package pangram

import (
	"testing"
)

func TestPangram(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsPangram(tc.input); actual != tc.expected {
				t.Fatalf("IsPangram(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkPangram(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsPangram(test.input)
		}
	}
}
