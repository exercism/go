//go:build asktoomuch
// +build asktoomuch

package series

import (
	"fmt"
	"testing"
)

func TestAskTooMuch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("first with %d digits in %s", tc.digits, tc.input), func(t *testing.T) {
			defer func() {
				if recover() != nil {
					t.Fatalf("UnsafeFirst(%d, %q) panicked!", tc.digits, tc.input)
				}
			}()
			got := UnsafeFirst(tc.digits, tc.input)
			switch {
			case len(tc.expected) > 0:
				if got != tc.expected[0] {
					t.Fatalf("UnsafeFirst(%d, %q) = %q, want: %q", tc.digits, tc.input, got, tc.expected[0])
				}
			case len(got) != tc.digits:
				t.Fatalf("UnsafeFirst(%d, %q) = %q, but %q doesn't have %d characters", tc.digits, tc.input, got, got, tc.digits)
			default:
				t.Fatalf("UnsafeFirst(%d, %q) = %q, but %q isn't in %q", tc.digits, tc.input, got, got, tc.input)
			}
		})
	}
}
