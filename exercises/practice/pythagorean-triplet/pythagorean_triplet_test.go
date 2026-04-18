package pythagorean

import (
	"reflect"
	"testing"
)

type rangeTest struct {
	description string
	min, max    int
	expected    []Triplet
}

var rangeTests = []rangeTest{
	{
		description: "Triplets in range 1..10",
		min:         1,
		max:         10,
		expected:    []Triplet{{3, 4, 5}, {6, 8, 10}},
	},
	{
		description: "Triplets in range 11..20",
		min:         11,
		max:         20,
		expected:    []Triplet{{12, 16, 20}},
	},
}

func TestRange(t *testing.T) {
	for _, tc := range rangeTests {
		t.Run(tc.description, func(t *testing.T) {
			got := Range(tc.min, tc.max)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Range(%d, %d) = %v, want: %v", tc.min, tc.max, got, tc.expected)
			}
		})
	}
}

type sumTest struct {
	description string
	sum         int
	expected    []Triplet
}

var sumTests = []sumTest{
	{
		description: "Triplets with sum 180",
		sum:         180,
		expected:    []Triplet{{18, 80, 82}, {30, 72, 78}, {45, 60, 75}},
	},
	{
		description: "Triplets with sum 1000",
		sum:         1000,
		expected:    []Triplet{{200, 375, 425}},
	},
}

func TestSum(t *testing.T) {
	for _, tc := range sumTests {
		t.Run(tc.description, func(t *testing.T) {
			got := Sum(tc.sum)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Sum(%d) = %v, want: %v", tc.sum, got, tc.expected)
			}
		})
	}
}

func BenchmarkRange(b *testing.B) {
	for range b.N {
		Range(1, 100)
	}
}

func BenchmarkSum(b *testing.B) {
	for range b.N {
		Sum(1000)
	}
}
