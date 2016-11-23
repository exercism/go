package protein

import (
	"reflect"
	"testing"
)

const targetTestVersion = 1

type codonCase struct {
	input    string
	expected string
}

var codonTestCases = []codonCase{
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

type rnaCase struct {
	input    string
	expected []string
}

var proteinTestCases = []rnaCase{
	{"AUGUUUUCUUAAAUG", []string{"Methionine", "Phenylalanine", "Serine"}},
	{"AUGUUUUGG", []string{"Methionine", "Phenylalanine", "Tryptophan"}},
	{"AUGUUUUAA", []string{"Methionine", "Phenylalanine"}},
	{"UGGUGUUAUUAAUGGUUU", []string{"Tryptophan", "Cysteine", "Tyrosine"}},
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
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Protein Translation test [%s], expected %q, actual %q", test.input, test.expected, actual)
		}
	}
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}
