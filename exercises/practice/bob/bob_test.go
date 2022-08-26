package bob

import "testing"

func TestHey(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Hey(tc.input)
			if actual != tc.expected {
				t.Fatalf("Hey(%q) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkHey(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Hey(tc.input)
		}
	}
}
