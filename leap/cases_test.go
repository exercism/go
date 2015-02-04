package leap

// Source: exercism/x-common
// Commit: 945d08e Merge pull request #50 from soniakeys/master

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{1996, true, "leap year"},
	{1997, false, "non-leap year"},
	{1998, false, "non-leap even year"},
	{1900, false, "century"},
	{2400, true, "fourth century"},
	{2000, true, "Y2K"},
}
