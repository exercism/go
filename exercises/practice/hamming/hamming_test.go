package hamming

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Distance(tc.s1, tc.s2)
			if tc.expectError {
				if err == nil {
					t.Errorf("Distance(%q, %q) expected error, got: %d", tc.s1, tc.s2, got)
				}
			} else {
				if err != nil {
					t.Errorf("Distance(%q, %q) returned unexpected error: %v", tc.s1, tc.s2, err)
				}
				if got != tc.want {
					t.Errorf("Distance(%q, %q) = %d, want %d", tc.s1, tc.s2, got, tc.want)
				}
			}
		})
	}
}

func BenchmarkHamming(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			_, _ = Distance(tc.s1, tc.s2)
		}
	}
}
