package protein

import (
	"testing"
)

const targetTestVersion = 1

type testCase struct {
	input    string
	expected string
}

var codonTestCases = []testCase{
	{"AUG", "Methionine"},
	{"UUU", "Phenylalanine"},
	{"UUC", "Phenylalanine"},
	{"UUA", "Leucine"},
	{"UUG", "Leucine"},
	{"UCU", "Serine"},
	{"UCC", "Serine"},
	{"UCA", "Serine"},
	{"UCG", "Serine"},
	{"UAU", "Tyrosine"},
	{"UAC", "Tyrosine"},
	{"UGU", "Cysteine"},
	{"UGC", "Cysteine"},
	{"UGG", "Tryptophan"},
	{"UAA", "STOP"},
	{"UAG", "STOP"},
	{"UGA", "STOP"},
}

var proteinTestCases = []testCase{
	{"AUGUUUUCUUAAAUG", "Methionine Phenylalanine Serine"},
	{"AUGUUUUGG", "Methionine Phenylalanine Tryptophan"},
	{"AUGUUUUAA", "Methionine Phenylalanine"},
	{"UGGUGUUAUUAAUGGUUU", "Tryptophan Cysteine Tyrosine"},
}

func TestCodon(t *testing.T) {
	for _, test := range codonTestCases {
		actual := FromCodon(test.input)
		if actual != test.expected {
			t.Errorf("Protein Translation test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func TestProtein(t *testing.T) {
	for _, test := range proteinTestCases {
		actual := FromRNA(test.input)
		if actual != test.expected {
			t.Errorf("Protein Translation test [%s], expected [%s], actual [%s]", test.input, test.expected, actual)
		}
	}
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}
