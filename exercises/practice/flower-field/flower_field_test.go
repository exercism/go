package flowerfield

import (
	"testing"
)

type flowerfieldTestCase struct {
	description string
	garden      []string
	expected    []string
}

var flowerfieldTestCases = []flowerfieldTestCase{
	{
		description: "no rows",
		garden:      []string{},
		expected:    []string{},
	},
	{
		description: "no columns",
		garden:      []string{""},
		expected:    []string{""},
	},
	{
		description: "no flowers",
		garden: []string{
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
		description: "garden with only flowers",
		garden: []string{
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
		description: "flower surrounded by spaces",
		garden: []string{
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
		description: "space surrounded by flowers",
		garden: []string{
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
		garden:      []string{" * * "},
		expected:    []string{"1*2*1"},
	},
	{
		description: "horizontal line, flowers at edges",
		garden:      []string{"*   *"},
		expected:    []string{"*1 1*"},
	},
	{
		description: "vertical line",
		garden: []string{
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
		description: "vertical line, flowers at edges",
		garden: []string{
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
		garden: []string{
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
		description: "large garden",
		garden: []string{
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
	for _, tc := range flowerfieldTestCases {
		t.Run(tc.description, func(t *testing.T) {
			got := Annotate(tc.garden)
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
