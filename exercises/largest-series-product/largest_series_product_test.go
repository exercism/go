package lsproduct

import "testing"

const targetTestVersion = 3

func TestLargestSeriesProduct(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	for _, test := range tests {
		p, err := LargestSeriesProduct(test.digits, test.span)
		switch {
		case err != nil:
			if test.ok {
				t.Fatalf("LargestSeriesProduct(%s, %d) returned error %q.  "+
					"Error not expected.",
					test.digits, test.span, err)
			}
		case !test.ok:
			t.Fatalf("LargestSeriesProduct(%s, %d) = %d, %v.  Expected error",
				test.digits, test.span, p, err)
		case int64(p) != test.product:
			t.Fatalf("LargestSeriesProduct(%s, %d) = %d, want %d",
				test.digits, test.span, p, test.product)
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
