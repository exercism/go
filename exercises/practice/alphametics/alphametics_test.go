package alphametics

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		s, err := Solve(tc.input)
		switch {
		case tc.errorExpected:
			if err == nil {
				t.Fatalf("FAIL: %s\nSolve(%q)\nExpected error\nActual: %#v",
					tc.description, tc.input, s)
			}
		case err != nil:
			t.Fatalf("FAIL: %s\nSolve(%q)\nExpected: %#v\nGot error: %q",
				tc.description, tc.input, tc.expected, err)
		case !reflect.DeepEqual(s, tc.expected):
			t.Fatalf("FAIL: %s\nSolve(%q)\nExpected: %#v\nActual: %#v",
				tc.description, tc.input, tc.expected, s)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSolve(b *testing.B) {
	if testing.Short() {
		t.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Solve(tc.input)
		}
	}
}
