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
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nGot error: %q",
				test.input, test.expected, err)
		}
		if actual != test.expected {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nActual: %s",
				test.input, test.expected, actual)
		}
		t.Logf("PASS: Protein translation test: %s", test.input)
	}

	for _, test := range codonInvalidBaseCases {
		actual, err := FromCodon(test.input)
		if !errors.As(err, &ErrInvalidBase{}) {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected type error: %T\nActual type error: %T",
				test.input, ErrInvalidBase{}, err)
		}
		if actual != test.expected {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nActual: %s",
				test.input, test.expected, actual)
		}
		t.Logf("PASS: Protein translation test: %s", test.input)
	}

	for _, test := range codonStopCases {
		actual, err := FromCodon(test.input)
		if !errors.As(err, &ErrStop{}) {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected type error: %T\nActual type error: %T",
				test.input, ErrStop{}, err)
		}
		if actual != test.expected {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nActual: %s",
				test.input, test.expected, actual)
		}
		t.Logf("PASS: Protein translation test: %s", test.input)
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
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nGot error: %q",
				test.input, test.expected, err)
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FAIL: RNA Translation test: %s\nExpected: %q\nActual %q", test.input, test.expected, actual)
		}
		t.Logf("PASS: RNA translation test: %s", test.input)
	}
	for _, test := range proteinInvalidBaseCases {
		actual, err := FromRNA(test.input)
		if !errors.As(err, &ErrInvalidBase{}) {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected type error: %T\nActual type error: %T",
				test.input, ErrInvalidBase{}, err)
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FAIL: RNA Translation test: %s\nExpected: %q\nActual %q", test.input, test.expected, actual)
		}
		t.Logf("PASS: RNA translation test: %s", test.input)
	}
}

func BenchmarkCodon(b *testing.B) {
	for _, test := range codonSuccessCases {
		for i := 0; i < b.N; i++ {
			FromCodon(test.input)
		}
	}
	for _, test := range codonInvalidBaseCases {
		for i := 0; i < b.N; i++ {
			FromCodon(test.input)
		}
	}
	for _, test := range codonStopCases {
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
	for _, test := range proteinInvalidBaseCases {
		for i := 0; i < b.N; i++ {
			FromRNA(test.input)
		}
	}
}
