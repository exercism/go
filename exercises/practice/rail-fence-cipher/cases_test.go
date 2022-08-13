package railfence

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 42dd0ce Remove version (#1678)

type testCase struct {
	description string
	message     string
	rails       int
	expected    string
}

var encodeTests = []testCase{
	{
		description: "encode with two rails",
		message:     "XOXOXOXOXOXOXOXOXO",
		rails:       2,
		expected:    "XXXXXXXXXOOOOOOOOO",
	},
	{
		description: "encode with three rails",
		message:     "WEAREDISCOVEREDFLEEATONCE",
		rails:       3,
		expected:    "WECRLTEERDSOEEFEAOCAIVDEN",
	},
	{
		description: "encode with ending in the middle",
		message:     "EXERCISES",
		rails:       4,
		expected:    "ESXIEECSR",
	},
}

var decodeTests = []testCase{
	{
		description: "decode with three rails",
		message:     "TEITELHDVLSNHDTISEIIEA",
		rails:       3,
		expected:    "THEDEVILISINTHEDETAILS",
	},
	{
		description: "decode with five rails",
		message:     "EIEXMSMESAORIWSCE",
		rails:       5,
		expected:    "EXERCISMISAWESOME",
	},
	{
		description: "decode with six rails",
		message:     "133714114238148966225439541018335470986172518171757571896261",
		rails:       6,
		expected:    "112358132134558914423337761098715972584418167651094617711286",
	},
}
