package scrabble

import "testing"

const testVersion = 2

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
	{"oxyphenbutazone", 41},
	{"alacrity", 13},
}

func TestScore(t *testing.T) {
	for _, test := range tests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}

	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v.", TestVersion, testVersion)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Score(test.input)
		}
	}
}
