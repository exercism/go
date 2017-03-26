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

func determineExpected(expected interface{}) (valid bool, list []int) {
	_, isError := expected.(int)
	if isError {
		return false, nil
	}
	ilist, isIntList := expected.([]interface{})
	if isIntList {
		list = make([]int, 0)
		for _, iv := range ilist {
			v, isInt := iv.(int)
			if isInt {
				list = append(list, v)
			}
		}
		return true, list
	}
	return false, nil
}

func TestChange(t *testing.T) {
	for _, tc := range testCases {
		valid, expectedChange := determineExpected(tc.expected)
		actual, err := Change(tc.coins, tc.target)
		if valid {
			if err != nil {
				t.Fatalf("%s : Change(%v, %d): expected %v, got error %s",
					tc.description, tc.coins, tc.target, expectedChange, err)
			} else {
				if !reflect.DeepEqual(actual, expectedChange) {
					t.Fatalf("%s : Change(%v, %d): expected %v, actual %v",
						tc.description, tc.coins, tc.target, expectedChange, actual)
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
	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			Change(tc.coins, tc.target)
		}
	}
}
