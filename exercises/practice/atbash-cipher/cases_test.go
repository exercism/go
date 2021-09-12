package atbash

// Source: exercism/problem-specifications
// Commit: dda678b atbash-cipher: Apply new "input" policy
// Problem Specifications Version: 1.1.0

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
