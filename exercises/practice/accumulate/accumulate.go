package accumulate

import (
	"fmt"
	"strings"
	"testing"
)

func TestAccumulate(t *testing.T) {
	tests := []struct {
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
