package forth

// API:
// func Forth([]string) ([]int, error)
//

import (
	"reflect"
	"testing"
)

func TestForth(t *testing.T) {
	for _, tg := range testGroups {
		for _, tc := range tg.tests {
			if v, err := Forth(tc.input); err == nil {
				var _ error = err
				if tc.expected == nil {
					t.Fatalf("FAIL: %s | %s\n\tForth(%#v) expected an error, got %v",
						tg.group, tc.description, tc.input, v)
				} else if !reflect.DeepEqual(v, tc.expected) {
					t.Fatalf("FAIL: %s | %s\n\tForth(%#v) expected %v, got %v",
						tg.group, tc.description, tc.input, tc.expected, v)
				}
			} else if tc.expected != nil {
				t.Fatalf("FAIL: %s | %s\n\tForth(%#v) expected %v, got an error: %q",
					tg.group, tc.description, tc.input, tc.expected, err)
			}
			t.Logf("PASS: %s | %s", tg.group, tc.description)
		}
	}
}

func BenchmarkForth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tg := range testGroups {
			for _, tc := range tg.tests {
				Forth(tc.input)
			}
		}
	}
}
