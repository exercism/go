package swiftscheduling

import "testing"

func TestDeliveryDate(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := DeliveryDate(tc.start, tc.delivery); actual != tc.expected {
				t.Fatalf("DeliveryDate(%q, %q) = %q, want: %q", tc.start, tc.delivery, actual, tc.expected)
			}
		})
	}
}

func BenchmarkDeliveryDate(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			DeliveryDate(tc.start, tc.delivery)
		}
	}
}
