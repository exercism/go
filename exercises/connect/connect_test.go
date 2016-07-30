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
		newLines[i] = strings.Replace(l, " ", "", -1)
	}
	return newLines
}

const targetTestVersion = 2

func TestResultOf(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, tt := range testCases {
		actual, err := ResultOf(prepare(tt.board))
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
		board := prepare(tt.board)
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ResultOf(board)
		}

		b.StopTimer()
	}

}
