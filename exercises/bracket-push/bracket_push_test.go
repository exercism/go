package brackets

import (
	"testing"
)

const targetTestVersion = 3

var testCases = []struct {
	input    string
	expected bool
}{
	{
		input:    "",
		expected: true,
	},
	{
		input:    "{}",
		expected: true,
	},
	{
		input:    "{{",
		expected: false,
	},
	{
		input:    "}{",
		expected: false,
	},
	{
		input:    "{}[]",
		expected: true,
	},
	{
		input:    "{[]}",
		expected: true,
	},
	{
		input:    "{[}]",
		expected: false,
	},
	{
		input:    "{[)][]}",
		expected: false,
	},
	{
		input:    "{[]([()])}",
		expected: true,
	},
}

func TestBracket(t *testing.T) {
	for _, tt := range testCases {
		actual, err := Bracket(tt.input)
		// We don't expect errors for any of the test cases
		if err != nil {
			t.Fatalf("Bracket(%q) returned error %q.  Error not expected.", tt.input, err)
		}
		if actual != tt.expected {
			t.Fatalf("Bracket(%q) was expected to return %v but returned %v.",
				tt.input, tt.expected, actual)
		}
	}
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func BenchmarkBracket(b *testing.B) {
	b.StopTimer()
	for _, tt := range testCases {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Bracket(tt.input)
		}
		b.StopTimer()
	}
}
