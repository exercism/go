package wordsearch

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Solve(tc.words, tc.puzzle)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("Solve(%v,%v) expected error, got:%v", tc.words, tc.puzzle, actual)
				}
			case err != nil:
				t.Fatalf("Solve(%v,%v) returned error: %v, want:%v", tc.words, tc.puzzle, err, tc.expected)
			case !reflect.DeepEqual(actual, tc.expected):
				t.Fatalf("Solve(%v,%v) = %v, want:%v", tc.words, tc.puzzle, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Solve(tc.words, tc.puzzle)
		}
	}
}
