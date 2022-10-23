package bottlesong

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
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

// Title is a copy of strings.Title function of the stdlib.
// The copy is here because strings.Title is deprecated but we still
// want to use this function as the alternative would require us to support
// external dependencies which we don't yet (tracking issue https://github.com/exercism/go/issues/2379).
// Students should still be able to use strings.Title if they want.
// Since this exercise is currently deprecated, this shouldn't matter too much.
func Title(s string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

// Copy of strings.isSeparator function of the stdlib.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
