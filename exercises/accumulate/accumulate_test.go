package accumulate

import (
	"fmt"
	"strings"
	"testing"
)

const targetTestVersion = 1

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

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
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
	b.StopTimer()
	for _, test := range tests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Accumulate(test.given, test.converter)
		}

		b.StopTimer()
	}
}
