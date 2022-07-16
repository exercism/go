package connect

import (
	"strings"
	"testing"
)

// Simply strip the spaces of all the strings to get a canonical
// input. The spaces are only for readability of the tests.
func prepare(lines []string) []string {
	newLines := make([]string, len(lines))
	for i, l := range lines {
		newLines[i] = strings.ReplaceAll(l, " ", "")
	}
	return newLines
}

func TestResultOf(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := ResultOf(prepare(tc.board))
			// We don't expect errors for any of the test cases
			if err != nil {
				t.Errorf("ResultOf() returned error %v\nboard: \n%s\nwant: %q", err, strings.Join(tc.board, "\n"), tc.expected)
			} else if actual != tc.expected {
				t.Errorf("ResultOf() returned wrong result \nboard: \n%s\ngot: %q\nwant: %q", strings.Join(tc.board, "\n"), actual, tc.expected)
			}
		})
	}
}

func BenchmarkResultOf(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	b.StopTimer()

	for _, tt := range testCases {
		board := prepare(tt.board)
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ResultOf(board)
		}

		b.StopTimer()
	}

}
