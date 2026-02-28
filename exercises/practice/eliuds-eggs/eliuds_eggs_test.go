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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			EggCount(tc.input)
		}
	}
}
