package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	for _, tc := range testCases {
		if actual := Flatten(tc.input); !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("FAIL: %s\nExpected: %v\nActual: %v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkFlatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Flatten(tc.input)
		}
	}
}
