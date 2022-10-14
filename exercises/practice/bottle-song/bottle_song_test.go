package bottlesong

import (
	"fmt"
	"testing"
)

func TestRecite(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Recite(tc.input.startBottles, tc.input.takeDown)
			if !equal(actual, tc.expected) {
				t.Errorf("Recite(%d, %d) = %q, want: %q", tc.input.startBottles, tc.input.takeDown, actual, tc.expected)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	if len(a) == 0 && len(b) == 0 {
		return true
	}

	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
