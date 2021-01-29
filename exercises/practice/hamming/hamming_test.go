package hamming

import "testing"

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		got, err := Distance(tc.s1, tc.s2)
		if tc.expectError {
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("Distance(%q, %q); expected error, got nil.",
					tc.s1, tc.s2)
			}
		} else {
			// we do not expect error
			if err != nil {
				t.Fatalf("Distance(%q, %q) returned unexpected error: %v",
					tc.s1, tc.s2, err)
			}
			if got != tc.want {
				t.Fatalf("Distance(%q, %q) = %d, want %d.",
					tc.s1, tc.s2, got, tc.want)
			}

		}
	}
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			// ignoring errors and results because we're just timing function execution
			_, _ = Distance(tc.s1, tc.s2)
		}
	}
}
