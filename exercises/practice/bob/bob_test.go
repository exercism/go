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
	for range b.N {
		for _, tc := range testCases {
			Hey(tc.input)
		}
	}
}
