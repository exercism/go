package octal

import (
	"testing"
)

var testCases = []struct {
	input    string
	expected int64
}{
	{"1", 1},
	{"10", 8},
	{"1234567", 342391},
	{"carrot", 0},
}

func TestToDecimal(t *testing.T) {
	for _, test := range testCases {
		actual := ToDecimal(test.input)
		if actual != test.expected {
			t.Fatalf("ToDecimal(%s): expected[%d], actual [%d]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkToDecimal(b *testing.B) {
	b.StopTimer()

	for _, test := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ToDecimal(test.input)
		}

		b.StopTimer()
	}
}
