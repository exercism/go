package proverb

import (
	"testing"
)

func TestProverb(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Proverb(test.input)
		if actual != test.expected {
			t.Errorf("Proverb test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkProverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Proverb(test.input)
		}
	}
}
