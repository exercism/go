package lsproduct

import "testing"

func TestLargestSeriesProduct(t *testing.T) {
	for _, test := range tests {
		p, err := LargestSeriesProduct(test.digits, test.span)
		if test.ok {
			// we do not expect error
			if err != nil {
				t.Fatalf("LargestSeriesProduct(%s, %d) returned error %q.  "+
					"Error not expected.",
					test.digits, test.span, err)
			}

			if int64(p) != test.product {
				t.Fatalf("LargestSeriesProduct(%s, %d) = %d, want %d",
					test.digits, test.span, p, test.product)
			}
		} else { // expect error
			// check if err is of error type
			var _ error = err

			// we expect error
			if err == nil {
				t.Fatalf("LargestSeriesProduct(%s, %d) = %d, %v."+
					"  Expected error got nil",
					test.digits, test.span, p, err)
			}
		}
	}
}

func BenchmarkLargestSeriesProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			LargestSeriesProduct(test.digits, test.span)
		}
	}
}
