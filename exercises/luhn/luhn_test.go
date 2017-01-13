package luhn

import "testing"

const targetTestVersion = 1

var testCases = []struct {
	input       string
	description string
	ok          bool
}{
	{"1", "single digit strings can not be valid", false},
	{"0", "a single zero is invalid", false},
	{"046 454 286", "valid Canadian SIN", true},
	{"046 454 287", "invalid Canadian SIN", false},
	{"8273 1232 7352 0569", "invalid credit card", false},
	{"827a 1232 7352 0569", "strings that contain non-digits are not valid", false},
}

func TestValid(t *testing.T) {
	for _, test := range testCases {
		if ok := Valid(test.input); ok != test.ok {
			t.Fatalf("Valid(%s): %s\n\t Expected: %t\n\t Got: %t", test.input, test.description, ok, test.ok)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("2323 2005 7766 3554")
	}
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}
