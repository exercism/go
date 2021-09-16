package zebra

import "testing"

func TestSolvePuzzle(t *testing.T) {
	expected := Solution{DrinksWater: "Norwegian", OwnsZebra: "Japanese"}
	actual := SolvePuzzle()
	if expected != actual {
		t.Fatalf("FAILED:\nExpected: %#v\nActual: %#v",
			expected, actual)
	}
}

func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		SolvePuzzle()
	}
}
