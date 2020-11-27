package protein

import (
	"errors"
	"reflect"
	"testing"
)

type codonCase struct {
	input    string
	expected string
}

var codonSuccessCases = []codonCase{
	{
		"AUG",
		"Methionine",
	},
	{
		"UUU",
		"Phenylalanine",
	},
	{
		"UUC",
		"Phenylalanine",
	},
	{
		"UUA",
		"Leucine",
	},
	{
		"UUG",
		"Leucine",
	},
	{
		"UCG",
		"Serine",
	},
	{
		"UAU",
		"Tyrosine",
	},
	{
		"UAC",
		"Tyrosine",
	},
	{
		"UGU",
		"Cysteine",
	},
	{
		"UGG",
		"Tryptophan",
	},
}

var codonInvalidBaseCases = []codonCase{
	{
		"ABC",
		"",
	},
}

var codonStopCases = []codonCase{
	{
		"UAA",
		"",
	},
	{
		"UAG",
		"",
	},
	{
		"UGA",
		"",
	},
}

func TestCodon(t *testing.T) {
	for _, test := range codonSuccessCases {
		actual, err := FromCodon(test.input)
		if err != nil {
			t.Fatalf("FromCodon(%q): %v", test.input, err)
		}
		if actual != test.expected {
			t.Fatalf("FromCodon(%q) Expected: %q VS Actual: %q", test.input, test.expected, actual)
		}
	}

	for _, test := range codonInvalidBaseCases {
		_, err := FromCodon(test.input)
		if !errors.As(err, &ErrInvalidBase{}) {
			t.Fatalf("FromCodon(%q): Wrong type of error. Expected: %T VS Actual: %T", test.input, ErrInvalidBase{}, err)
		}
	}

	for _, test := range codonStopCases {
		_, err := FromCodon(test.input)
		if !errors.As(err, &ErrStop{}) {
			t.Fatalf("FromCodon(%q): Wrong type of error. Expected: %T VS Actual: %T", test.input, ErrInvalidBase{}, err)
		}
	}
}

type rnaCase struct {
	input    string
	expected []string
}

var proteinSuccessCases = []rnaCase{
	{
		"AUGUUUUCUUAAAUG",
		[]string{"Methionine", "Phenylalanine", "Serine"},
	},
	{
		"AUGUUUUGG",
		[]string{"Methionine", "Phenylalanine", "Tryptophan"},
	},
	{
		"AUGUUUUAA",
		[]string{"Methionine", "Phenylalanine"},
	},
	{
		"UGGUGUUAUUAAUGGUUU",
		[]string{"Tryptophan", "Cysteine", "Tyrosine"},
	},
}

var proteinInvalidBaseCases = []rnaCase{
	{
		"UGGAGAAUUAAUGGUUU",
		[]string{"Tryptophan"},
	},
}

func TestProtein(t *testing.T) {
	for _, test := range proteinSuccessCases {
		actual, err := FromRNA(test.input)
		if err != nil {
			t.Fatalf("FromRNA(%q): %v", test.input, err)
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FromRNA(%q) Expected: %q VS Actual: %q", test.input, test.expected, actual)
		}
	}
	for _, test := range proteinInvalidBaseCases {
		_, err := FromRNA(test.input)
		if !errors.As(err, &ErrInvalidBase{}) {
			t.Fatalf("FromRNA(%q): Wrong type of error. Expected: %T VS Actual: %T", test.input, ErrInvalidBase{}, err)
		}
	}
}

func BenchmarkCodon(b *testing.B) {
	for _, test := range codonSuccessCases {
		for i := 0; i < b.N; i++ {
			FromCodon(test.input)
		}
	}
}

func BenchmarkProtein(b *testing.B) {
	for _, test := range proteinSuccessCases {
		for i := 0; i < b.N; i++ {
			FromRNA(test.input)
		}
	}
}
