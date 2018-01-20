package change

import (
	"reflect"
	"testing"
)

func TestChange(t *testing.T) {
	for _, tc := range testCases {
		actual, err := Change(tc.coins, tc.target)
		if tc.valid {
			if err != nil {
				t.Fatalf("%s : Change(%v, %d): expected %v, got error %s",
					tc.description, tc.coins, tc.target, tc.expectedChange, err.Error())
			} else {
				if !reflect.DeepEqual(actual, tc.expectedChange) {
					t.Fatalf("%s : Change(%v, %d): expected %#v, actual %#v",
						tc.description, tc.coins, tc.target, tc.expectedChange, actual)
				} else {
					t.Logf("PASS: %s", tc.description)
				}
			}
		} else {
			if err == nil {
				t.Fatalf("%s : Change(%v, %d): expected error, got %v",
					tc.description, tc.coins, tc.target, actual)
			} else {
				t.Logf("PASS: %s", tc.description)
			}
		}
	}
}

func BenchmarkChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Change(tc.coins, tc.target)
		}
	}
}
