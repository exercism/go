package proteintranslation

import (
	"slices"
	"testing"
)

type notNilTestCase struct {
	description string
	err         error
}

var notNilTestCases = []notNilTestCase{
	{
		description: "ErrStop",
		err:         ErrStop,
	},
	{
		description: "ErrInvalidBase",
		err:         ErrInvalidBase,
	},
}

func TestErrorsNotNil(t *testing.T) {
	for _, tc := range notNilTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.err == nil {
				t.Fatalf("FAIL: %s cannot be nil", tc.description)
			}
		})
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
					t.Fatalf("FromCodon(%q) expected error: %q, got: %q", tc.input, tc.expectedError, err)
				}
			case err != nil:
				t.Fatalf("FromCodon(%q) returned unexpected error: %q, want: %q", tc.input, err, tc.expected)
			case got != tc.expected:
				t.Fatalf("FromCodon(%q) = %q, want: %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestProtein(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := FromRNA(tc.input)
			switch {
			case tc.expectedError != nil:
				if err != tc.expectedError {
					t.Fatalf("FromRNA(%q) expected error: %v, got: %v", tc.input, tc.expectedError, err)
				}
			case err != nil:
				t.Fatalf("FromRNA(%q) returned unexpected error: %v, want: %q", tc.input, err, tc.expected)
			case !slices.Equal(got, tc.expected):
				t.Fatalf("FromRNA(%q)\n got: %v\nwant: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkCodon(b *testing.B) {
	for _, test := range codonTestCases {
		for range b.N {
			FromCodon(test.input)
		}
	}
}

func BenchmarkProtein(b *testing.B) {
	for _, test := range testCases {
		for range b.N {
			FromRNA(test.input)
		}
	}
}
