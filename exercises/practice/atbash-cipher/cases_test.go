package atbash

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: d137db1 Format using prettier (#1917)

var testCases = []struct {
	description string
	phrase      string
	expected    string
}{
	{
		description: "encode yes",
		phrase:      "yes",
		expected:    "bvh",
	}, {
		description: "encode no",
		phrase:      "no",
		expected:    "ml",
	}, {
		description: "encode OMG",
		phrase:      "OMG",
		expected:    "lnt",
	}, {
		description: "encode spaces",
		phrase:      "O M G",
		expected:    "lnt",
	}, {
		description: "encode mindblowingly",
		phrase:      "mindblowingly",
		expected:    "nrmwy oldrm tob",
	}, {
		description: "encode numbers",
		phrase:      "Testing,1 2 3, testing.",
		expected:    "gvhgr mt123 gvhgr mt",
	}, {
		description: "encode deep thought",
		phrase:      "Truth is fiction.",
		expected:    "gifgs rhurx grlm",
	}, {
		description: "encode all the letters",
		phrase:      "The quick brown fox jumps over the lazy dog.",
		expected:    "gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt",
	},
}
