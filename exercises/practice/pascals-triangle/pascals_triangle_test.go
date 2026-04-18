package pascalstriangle

import (
	"fmt"
	"slices"
	"testing"
)

func TestTriangle(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := Triangle(tc.rows); !slices.EqualFunc(got, tc.expected, slices.Equal) {
				help := getHelp(got, tc.expected)
				t.Fatalf("Triangle(%d)\nhelp: %s\ncomplete got: %s\ncomplete want: %s\n", tc.rows, help, format(got), format(tc.expected))
			}
		})
	}
}

func getHelp(got, want [][]int) string {
	if len(got) != len(want) {
		return fmt.Sprintf("expected %d rows, got: %d", len(want), len(got))
	}
	for i, gotLine := range got {
		if !slices.Equal(gotLine, want[i]) {
			return fmt.Sprintf("first difference in row with index: %d\n got: %v\nwant: %v", i, gotLine, want[i])
		}
	}
	return ""
}

func format(t [][]int) (s string) {
	for _, r := range t {
		s = fmt.Sprintf("%s\n%v", s, r)
	}
	return s
}

// BenchmarkPascalsTriangleFixed will generate Pascals Triangles against the
// solution using triangles of fixed size 20.
func BenchmarkPascalsTriangleFixed(b *testing.B) {
	for range b.N {
		Triangle(len(testCases))
	}
}

// BenchmarkPascalsTriangleIncreasing will generate Pascals Triangles against the
// solution using triangles of an increasingly larger size from 1 to 20.
func BenchmarkPascalsTriangleIncreasing(b *testing.B) {
	for range b.N {
		for x := range len(testCases) {
			Triangle(x)
		}
	}
}
