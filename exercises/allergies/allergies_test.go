package allergies

import (
	"reflect"
	"testing"
)

func TestAllergies(t *testing.T) {
	allergicToTests, listTests := getTests()
	testAllergicTo(t, allergicToTests)
	testAllergies(t, listTests)
}

func testAllergies(t *testing.T, listTests []lTest) {
	for _, test := range listTests {
		if actual := Allergies(test.score); !reflect.DeepEqual(actual, test.expected) && (len(actual) > 0 || len(test.expected) > 0) {
			t.Fatalf("FAIL: Allergies(%d): expected %#v, actual %#v", test.score, test.expected, actual)
		} else {
			t.Logf("PASS: Allergic to %#v", test.expected)
		}
	}
}

func testAllergicTo(t *testing.T, allergicToTests []aTest) {
	for _, test := range allergicToTests {
		actual := AllergicTo(test.score, test.item)
		if actual != test.expected {
			t.Fatalf("FAIL: AllergicTo(%d, %s): expected %t, actual %t", test.score, test.item, test.expected, actual)
		} else {
			t.Logf("PASS: AllergicTo(%d, %s) %t", test.score, test.item, actual)
		}
	}
}

func BenchmarkAllergies(b *testing.B) {
	allergicToTests, listTests := getTests()
	benchmarkAllergicTo(b, allergicToTests)
	benchmarkAllergies(b, listTests)
}

func benchmarkAllergies(b *testing.B, listTests []lTest) {
	for i := 0; i < b.N; i++ {
		for _, test := range listTests {
			Allergies(test.score)
		}
	}
}

func benchmarkAllergicTo(b *testing.B, allergicToTests []aTest) {
	for i := 0; i < b.N; i++ {
		for _, test := range allergicToTests {
			AllergicTo(test.score, test.item)
		}
	}
}
