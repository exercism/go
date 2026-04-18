package series

import (
	"slices"
	"testing"
)

func TestUnsafeFirst(t *testing.T) {
	for _, tc := range simpleTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := UnsafeFirst(tc.digits, tc.input); got != tc.expected {
				t.Fatalf("UnsafeFirst(%d, %q) = %q, want: %q", tc.digits, tc.input, got, tc.expected)
			}
		})
	}
}

func TestAll(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := All(tc.digits, tc.input); !slices.Equal(got, tc.expected) {
				t.Fatalf("All(%d, %q)\n got: %v, want: %v", tc.digits, tc.input, got, tc.expected)
			}
		})
	}
}
