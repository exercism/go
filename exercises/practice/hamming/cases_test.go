package hamming

// Source: exercism/problem-specifications
// Commit: 7a8722a Reorder keys (#1960)

var testCases = []struct {
	description string
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{
		description: "empty strands",
		s1:          "",
		s2:          "",
		want:        0,
		expectError: false,
	},
	{
		description: "single letter identical strands",
		s1:          "A",
		s2:          "A",
		want:        0,
		expectError: false,
	},
	{
		description: "single letter different strands",
		s1:          "G",
		s2:          "T",
		want:        1,
		expectError: false,
	},
	{
		description: "long identical strands",
		s1:          "GGACTGAAATCTG",
		s2:          "GGACTGAAATCTG",
		want:        0,
		expectError: false,
	},
	{
		description: "long different strands",
		s1:          "GGACGGATTCTG",
		s2:          "AGGACGGATTCT",
		want:        9,
		expectError: false,
	},
	{
		description: "disallow first strand longer",
		s1:          "AATG",
		s2:          "AAA",
		want:        0,
		expectError: true,
	},
	{
		description: "disallow second strand longer",
		s1:          "ATA",
		s2:          "AGTG",
		want:        0,
		expectError: true,
	},
	{
		description: "disallow empty first strand",
		s1:          "",
		s2:          "G",
		want:        0,
		expectError: true,
	},
	{
		description: "disallow empty second strand",
		s1:          "G",
		s2:          "",
		want:        0,
		expectError: true,
	},
}
