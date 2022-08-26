package alphametics

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			s, err := Solve(tc.input)
			switch {
			case tc.errorExpected:
				if err == nil {
					t.Fatalf("Solve(%q) expected error, got: %#v", tc.input, s)
				}
			case err != nil:
				t.Fatalf("Solve(%q)\nexpected: %#v\ngot error: %q", tc.input, tc.expected, err)
			case !reflect.DeepEqual(s, tc.expected):
				t.Fatalf("Solve(%q)\ngot: %#v\nwant:%#v", tc.input, s, tc.expected)
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
			Solve(tc.input)
		}
	}
}
