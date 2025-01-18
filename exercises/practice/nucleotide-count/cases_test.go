package dna

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 5fc501b Remove unneeded nesting (#1798)

var testCases = []struct {
	description   string
	strand        string
	expected      Histogram
	errorExpected bool
}{
	{
		description:   "empty strand",
		strand:        "",
		errorExpected: false,
		expected:      Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0},
	},
	{
		description:   "can count one nucleotide in single-character input",
		strand:        "G",
		errorExpected: false,
		expected:      Histogram{'A': 0, 'C': 0, 'G': 1, 'T': 0},
	},
	{
		description:   "strand with repeated nucleotide",
		strand:        "GGGGGGG",
		errorExpected: false,
		expected:      Histogram{'A': 0, 'C': 0, 'G': 7, 'T': 0},
	},
	{
		description:   "strand with multiple nucleotides",
		strand:        "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		errorExpected: false,
		expected:      Histogram{'A': 20, 'C': 12, 'G': 17, 'T': 21},
	},
	{
		description:   "strand with invalid nucleotides",
		strand:        "AGXXACT",
		errorExpected: true,
		expected:      Histogram{},
	},
}
