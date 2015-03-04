package connect

import (
	"strings"
	"testing"
)

var testCases = []struct {
	description string
	lines       []string
	expected    string
}{
	{
		description: "empty board has no winner",
		lines: []string{
			". . . . .",
			" . . . . .",
			"  . . . . .",
			"   . . . . .",
			"    . . . . .",
		},
		expected: "",
	},
	{
		description: "1x1 board, black",
		lines: []string{
			"X",
		},
		expected: "black",
	},
	{
		description: "1x1 board, white",
		lines: []string{
			"O",
		},
		expected: "white",
	},
	{
		description: "convulted path",
		lines: []string{
			". X X . .",
			" X . X . X",
			"  . X . X .",
			"   . X X . .",
			"    O O O O O",
		},
		expected: "black",
	},
	{
		description: "rectangle, white wins",
		lines: []string{
			". O . .",
			" O X X X",
			"  O O O .",
			"   X X O X",
			"    . O X .",
		},
		expected: "white",
	},
	{
		description: "rectangle, black wins",
		lines: []string{
			". O . .",
			" O X X X",
			"  O X O .",
			"   X X O X",
			"    . O X .",
		},
		expected: "black",
	},
	{
		description: "spiral, black wins",
		lines: []string{
			"OXXXXXXXX",
			"OXOOOOOOO",
			"OXOXXXXXO",
			"OXOXOOOXO",
			"OXOXXXOXO",
			"OXOOOXOXO",
			"OXXXXXOXO",
			"OOOOOOOXO",
			"XXXXXXXXO",
		},
		expected: "black",
	},
	{
		description: "spiral, nobody wins",
		lines: []string{
			"OXXXXXXXX",
			"OXOOOOOOO",
			"OXOXXXXXO",
			"OXOXOOOXO",
			"OXOX.XOXO",
			"OXOOOXOXO",
			"OXXXXXOXO",
			"OOOOOOOXO",
			"XXXXXXXXO",
		},
		expected: "",
	},
}

// Simply strip the spaces of all the strings to get a canonical
// input. The spaces are only for readability of the tests.
func prepare(lines []string) []string {
	newLines := make([]string, len(lines))
	for i, l := range lines {
		newLines[i] = strings.Replace(l, " ", "", -1)
	}
	return newLines
}

func TestResultOf(t *testing.T) {
	for _, tt := range testCases {
		actual, err := ResultOf(prepare(tt.lines))
		// We don't expect errors for any of the test cases
		if err != nil {
			t.Fatalf("ResultOf for board %q returned error %q.  Error not expected.",
				tt.description, err)
		}
		if actual != tt.expected {
			t.Fatalf("ResultOf for board %q was expected to return %q but returned %q.",
				tt.description, tt.expected, actual)
		}
	}
}

func BenchmarkResultOf(b *testing.B) {

	b.StopTimer()

	for _, tt := range testCases {
		board := prepare(tt.lines)
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ResultOf(board)
		}

		b.StopTimer()
	}

}
