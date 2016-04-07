package scrabble

import "testing"

const targetTestVersion = 4

var tests = []struct {
	input    string
	expected int
}{
	{"", 0},
	{" \t\n", 0},
	{"a", 1},
	{"f", 4},
	{"street", 6},
	{"quirky", 22},
	{"OXYPHENBUTAZONE", 41},
	{"alacrity", 13},
}

func TestScore(t *testing.T) {
	for _, test := range tests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}

	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Score(test.input)
		}
	}
}
