package accumulate

import (
	"fmt"
	"strings"
	"testing"
)

func echo(c string) string {
	return c
}

func capitalize(word string) string {
	return strings.Title(word)
}

var tests = []struct {
	expected    []string
	given       []string
	converter   func(string) string
	description string
}{
	{[]string{}, []string{}, echo, "echo"},
	{[]string{"echo", "echo", "echo", "echo"}, []string{"echo", "echo", "echo", "echo"}, echo, "echo"},
	{[]string{"First", "Letter", "Only"}, []string{"first", "letter", "only"}, capitalize, "capitalize"},
	{[]string{"HELLO", "WORLD"}, []string{"hello", "world"}, strings.ToUpper, "strings.ToUpper"},
}

func TestAccumulate(t *testing.T) {
	for _, test := range tests {
		actual := Accumulate(test.given, test.converter)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("Accumulate(%q, %q): expected %q, actual %q", test.given, test.description, test.expected, actual)
		}
	}
}

func BenchmarkAccumulate(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, test := range tests {
			Accumulate(test.given, test.converter)
		}

	}
}
