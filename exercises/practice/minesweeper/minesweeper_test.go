package minesweeper

import (
	"testing"
)

type minesweeperTestCase struct {
	description string
	minefield   []string
	expected    []string
}

var minesweeperTestCases = []minesweeperTestCase{
	{
		description: "no rows",
		minefield:   []string{},
		expected:    []string{},
	},
	{
		description: "no columns",
		minefield:   []string{""},
		expected:    []string{""},
	},
	{
		description: "no mines",
		minefield: []string{
			"   ",
			"   ",
			"   ",
		},
		expected: []string{
			"   ",
			"   ",
			"   ",
		},
	},
	{
		description: "minefield with only mines",
		minefield: []string{
			"***",
			"***",
			"***",
		},
		expected: []string{
			"***",
			"***",
			"***",
		},
	},
	{
		description: "mine surrounded by spaces",
		minefield: []string{
			"   ",
			" * ",
			"   ",
		},
		expected: []string{
			"111",
			"1*1",
			"111",
		},
	},
	{
		description: "space surrounded by mines",
		minefield: []string{
			"***",
			"* *",
			"***",
		},
		expected: []string{
			"***",
			"*8*",
			"***",
		},
	},
	{
		description: "horizontal line",
		minefield:   []string{" * * "},
		expected:    []string{"1*2*1"},
	},
	{
		description: "horizontal line, mines at edges",
		minefield:   []string{"*   *"},
		expected:    []string{"*1 1*"},
	},
	{
		description: "vertical line",
		minefield: []string{
			" ",
			"*",
			" ",
			"*",
			" ",
		},
		expected: []string{
			"1",
			"*",
			"2",
			"*",
			"1",
		},
	},
	{
		description: "vertical line, mines at edges",
		minefield: []string{
			"*",
			" ",
			" ",
			" ",
			"*",
		},
		expected: []string{
			"*",
			"1",
			" ",
			"1",
			"*",
		},
	},
	{
		description: "cross",
		minefield: []string{
			"  *  ",
			"  *  ",
			"*****",
			"  *  ",
			"  *  ",
		},
		expected: []string{
			" 2*2 ",
			"25*52",
			"*****",
			"25*52",
			" 2*2 ",
		},
	},
	{
		description: "large minefield",
		minefield: []string{
			" *  * ",
			"  *   ",
			"    * ",
			"   * *",
			" *  * ",
			"      ",
		},
		expected: []string{
			"1*22*1",
			"12*322",
			" 123*2",
			"112*4*",
			"1*22*2",
			"111111",
		},
	},
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAnnotate(t *testing.T) {
	for _, tc := range minesweeperTestCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Annotate(tc.minefield)
			want := tc.expected
			if !slicesEqual(want, got) {
				t.Fatalf("expected: %v, got: %v", want, got)
			}
		})
	}
}

var benchmarkResult []string

func BenchmarkAnnotate(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	var result []string
	board := []string{
		"1*22*1",
		"12*322",
		" 123*2",
		"112*4*",
		"1*22*2",
		"111111",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = Annotate(board)
	}
	benchmarkResult = result
}
