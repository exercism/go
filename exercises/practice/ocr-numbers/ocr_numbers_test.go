package ocrnumbers

import (
	"slices"
	"testing"
)

var _ = recognizeDigit // step 1.

func TestRecognize(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got, err := Recognize(tc.input); tc.errorMsg != "" {
				if err == nil {
					t.Fatalf("Recognize(%q) = %#v, want: error %v", tc.input, got, tc.errorMsg)
				} else if err.Error() != tc.errorMsg {
					t.Fatalf("Recognize(%q) = error %v, want: error %v", tc.input, err, tc.errorMsg)
				}
			} else if !slices.Equal(got, tc.expected) {
				t.Fatalf("Recognize(%q) = %#v, want: %#v", tc.input, got, tc.expected)
			}
		})
	}
}
