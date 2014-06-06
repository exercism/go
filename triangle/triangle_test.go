package triangle

import (
	"math"
	"testing"
)

type testCase struct {
	want    Kind
	a, b, c float64
}

// basic test cases
var testData = []testCase{
	{Equ, 2, 2, 2},    // same length
	{Equ, 10, 10, 10}, // a little bigger
	{Iso, 3, 4, 4},    // last two sides equal
	{Iso, 4, 3, 4},    // first and last sides equal
	{Iso, 4, 4, 3},    // first two sides equal
	{Iso, 10, 10, 2},  // again
	{Sca, 3, 4, 5},    // no sides equal
	{Sca, 10, 11, 12}, // again
	{Sca, 5, 4, 2},    // descending order
	{Sca, .4, .6, .3}, // small sides
	{NaT, 0, 0, 0},    // zero length
	{NaT, 3, 4, -5},   // negative length
	{NaT, 1, 1, 3},    // fails triangle inequality
	{NaT, 2, 4, 2},    // another
	{NaT, 7, 3, 2},    // another
}

// generate cases with NaN and Infs, append to basic cases
func init() {
	nan := math.NaN()
	pinf := math.Inf(1)
	ninf := math.Inf(-1)
	nf := make([]testCase, 4*4*4)
	i := 0
	for _, a := range []float64{3, nan, pinf, ninf} {
		for _, b := range []float64{4, nan, pinf, ninf} {
			for _, c := range []float64{5, nan, pinf, ninf} {
				nf[i] = testCase{NaT, a, b, c}
				i++
			}
		}
	}
	testData = append(testData, nf[1:]...)
}

func TestKind(t *testing.T) {
	for _, test := range testData {
		got := KindFromSides(test.a, test.b, test.c)
		if got != test.want {
			t.Fatalf("Triangle with sides, %g, %g, %g = %v, want %v",
				test.a, test.b, test.c, got, test.want)
		}
	}
}

func BenchmarkKind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			KindFromSides(test.a, test.b, test.c)
		}
	}
}
