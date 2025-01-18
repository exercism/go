package triangle

import (
	"testing"
)

type testCase struct {
	description string
	expected    Kind
	a, b, c     float64
}

var testCases = []testCase{
	{description: "same length (2)", expected: Equ, a: 2, b: 2, c: 2},
	{description: "same length (10)", expected: Equ, a: 10, b: 10, c: 10},
	{description: "b = c = 4", expected: Iso, a: 3, b: 4, c: 4},
	{description: "a = c = 4", expected: Iso, a: 4, b: 3, c: 4},
	{description: "a = b = 4", expected: Iso, a: 4, b: 4, c: 3},
	{description: "a = b = 10", expected: Iso, a: 10, b: 10, c: 2},
	{description: "no sides equal (3, 4, 5)", expected: Sca, a: 3, b: 4, c: 5},
	{description: "no sides equal (10, 11, 12)", expected: Sca, a: 10, b: 11, c: 12},
	{description: "no sides equal (5, 4, 2)", expected: Sca, a: 5, b: 4, c: 2},
	{description: "no sides equal (0.4, 0.6, 0.3)", expected: Sca, a: .4, b: .6, c: .3},
	{description: "no sides equal (5, 4, 6) | 2a=b+c", expected: Sca, a: 5, b: 4, c: 6},
	{description: "no sides equal (6, 4, 5) | 2c=a+b", expected: Sca, a: 6, b: 4, c: 5},
	{description: "all sides zero", expected: NaT, a: 0, b: 0, c: 0},
	{description: "negative length", expected: NaT, a: 3, b: 4, c: -5},
	{description: "not a triangle (1, 1, 3)", expected: NaT, a: 1, b: 1, c: 3},
	{description: "not a triangle (2, 5, 2)", expected: NaT, a: 2, b: 5, c: 2},
	{description: "not a triangle (7, 3, 2)", expected: NaT, a: 7, b: 3, c: 2},
}

// Test that the kinds are not equal to each other.
// If they are equal, then TestKind will return false positives.
func TestKindsNotEqual(t *testing.T) {
	kindsAndNames := []struct {
		kind Kind
		name string
	}{
		{Equ, "Equ"},
		{Iso, "Iso"},
		{Sca, "Sca"},
		{NaT, "NaT"},
	}

	for i, pair1 := range kindsAndNames {
		for j := i + 1; j < len(kindsAndNames); j++ {
			pair2 := kindsAndNames[j]
			if pair1.kind == pair2.kind {
				t.Fatalf("the value of %s should not be equal to the value of %s", pair1.name, pair2.name)
			}
		}
	}
}

func TestKind(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			got := KindFromSides(test.a, test.b, test.c)
			if got != test.expected {
				t.Fatalf("KindFromSides(%v, %v, %v) = %v, want: %v", test.a, test.b, test.c, got, test.expected)
			}
		})
	}
}

func BenchmarkKind(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			KindFromSides(test.a, test.b, test.c)
		}
	}
}
