package zebra

import "testing"

func TestSolvePuzzle(t *testing.T) {
	expected := Solution{DrinksWater: "Norwegian", OwnsZebra: "Japanese"}
	actual := SolvePuzzle()
	if expected.DrinksWater != actual.DrinksWater {
		t.Fatalf("FAILED: The resident who drinks water should be %q. Actual: %q",
			expected.DrinksWater, actual.DrinksWater)
	}
	if expected.OwnsZebra != actual.OwnsZebra {
		t.Fatalf("FAILED: The resident who owns the zebra should be %q. Actual: %q",
			expected.OwnsZebra, actual.OwnsZebra)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePuzzle()
	}
}
