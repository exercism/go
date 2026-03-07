package microblog

import "testing"

func TestTruncate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Truncate(tc.input); actual != tc.expected {
				t.Fatalf("Truncate(%q) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkTruncate(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Truncate(tc.input)
		}
	}
}
