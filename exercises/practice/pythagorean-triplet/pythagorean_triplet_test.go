package pythagorean

import (
	"reflect"
	"testing"
)

var rangeTests = []struct {
	min, max int
	ts       []Triplet
}{
	{1, 10, []Triplet{{3, 4, 5}, {6, 8, 10}}},
	{11, 20, []Triplet{{12, 16, 20}}},
}

func TestRange(t *testing.T) {
	for _, test := range rangeTests {
		ts := Range(test.min, test.max)
		if !reflect.DeepEqual(ts, test.ts) {
			t.Fatalf("Range(%d, %d) = %v, want %v",
				test.min, test.max, ts, test.ts)
		}
	}
}

var sumTests = []struct {
	sum int
	ts  []Triplet
}{
	{180, []Triplet{{18, 80, 82}, {30, 72, 78}, {45, 60, 75}}},
	{1000, []Triplet{{200, 375, 425}}},
}

func TestSum(t *testing.T) {
	for _, test := range sumTests {
		ts := Sum(test.sum)
		if !reflect.DeepEqual(ts, test.ts) {
			t.Fatalf("Sum(%d) = %v, want %v",
				test.sum, ts, test.ts)
		}
	}
}

func BenchmarkRange(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Range(1, 100)
	}
}

func BenchmarkSum(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Sum(1000)
	}
}
