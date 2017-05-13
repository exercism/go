package atbash

// Source: exercism/x-common
// Commit: bb4f220 atbash-cipher: Fix canonical-data.json formatting
// x-common version: 1.0.0

type atbashTest struct {
	s        string
	expected string
}

var tests = []atbashTest{
	{
		s:        "yes",
		expected: "bvh",
	},
	{
		s:        "no",
		expected: "ml",
	},
	{
		s:        "OMG",
		expected: "lnt",
	},
	{
		s:        "O M G",
		expected: "lnt",
	},
	{
		s:        "mindblowingly",
		expected: "nrmwy oldrm tob",
	},
	{
		s:        "Testing,1 2 3, testing.",
		expected: "gvhgr mt123 gvhgr mt",
	},
	{
		s:        "Truth is fiction.",
		expected: "gifgs rhurx grlm",
	},
	{
		s:        "The quick brown fox jumps over the lazy dog.",
		expected: "gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt",
	},
	{
		s:        "exercism",
		expected: "vcvix rhn",
	},
	{
		s:        "anobstacleisoftenasteppingstone",
		expected: "zmlyh gzxov rhlug vmzhg vkkrm thglm v",
	},
	{
		s:        "testing123testing",
		expected: "gvhgr mt123 gvhgr mt",
	},
	{
		s:        "thequickbrownfoxjumpsoverthelazydog",
		expected: "gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt",
	},
}
