package affinecipher

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

type testCase struct {
	description string
	inputPhrase string
	inputA      int
	inputB      int
	expectError bool
	expected    string
}

var encodeTests = []testCase{
	{
		description: "encode yes",
		inputPhrase: "yes",
		inputA:      5,
		inputB:      7,
		expectError: false,
		expected:    "xbt",
	},
	{
		description: "encode no",
		inputPhrase: "no",
		inputA:      15,
		inputB:      18,
		expectError: false,
		expected:    "fu",
	},
	{
		description: "encode OMG",
		inputPhrase: "OMG",
		inputA:      21,
		inputB:      3,
		expectError: false,
		expected:    "lvz",
	},
	{
		description: "encode O M G",
		inputPhrase: "O M G",
		inputA:      25,
		inputB:      47,
		expectError: false,
		expected:    "hjp",
	},
	{
		description: "encode mindblowingly",
		inputPhrase: "mindblowingly",
		inputA:      11,
		inputB:      15,
		expectError: false,
		expected:    "rzcwa gnxzc dgt",
	},
	{
		description: "encode numbers",
		inputPhrase: "Testing,1 2 3, testing.",
		inputA:      3,
		inputB:      4,
		expectError: false,
		expected:    "jqgjc rw123 jqgjc rw",
	},
	{
		description: "encode deep thought",
		inputPhrase: "Truth is fiction.",
		inputA:      5,
		inputB:      17,
		expectError: false,
		expected:    "iynia fdqfb ifje",
	},
	{
		description: "encode all the letters",
		inputPhrase: "The quick brown fox jumps over the lazy dog.",
		inputA:      17,
		inputB:      33,
		expectError: false,
		expected:    "swxtj npvyk lruol iejdc blaxk swxmh qzglf",
	},
	{
		description: "encode with a not coprime to m",
		inputPhrase: "This is a test.",
		inputA:      6,
		inputB:      17,
		expectError: true,
		expected:    "",
	},
}

var decodeTests = []testCase{
	{
		description: "decode exercism",
		inputPhrase: "tytgn fjr",
		inputA:      3,
		inputB:      7,
		expectError: false,
		expected:    "exercism",
	},
	{
		description: "decode a sentence",
		inputPhrase: "qdwju nqcro muwhn odqun oppmd aunwd o",
		inputA:      19,
		inputB:      16,
		expectError: false,
		expected:    "anobstacleisoftenasteppingstone",
	},
	{
		description: "decode numbers",
		inputPhrase: "odpoz ub123 odpoz ub",
		inputA:      25,
		inputB:      7,
		expectError: false,
		expected:    "testing123testing",
	},
	{
		description: "decode all the letters",
		inputPhrase: "swxtj npvyk lruol iejdc blaxk swxmh qzglf",
		inputA:      17,
		inputB:      33,
		expectError: false,
		expected:    "thequickbrownfoxjumpsoverthelazydog",
	},
	{
		description: "decode with no spaces in input",
		inputPhrase: "swxtjnpvyklruoliejdcblaxkswxmhqzglf",
		inputA:      17,
		inputB:      33,
		expectError: false,
		expected:    "thequickbrownfoxjumpsoverthelazydog",
	},
	{
		description: "decode with too many spaces",
		inputPhrase: "vszzm    cly   yd cg    qdp",
		inputA:      15,
		inputB:      16,
		expectError: false,
		expected:    "jollygreengiant",
	},
	{
		description: "decode with a not coprime to m",
		inputPhrase: "Test",
		inputA:      13,
		inputB:      5,
		expectError: true,
		expected:    "",
	},
}
