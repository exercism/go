package allergies

import (
	"fmt"
	"sort"
	"testing"
)

const targetTestVersion = 1

var allergiesTests = []struct {
	expected []string
	input    uint
}{
	{[]string{}, 0},
	{[]string{"eggs"}, 1},
	{[]string{"peanuts"}, 2},
	{[]string{"strawberries"}, 8},
	{[]string{"eggs", "peanuts"}, 3},
	{[]string{"eggs", "shellfish"}, 5},
	{[]string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 248},
	{[]string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 255},
	{[]string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}, 509},
}

func TestAllergies(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, test := range allergiesTests {
		actual := Allergies(test.input)
		sort.Strings(actual)
		sort.Strings(test.expected)
		if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", test.expected) {
			t.Fatalf("FAIL: Allergies(%d): expected %q, actual %q", test.input, test.expected, actual)
		} else {
			t.Logf("PASS: Allergic to %q", test.expected)
		}
	}
}

func BenchmarkAllergies(b *testing.B) {
	b.StopTimer()
	for _, test := range allergicToTests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Allergies(test.i)
		}
		b.StopTimer()
	}
}

var allergicToTests = []struct {
	expected bool
	i        uint
	allergen string
}{
	{false, 0, "peanuts"},
	{false, 0, "cats"},
	{false, 0, "strawberries"},
	{true, 1, "eggs"},
	{true, 5, "eggs"},
}

func TestAllergicTo(t *testing.T) {
	for _, test := range allergicToTests {
		actual := AllergicTo(test.i, test.allergen)
		if actual != test.expected {
			t.Fatalf("FAIL: AllergicTo(%d, %s): expected %t, actual %t", test.i, test.allergen, test.expected, actual)
		} else {
			t.Logf("PASS: AllergicTo(%d, %s) %t", test.i, test.allergen, actual)
		}
	}
}

func BenchmarkAllergicTo(b *testing.B) {
	b.StopTimer()
	for _, test := range allergicToTests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			AllergicTo(test.i, test.allergen)
		}
		b.StopTimer()
	}
}
