package acronym

import (
	"testing"
)

const targetTestVersion = 2

type testCase struct {
	input    string
	expected string
}

var stringTestCases = []testCase{
	{"Portable Network Graphics", "PNG"},
	{"HyperText Markup Language", "HTML"},
	{"Ruby on Rails", "ROR"},
	{"PHP: Hypertext Preprocessor", "PHP"},
	{"First In, First Out", "FIFO"},
	{"Complementary metal-oxide semiconductor", "CMOS"},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestAcronym(t *testing.T) {
	for _, test := range stringTestCases {
		actual := Abbreviate(test.input)
		if actual != test.expected {
			t.Errorf("Acronym test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}
