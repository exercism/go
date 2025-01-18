package pythagorean

import (
	"fmt"
	"reflect"
	"testing"
)

var rangeTests = []struct {
	min, max int
	expected []Triplet
}{
	{
		min:      1,
		max:      10,
		expected: []Triplet{{3, 4, 5}, {6, 8, 10}},
	},
	{
		min:      11,
		max:      20,
		expected: []Triplet{{12, 16, 20}},
	},
}

func TestRange(t *testing.T) {
	for _, tc := range rangeTests {
		t.Run(fmt.Sprintf("Triplets in Range %d-%d", tc.min, tc.max), func(t *testing.T) {
			got := Range(tc.min, tc.max)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Range(%d, %d) = %v, want: %v", tc.min, tc.max, got, tc.expected)
			}
		})
	}
}

var sumTests = []struct {
	sum      int
	expected []Triplet
}{
	{
		sum:      180,
		expected: []Triplet{{18, 80, 82}, {30, 72, 78}, {45, 60, 75}},
	},
	{
		sum:      1000,
		expected: []Triplet{{200, 375, 425}},
	},
}

func TestSum(t *testing.T) {
	for _, tc := range sumTests {
		t.Run(fmt.Sprintf("Triplets with perimeter %d", tc.sum), func(t *testing.T) {
			got := Sum(tc.sum)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Sum(%d) = %v, want: %v", tc.sum, got, tc.expected)
			}
		})
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
