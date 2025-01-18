package forth

import (
	"reflect"
	"testing"
)

func TestForth(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			v, err := Forth(tc.input)
			if err == nil {
				if tc.expected == nil {
					t.Fatalf("Forth(%#v) expected an error, got %v", tc.input, v)
				} else if !reflect.DeepEqual(v, tc.expected) {
					t.Fatalf("Forth(%#v) expected %v, got %v", tc.input, tc.expected, v)
				}
			} else if tc.expected != nil {
				t.Fatalf("Forth(%#v) expected %v, got an error: %q", tc.input, tc.expected, err)
			}
		})
	}
}

func BenchmarkForth(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Forth(tc.input)
		}
	}
}
