package strand

import (
	"testing"
)

type rnaTest struct {
	input    string
	expected string
}

var rnaTests = []rnaTest{
	{"C", "G"},
	{"G", "C"},
	{"A", "U"},
	{"T", "A"},
	{"ACGTGGTCTTAA", "UGCACCAGAAUU"},
}

func TestRnaTranscription(t *testing.T) {
	for _, test := range rnaTests {
		actual := ToRna(test.input)
		if actual != test.expected {
			t.Errorf("ToRna(%s): expected %s, actual %s", test.input, test.expected, actual)
		}
	}
}

func BenchmarkRnaTranscription(b *testing.B) {
	b.StopTimer()
	for _, test := range rnaTests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ToRna(test.input)
		}
		b.StopTimer()
	}
}
