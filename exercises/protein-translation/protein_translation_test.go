package protein

import (
	"reflect"
	"testing"
)

type codonCase struct {
	input         string
	expected      string
	errorExpected error
}

var codonTestCases = []codonCase{
	{
		"AUG",
		"Methionine",
		nil,
	},
	{
		"UUU",
		"Phenylalanine",
		nil,
	},
	{
		"UUC",
		"Phenylalanine",
		nil,
	},
	{
		"UUA",
		"Leucine",
		nil,
	},
	{
		"UUG",
		"Leucine",
		nil,
	},
	{
		"UCG",
		"Serine",
		nil,
	},
	{
		"UAU",
		"Tyrosine",
		nil,
	},
	{
		"UAC",
		"Tyrosine",
		nil,
	},
	{
		"UGU",
		"Cysteine",
		nil,
	},
	{
		"UGG",
		"Tryptophan",
		nil,
	},
	{
		"UAA",
		"",
		STOP,
	},
	{
		"UAG",
		"",
		STOP,
	},
	{
		"UGA",
		"",
		STOP,
	},
	{
		"ABC",
		"",
		ErrInvalidBase,
	},
}

func TestCodon(t *testing.T) {
	for _, test := range codonTestCases {
		actual, err := FromCodon(test.input)
		if test.errorExpected != nil {
			if test.errorExpected != err {
				t.Errorf("FAIL: Protein translation test: %s\nExpected error: %q\nActual error: %q",
					test.input, test.errorExpected, err)
			}
		} else if err != nil {
			t.Errorf("FAIL: Protein translation test: %s\nExpected: %s\nGot error: %q",
				test.input, test.expected, err)
		}
		if actual != test.expected {
			t.Errorf("FAIL: Protein translation test: %s\nExpected: %s\nActual: %s",
				test.input, test.expected, actual)
		}
		t.Logf("PASS: %s", test.input)
	}
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

func TestProtein(t *testing.T) {
	for _, test := range proteinTestCases {
		actual := FromRNA(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("FAIL: RNA Translation test: %s\nExpected: %q\nActual %q", test.input, test.expected, actual)
		}
	}
}
