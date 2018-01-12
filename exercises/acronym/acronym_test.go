package acronym

import (
	"testing"
)

func TestAcronym(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Abbreviate(test.input)
		if actual != test.expected {
			t.Errorf("Acronym test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func BenchmarkAcronym(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Abbreviate(test.input)
		}
	}
}
