package strand

// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "Empty RNA sequence",
		input:       "",
		expected:    "",
	},
	{
		description: "RNA complement of cytosine is guanine",
		input:       "C",
		expected:    "G",
	},
	{
		description: "RNA complement of guanine is cytosine",
		input:       "G",
		expected:    "C",
	},
	{
		description: "RNA complement of thymine is adenine",
		input:       "T",
		expected:    "A",
	},
	{
		description: "RNA complement of adenine is uracil",
		input:       "A",
		expected:    "U",
	},
	{
		description: "RNA complement",
		input:       "ACGTGGTCTTAA",
		expected:    "UGCACCAGAAUU",
	},
}
