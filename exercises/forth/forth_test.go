package forth

// API:
// func Forth([]string) ([]int, error)
//

import (
	"reflect"
	"testing"
)

const targetTestVersion = 1

type testCase struct {
	description string
	input       []string
	expected    []int // nil slice indicates error expected.
}

type testcaseSection struct {
	name  string
	tests []testCase
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func runTestCases(t *testing.T, sectionName string, testCases []testCase) {
	for _, tc := range testCases {
		description := sectionName + " - " + tc.description
		if v, err := Forth(tc.input); err == nil {
			var _ error = err
			if tc.expected == nil {
				t.Fatalf("%s\n\tForth(%#v) expected an error, got %v",
					description, tc.input, v)
			} else if !reflect.DeepEqual(v, tc.expected) {
				t.Fatalf("%s\n\tForth(%#v) expected %v, got %v",
					description, tc.input, tc.expected, v)
			}
		} else if tc.expected != nil {
			t.Fatalf("%s\n\tForth(%#v) expected %v, got an error: %q",
				description, tc.input, tc.expected, err)
		}
	}
}

func TestForth(t *testing.T) {
	for _, testSection := range testSections {
		runTestCases(t, testSection.name, testSection.tests)
	}
}

func BenchmarkForth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testSection := range testSections {
			for _, test := range testSection.tests {
				Forth(test.input)
			}
		}
	}
}
