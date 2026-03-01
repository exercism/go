package proteintranslation

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 87bd801 fix: prevent matching STOP in a valid non-STOP sequence (#2458)

var testCases = []struct {
	description   string
	input         string
	expected      []string
	expectedError error
}{
	{
		description:   "Empty RNA sequence results in no proteins",
		input:         "",
		expected:      []string(nil),
		expectedError: nil,
	},
	{
		description:   "Methionine RNA sequence",
		input:         "AUG",
		expected:      []string{"Methionine"},
		expectedError: nil,
	},
	{
		description:   "Phenylalanine RNA sequence 1",
		input:         "UUU",
		expected:      []string{"Phenylalanine"},
		expectedError: nil,
	},
	{
		description:   "Phenylalanine RNA sequence 2",
		input:         "UUC",
		expected:      []string{"Phenylalanine"},
		expectedError: nil,
	},
	{
		description:   "Leucine RNA sequence 1",
		input:         "UUA",
		expected:      []string{"Leucine"},
		expectedError: nil,
	},
	{
		description:   "Leucine RNA sequence 2",
		input:         "UUG",
		expected:      []string{"Leucine"},
		expectedError: nil,
	},
	{
		description:   "Serine RNA sequence 1",
		input:         "UCU",
		expected:      []string{"Serine"},
		expectedError: nil,
	},
	{
		description:   "Serine RNA sequence 2",
		input:         "UCC",
		expected:      []string{"Serine"},
		expectedError: nil,
	},
	{
		description:   "Serine RNA sequence 3",
		input:         "UCA",
		expected:      []string{"Serine"},
		expectedError: nil,
	},
	{
		description:   "Serine RNA sequence 4",
		input:         "UCG",
		expected:      []string{"Serine"},
		expectedError: nil,
	},
	{
		description:   "Tyrosine RNA sequence 1",
		input:         "UAU",
		expected:      []string{"Tyrosine"},
		expectedError: nil,
	},
	{
		description:   "Tyrosine RNA sequence 2",
		input:         "UAC",
		expected:      []string{"Tyrosine"},
		expectedError: nil,
	},
	{
		description:   "Cysteine RNA sequence 1",
		input:         "UGU",
		expected:      []string{"Cysteine"},
		expectedError: nil,
	},
	{
		description:   "Cysteine RNA sequence 2",
		input:         "UGC",
		expected:      []string{"Cysteine"},
		expectedError: nil,
	},
	{
		description:   "Tryptophan RNA sequence",
		input:         "UGG",
		expected:      []string{"Tryptophan"},
		expectedError: nil,
	},
	{
		description:   "STOP codon RNA sequence 1",
		input:         "UAA",
		expected:      []string(nil),
		expectedError: nil,
	},
	{
		description:   "STOP codon RNA sequence 2",
		input:         "UAG",
		expected:      []string(nil),
		expectedError: nil,
	},
	{
		description:   "STOP codon RNA sequence 3",
		input:         "UGA",
		expected:      []string(nil),
		expectedError: nil,
	},
	{
		description:   "Sequence of two protein codons translates into proteins",
		input:         "UUUUUU",
		expected:      []string{"Phenylalanine", "Phenylalanine"},
		expectedError: nil,
	},
	{
		description:   "Sequence of two different protein codons translates into proteins",
		input:         "UUAUUG",
		expected:      []string{"Leucine", "Leucine"},
		expectedError: nil,
	},
	{
		description:   "Translate RNA strand into correct protein list",
		input:         "AUGUUUUGG",
		expected:      []string{"Methionine", "Phenylalanine", "Tryptophan"},
		expectedError: nil,
	},
	{
		description:   "Translation stops if STOP codon at beginning of sequence",
		input:         "UAGUGG",
		expected:      []string(nil),
		expectedError: nil,
	},
	{
		description:   "Translation stops if STOP codon at end of two-codon sequence",
		input:         "UGGUAG",
		expected:      []string{"Tryptophan"},
		expectedError: nil,
	},
	{
		description:   "Translation stops if STOP codon at end of three-codon sequence",
		input:         "AUGUUUUAA",
		expected:      []string{"Methionine", "Phenylalanine"},
		expectedError: nil,
	},
	{
		description:   "Translation stops if STOP codon in middle of three-codon sequence",
		input:         "UGGUAGUGG",
		expected:      []string{"Tryptophan"},
		expectedError: nil,
	},
	{
		description:   "Translation stops if STOP codon in middle of six-codon sequence",
		input:         "UGGUGUUAUUAAUGGUUU",
		expected:      []string{"Tryptophan", "Cysteine", "Tyrosine"},
		expectedError: nil,
	},
	{
		description:   "Sequence of two non-STOP codons does not translate to a STOP codon",
		input:         "AUGAUG",
		expected:      []string{"Methionine", "Methionine"},
		expectedError: nil,
	},
	{
		description:   "Unknown amino acids, not part of a codon, can't translate",
		input:         "XYZ",
		expected:      nil,
		expectedError: ErrInvalidBase,
	},
	{
		description:   "Incomplete RNA sequence can't translate",
		input:         "AUGU",
		expected:      nil,
		expectedError: ErrInvalidBase,
	},
	{
		description:   "Incomplete RNA sequence can translate if valid until a STOP codon",
		input:         "UUCUUCUAAUGGU",
		expected:      []string{"Phenylalanine", "Phenylalanine"},
		expectedError: nil,
	},
}
