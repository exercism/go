package strand

// Source: exercism/problem-specifications
// Commit: 294c831 rna-transcription: add test for empty sequence (#1266)
// Problem Specifications Version: 1.3.0

var rnaTests = []struct {
	description string
	input       string
	expected    string
}{
	{"Empty RNA sequence",
		"", "",
	},
	{"RNA complement of cytosine is guanine",
		"C", "G",
	},
	{"RNA complement of guanine is cytosine",
		"G", "C",
	},
	{"RNA complement of thymine is adenine",
		"T", "A",
	},
	{"RNA complement of adenine is uracil",
		"A", "U",
	},
	{"RNA complement",
		"ACGTGGTCTTAA", "UGCACCAGAAUU",
	},
}
