package hamming

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Distance(tc.s1, tc.s2)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("Distance(%q, %q) expected error, got: %d", tc.s1, tc.s2, got)
				}
			case err != nil:
				t.Fatalf("Distance(%q, %q) returned error: %v, want: %d", tc.s1, tc.s2, err, tc.want)
			case got != tc.want:
				t.Fatalf("Distance(%q, %q) = %d, want %d", tc.s1, tc.s2, got, tc.want)
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
