package strand

// Source: exercism/problem-specifications
// Commit: b4c24e6 rna-transcription: Apply new "input" policy
// Problem Specifications Version: 1.2.0

var rnaTests = []struct {
	description string
	input       string
	expected    string
}{
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
