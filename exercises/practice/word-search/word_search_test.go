package wordsearch

import (
	"reflect"
	"testing"
)

// Define a function Solve(words []string, puzzle []string) (map[string][2][2]int, error).

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		actual, err := Solve(tc.words, tc.puzzle)
		switch {
		case err != nil:
			var _ error = err
			if !tc.expectError {
				t.Fatalf("FAIL: %s\nExpected %#v\nGot error: %v", tc.description, tc.expected, err)
			}
		case tc.expectError:
			t.Fatalf("FAIL: %s\nExpected error\nGot %v", tc.description, actual)
		case !reflect.DeepEqual(actual, tc.expected):
			t.Fatalf("FAIL: %s\nExpected %v,\nGot %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Solve(tc.words, tc.puzzle)
		}
	}
}
