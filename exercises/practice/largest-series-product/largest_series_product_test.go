package lsproduct

import "testing"

func TestLargestSeriesProduct(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := LargestSeriesProduct(tc.digits, tc.span)
			if tc.error == "" {
				// we do not expect error
				if err != nil {
					t.Fatalf("LargestSeriesProduct(%q, %d) returned error: %v\nwant: %d", tc.digits, tc.span, err, tc.expected)
				}

				if actual != tc.expected {
					t.Fatalf("LargestSeriesProduct(%q, %d) = %d, want: %d", tc.digits, tc.span, actual, tc.expected)
				}
			} else if err == nil {
				// expect error but got nil
				t.Fatalf("LargestSeriesProduct(%q, %d) = %d, want error: %q", tc.digits, tc.span, actual, tc.error)
			}
		})
	}
}

func BenchmarkLargestSeriesProduct(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			LargestSeriesProduct(test.digits, test.span)
		}
	}
}
