//go:build first
// +build first

package series

import "testing"

func TestFirst(t *testing.T) {
	for _, tc := range testCases {
		switch got, ok := First(tc.digits, tc.input); {
		case !ok:
			if len(tc.expected) > 0 {
				t.Fatalf("First(%d, %q) returned ok=false, want ok=true", tc.digits, tc.input)
			}
		case len(tc.expected) == 0:
			t.Fatalf("First(%d, %q) = %q, %v | expected ok=false", tc.digits, tc.input, got, ok)
		case got != tc.expected[0]:
			t.Fatalf("First(%d, %q) = %q, want: %q", tc.digits, tc.input, got, tc.expected[0])
		}
	}
}
