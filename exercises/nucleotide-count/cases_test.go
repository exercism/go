package dna

// Source: exercism/problem-specifications
// Commit: bbdcaee nucleotide-count: fix wrong order introduced in #951
// Problem Specifications Version: 1.2.0

// count all nucleotides in a strand
var testCases = []struct {
	description   string
	strand        string
	expected      Histogram
	errorExpected bool
}{
	{
		description: "empty strand",
		strand:      "",
		expected:    Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0},
	},
	{
		description: "can count one nucleotide in single-character input",
		strand:      "G",
		expected:    Histogram{'A': 0, 'C': 0, 'G': 1, 'T': 0},
	},
	{
		description: "strand with repeated nucleotide",
		strand:      "GGGGGGG",
		expected:    Histogram{'A': 0, 'C': 0, 'G': 7, 'T': 0},
	},
	{
		description: "strand with multiple nucleotides",
		strand:      "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		expected:    Histogram{'A': 20, 'C': 12, 'G': 17, 'T': 21},
	},
	{
		description:   "strand with invalid nucleotides",
		strand:        "AGXXACT",
		errorExpected: true,
	},
}
