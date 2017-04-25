package strand

// Source: exercism/x-common
// Commit: 0b20fff rna-transcription: Fix canonical-data.json formatting
// x-common version: 1.0.0

var rnaTests = []struct {
	input    string
	expected string
}{
	// rna complement of cytosine is guanine
	{"C", "G"},

	// rna complement of guanine is cytosine
	{"G", "C"},

	// rna complement of thymine is adenine
	{"T", "A"},

	// rna complement of adenine is uracil
	{"A", "U"},

	// rna complement
	{"ACGTGGTCTTAA", "UGCACCAGAAUU"},
}
