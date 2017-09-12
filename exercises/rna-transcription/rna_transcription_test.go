package strand

import "testing"

func TestRNATranscription(t *testing.T) {
	for _, test := range rnaTests {
		if actual := ToRNA(test.input); actual != test.expected {
			t.Errorf("ToRNA(%s): %s, expected %s",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkRNATranscription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range rnaTests {
			ToRNA(test.input)
		}
	}
}
