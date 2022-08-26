package sublist

import (
	"testing"
)

func TestSublist(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Sublist(tc.listOne, tc.listTwo); actual != tc.expected {
				t.Fatalf("Sublist(%#v,%#v) = %v, want: %v", tc.listOne, tc.listTwo, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSublist(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sublist(tc.listOne, tc.listTwo)
		}
	}
}
