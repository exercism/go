package change

import (
	"reflect"
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestChange(t *testing.T) {
	for _, tc := range testCases {
		actual := Change(tc.coins, tc.target)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("%s : Change(%v, %d): expected %v, actual %v",
				tc.description, tc.coins, tc.target, tc.expected, actual)
		} else {
			t.Logf("PASS: %s", tc.description)
		}
	}
}

func BenchmarkChange(b *testing.B) {
	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			Change(tc.coins, tc.target)
		}
	}
}
