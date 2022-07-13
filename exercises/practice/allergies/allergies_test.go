package allergies

import (
	"reflect"
	"testing"
)

func TestAllergies(t *testing.T) {
	for _, test := range listTests {
		t.Run(test.description, func(t *testing.T) {
			if actual := Allergies(test.score); !sameSliceElements(actual, test.expected) {
				t.Fatalf("FAIL: Allergies(%d): expected %#v, actual %#v", test.score, test.expected, actual)
			} else {
				t.Logf("PASS: Allergic to %#v", test.expected)
			}
		})
	}
}

func BenchmarkAllergies(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {

		for _, test := range allergicToTests {
			Allergies(test.input.score)
		}
	}
}

func TestAllergicTo(t *testing.T) {
	for _, test := range allergicToTests {
		t.Run(test.description, func(t *testing.T) {
			actual := AllergicTo(test.input.score, test.input.allergen)
			if actual != test.expected {
				t.Fatalf("FAIL: AllergicTo(%d, %s): expected %t, actual %t", test.input.score, test.input.allergen, test.expected, actual)
			} else {
				t.Logf("PASS: AllergicTo(%d, %s) %t", test.input.score, test.input.allergen, actual)
			}
		})
	}
}

func BenchmarkAllergicTo(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range allergicToTests {
			AllergicTo(test.input.score, test.input.allergen)
		}
	}
}

// stringSet is a set of strings
type stringSet map[string]bool

// sameSliceElements checks if the slices have the same number of elements
// regardless of their order
func sameSliceElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	return reflect.DeepEqual(sliceSet(a), sliceSet(b))
}

// sliceSet creates a new stringSet from a slice of strings
func sliceSet(list []string) stringSet {
	set := make(stringSet)

	for _, elem := range list {
		set[elem] = true
	}

	return set
}
