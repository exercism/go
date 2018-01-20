package allergies

import (
	"reflect"
	"testing"
)

func TestAllergies(t *testing.T) {
	for _, test := range listTests {
		if actual := Allergies(test.score); !reflect.DeepEqual(actual, test.expected) && (len(actual) > 0 || len(test.expected) > 0) {
			t.Fatalf("FAIL: Allergies(%d): expected %#v, actual %#v", test.score, test.expected, actual)
		} else {
			t.Logf("PASS: Allergic to %#v", test.expected)
		}
	}
}

func BenchmarkAllergies(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, test := range allergicToTests {
			Allergies(test.score)
		}
	}
}

func TestAllergicTo(t *testing.T) {
	for _, test := range allergicToTests {
		for _, response := range test.expected {
			actual := AllergicTo(test.score, response.substance)
			if actual != response.result {
				t.Fatalf("FAIL: AllergicTo(%d, %s): expected %t, actual %t", test.score, response.substance, response.result, actual)
			} else {
				t.Logf("PASS: AllergicTo(%d, %s) %t", test.score, response.substance, actual)
			}
		}
	}
}

func BenchmarkAllergicTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range allergicToTests {

			for _, response := range test.expected {
				AllergicTo(test.score, response.substance)
			}
		}
	}
}
