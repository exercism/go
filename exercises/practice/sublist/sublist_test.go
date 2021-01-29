package sublist

import (
	"testing"
)

func TestSublist(t *testing.T) {
	for _, tc := range testCases {
		if actual := Sublist(tc.listOne, tc.listTwo); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkSublist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sublist(tc.listOne, tc.listTwo)
		}
	}
}
