package strand

// Source: exercism/x-common
// Commit: 6985644 Merge pull request #121 from mikeyjcat/add-roman-numerals-test-definition

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
