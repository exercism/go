package change

import (
	"reflect"
	"testing"
)

func TestChange(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Change(tc.coins, tc.target)
			if tc.valid {
				if err != nil {
					t.Fatalf("Change(%v, %d): expected %v, got error %v", tc.coins, tc.target, tc.expectedChange, err)
				} else if !reflect.DeepEqual(actual, tc.expectedChange) {
					t.Fatalf("Change(%v, %d): expected %#v, actual %#v", tc.coins, tc.target, tc.expectedChange, actual)
				}
			} else {
				if err == nil {
					t.Fatalf("Change(%v, %d): expected error, got %v", tc.coins, tc.target, actual)
				}
			}
		})
	}
}

func BenchmarkChange(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Change(tc.coins, tc.target)
		}
	}
}
