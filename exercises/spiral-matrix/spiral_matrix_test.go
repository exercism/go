package spiralmatrix

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	description string
	input       int
	expected    [][]int
}{
	{
		description: "empty spiral",
		input:       0,
		expected:    [][]int{},
	},
	{
		description: "trivial spiral",
		input:       1,
		expected: [][]int{
			[]int{1},
		},
	},
	{
		description: "spiral of size 2",
		input:       2,
		expected: [][]int{
			[]int{1, 2},
			[]int{4, 3},
		},
	},
	{
		description: "spiral of size 3",
		input:       3,
		expected: [][]int{
			[]int{1, 2, 3},
			[]int{8, 9, 4},
			[]int{7, 6, 5},
		},
	},
	{
		description: "spiral of size 4",
		input:       4,
		expected: [][]int{
			[]int{1, 2, 3, 4},
			[]int{12, 13, 14, 5},
			[]int{11, 16, 15, 6},
			[]int{10, 9, 8, 7},
		},
	},
}

func TestSpiralMatrix(t *testing.T) {
	for _, testCase := range testCases {
		matrix := SpiralMatrix(testCase.input)
		if !reflect.DeepEqual(matrix, testCase.expected) {
			t.Fatalf("FAIL: %s\n\tSpiralMatrix(%v)\nexpected: %v\ngot     : %v",
				testCase.description, testCase.input, testCase.expected, matrix)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkSpiralMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			SpiralMatrix(testCase.input)
		}
	}
}
