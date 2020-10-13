package resistorcolor

import (
	"testing"
)

func TestResistorColor(t *testing.T) {
	for _, test := range stringTestCases {
		actual := ColorCode(test.input)
		if actual != test.expected {
			t.Errorf("Resistor color test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func TestColors(t *testing.T) {
	actual := Colors()
	if actual != colorsExpected {
		t.Errorf("Color list test expected [%s], actual [%s]", colorsExpected, actual)
	}
}

func BenchmarkResistorColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			ColorCode(test.input)
		}
	}
}
