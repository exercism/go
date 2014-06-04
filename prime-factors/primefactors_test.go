package prime

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	input    int
	expected []int
}{
	{1, []int{}},
	{2, []int{2}},
	{3, []int{3}},
	{4, []int{2, 2}},
	{6, []int{2, 3}},
	{8, []int{2, 2, 2}},
	{9, []int{3, 3}},
	{27, []int{3, 3, 3}},
	{625, []int{5, 5, 5, 5}},
	{901255, []int{5, 17, 23, 461}},
	{93819012551, []int{11, 9539, 894119}},
}

func TestPrimeFactors(t *testing.T) {
	for _, test := range tests {
		actual := Factors(test.input)
		sort.Ints(actual)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("prime.Factors(%d) = %v; expected %v", test.input, actual, test.expected)
		}
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for _, test := range tests {
		for i := 0; i < b.N; i++ {
			Factors(test.input)
		}
	}
}
