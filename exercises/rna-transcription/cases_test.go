package strand

// Source: exercism/problem-specifications
// Commit: cb1fd3a rna-transcription: rephrase negative test descriptions
// Problem Specifications Version: 1.0.1

var rnaTests = []struct {
	input    string
	expected string
}{
	// RNA complement of cytosine is guanine
	{"C", "G"},

	// RNA complement of guanine is cytosine
	{"G", "C"},

	// RNA complement of thymine is adenine
	{"T", "A"},

	// RNA complement of adenine is uracil
	{"A", "U"},

	// RNA complement
	{"ACGTGGTCTTAA", "UGCACCAGAAUU"},
}
