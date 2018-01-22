package railfence

// Source: exercism/problem-specifications
// Commit: c01f9ca rail-fence-cipher 1.0.1: Remove "test to" from all descriptions
// Problem Specifications Version: 1.0.1

type testCase struct {
	description string
	message     string
	rails       int
	expected    string
}

// encode
var encodeTests = []testCase{

	{"encode with two rails",
		"XOXOXOXOXOXOXOXOXO",
		2,
		"XXXXXXXXXOOOOOOOOO"},

	{"encode with three rails",
		"WEAREDISCOVEREDFLEEATONCE",
		3,
		"WECRLTEERDSOEEFEAOCAIVDEN"},

	{"encode with ending in the middle",
		"EXERCISES",
		4,
		"ESXIEECSR"},
}

// decode
var decodeTests = []testCase{

	{"decode with three rails",
		"TEITELHDVLSNHDTISEIIEA",
		3,
		"THEDEVILISINTHEDETAILS"},

	{"decode with five rails",
		"EIEXMSMESAORIWSCE",
		5,
		"EXERCISMISAWESOME"},

	{"decode with six rails",
		"133714114238148966225439541018335470986172518171757571896261",
		6,
		"112358132134558914423337761098715972584418167651094617711286"},
}
