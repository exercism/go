package accumulate

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

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

func echo(c string) string {
	return c
}

func capitalize(word string) string {
	return Title(word)
}

var tests = []struct {
	expected    []string
	given       []string
	converter   func(string) string
	description string
}{
	{[]string{}, []string{}, echo, "echo"},
	{
		[]string{"echo", "echo", "echo", "echo"},
		[]string{"echo", "echo", "echo", "echo"},
		echo,
		"echo",
	},
	{
		[]string{"First", "Letter", "Only"},
		[]string{"first", "letter", "only"},
		capitalize,
		"capitalize",
	},
	{
		[]string{"HELLO", "WORLD"},
		[]string{"hello", "world"},
		strings.ToUpper,
		"strings.ToUpper",
	},
}

func TestAccumulate(t *testing.T) {
	for _, test := range tests {
		in := make([]string, len(test.given))
		copy(in, test.given)
		actual := Accumulate(in, test.converter)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("Accumulate(%q, %q): expected %q, actual %q",
				test.given, test.description, test.expected, actual)
		}

		// check in place replacement
		for i, s := range in {
			if test.given[i] != s {
				t.Fatalf("Accumulate should return a new slice")
			}
		}
	}
}

func BenchmarkAccumulate(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Accumulate(test.given, test.converter)
		}
	}
}
