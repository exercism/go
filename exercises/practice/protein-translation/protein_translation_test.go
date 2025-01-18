package protein

import (
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
	expectedError error
}

var codonTestCases = []codonCase{
	{
		input:         "AUG",
		expected:      "Methionine",
		expectedError: nil,
	},
	{
		input:         "UUU",
		expected:      "Phenylalanine",
		expectedError: nil,
	},
	{
		input:         "UUC",
		expected:      "Phenylalanine",
		expectedError: nil,
	},
	{
		input:         "UUA",
		expected:      "Leucine",
		expectedError: nil,
	},
	{
		input:         "UUG",
		expected:      "Leucine",
		expectedError: nil,
	},
	{
		input:         "UCG",
		expected:      "Serine",
		expectedError: nil,
	},
	{
		input:         "UAU",
		expected:      "Tyrosine",
		expectedError: nil,
	},
	{
		input:         "UAC",
		expected:      "Tyrosine",
		expectedError: nil,
	},
	{
		input:         "UGU",
		expected:      "Cysteine",
		expectedError: nil,
	},
	{
		input:         "UGG",
		expected:      "Tryptophan",
		expectedError: nil,
	},
	{
		input:         "UAA",
		expected:      "",
		expectedError: ErrStop,
	},
	{
		input:         "UAG",
		expected:      "",
		expectedError: ErrStop,
	},
	{
		input:         "UGA",
		expected:      "",
		expectedError: ErrStop,
	},
	{
		input:         "ABC",
		expected:      "",
		expectedError: ErrInvalidBase,
	},
}

func TestCodon(t *testing.T) {
	for _, tc := range codonTestCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := FromCodon(tc.input)
			switch {
			case tc.expectedError != nil:
				if err != tc.expectedError {
					t.Fatalf("FromCodon(%q) expected error: %v, got: %v", tc.input, tc.expectedError, err)
				}
			case err != nil:
				t.Fatalf("FromCodon(%q) returned unexpected error: %v, want: %q", tc.input, err, tc.expected)
			case got != tc.expected:
				t.Fatalf("FromCodon(%q) = %q, want: %q", tc.input, got, tc.expected)
			}
		})
	}
}

type rnaCase struct {
	input         string
	expected      []string
	expectedError error
}

var rnaTestCases = []rnaCase{
	{
		input:         "AUGUUUUCUUAAAUG",
		expected:      []string{"Methionine", "Phenylalanine", "Serine"},
		expectedError: nil,
	},
	{
		input:         "AUGUUUUGG",
		expected:      []string{"Methionine", "Phenylalanine", "Tryptophan"},
		expectedError: nil,
	},
	{
		input:         "AUGUUUUAA",
		expected:      []string{"Methionine", "Phenylalanine"},
		expectedError: nil,
	},
	{
		input:         "UGGUGUUAUUAAUGGUUU",
		expected:      []string{"Tryptophan", "Cysteine", "Tyrosine"},
		expectedError: nil,
	},
	{
		input:         "UGGAGAAUUAAUGGUUU",
		expected:      nil,
		expectedError: ErrInvalidBase,
	},
}

func TestProtein(t *testing.T) {
	for _, tc := range rnaTestCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := FromRNA(tc.input)
			switch {
			case tc.expectedError != nil:
				if err != tc.expectedError {
					t.Fatalf("FromRNA(%q) expected error: %v, got: %v", tc.input, tc.expectedError, err)
				}
			case err != nil:
				t.Fatalf("FromRNA(%q) returned unexpected error: %v, want: %q", tc.input, err, tc.expected)
			case !slicesEqual(got, tc.expected):
				t.Fatalf("FromRNA(%q)\n got: %v\nwant: %v", tc.input, got, tc.expected)
			}
		})
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
	for _, test := range rnaTestCases {
		for i := 0; i < b.N; i++ {
			FromRNA(test.input)
		}
	}
}
