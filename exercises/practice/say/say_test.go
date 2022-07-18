package say

import (
	"testing"
)

func TestSay(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, ok := Say(tc.input)
			switch {
			case tc.expectError:
				if ok {
					t.Fatalf("Say(%d) expected error, got: %q", tc.input, actual)
				}
			case !ok:
				t.Fatalf("Say(%d) got ok:%t, want: %q", tc.input, ok, tc.expected)
			case actual != tc.expected:
				t.Fatalf("Say(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSay(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Say(tc.input)
		}
	}
}
