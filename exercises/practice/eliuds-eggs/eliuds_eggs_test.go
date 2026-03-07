package eliudseggs

import "testing"

func TestEggCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := EggCount(tc.input)
			if got != tc.expected {
				t.Fatalf("EggCount(%d) = %d, want: %d", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkEggCount(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			EggCount(tc.input)
		}
	}
}
