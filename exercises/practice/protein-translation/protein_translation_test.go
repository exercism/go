package protein

import (
	"reflect"
	"testing"
)

func TestErrorsNotNil(t *testing.T) {
	if ErrStop == nil {
		t.Fatalf("FAIL: ErrStop cannot be nil")
	}
	if ErrInvalidBase == nil {
		t.Fatalf("FAIL: ErrInvalidBase cannot be nil")
	}
}

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
		ErrStop,
	},
	{
		"UAG",
		"",
		ErrStop,
	},
	{
		"UGA",
		"",
		ErrStop,
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
				t.Fatalf("FAIL: Protein translation test: %s\nExpected error: %q\nActual error: %q",
					test.input, test.errorExpected, err)
			}
		} else if err != nil {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nGot error: %q",
				test.input, test.expected, err)
		}
		if actual != test.expected {
			t.Fatalf("FAIL: Protein translation test: %s\nExpected: %s\nActual: %s",
				test.input, test.expected, actual)
		}
		t.Logf("PASS: Protein translation test: %s", test.input)
	}
}

type rnaCase struct {
	input         string
	expected      []string
	errorExpected error
}

var proteinTestCases = []rnaCase{
	{
		"AUGUUUUCUUAAAUG",
		[]string{"Methionine", "Phenylalanine", "Serine"},
		nil,
	},
	{
		"AUGUUUUGG",
		[]string{"Methionine", "Phenylalanine", "Tryptophan"},
		nil,
	},
	{
		"AUGUUUUAA",
		[]string{"Methionine", "Phenylalanine"},
		nil,
	},
	{
		"UGGUGUUAUUAAUGGUUU",
		[]string{"Tryptophan", "Cysteine", "Tyrosine"},
		nil,
	},
	{
		"UGGAGAAUUAAUGGUUU",
		nil,
		ErrInvalidBase,
	},
}

func TestProtein(t *testing.T) {
	for _, test := range proteinTestCases {
		actual, err := FromRNA(test.input)
		if test.errorExpected != nil {
			if test.errorExpected != err {
				t.Fatalf("FAIL: RNA translation test: %s\nExpected error: %q\nActual error: %q",
					test.input, test.errorExpected, err)
			}
		} else if err != nil {
			t.Fatalf("FAIL: RNA translation test: %s\nExpected: %s\nGot error: %q",
				test.input, test.expected, err)
		}
		if !slicesEqual(actual, test.expected) {
			t.Fatalf("FAIL: RNA Translation test: %s\nExpected: %q\nActual %q", test.input, test.expected, actual)
		}
		t.Logf("PASS: RNA translation test: %s", test.input)
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func BenchmarkCodon(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, test := range codonTestCases {
		for i := 0; i < b.N; i++ {
			FromCodon(test.input)
		}
	}
}

func BenchmarkProtein(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for _, test := range proteinTestCases {
		for i := 0; i < b.N; i++ {
			FromRNA(test.input)
		}
	}
}
