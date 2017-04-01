package leap

// Source: exercism/x-common
// Commit: cc65ebe leap: Fix canonical-data.json formatting
// x-common version: 1.0.0

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{2015, false, "year not divisible by 4: common year"},
	{2016, true, "year divisible by 4, not divisible by 100: leap year"},
	{2100, false, "year divisible by 100, not divisible by 400: common year"},
	{2000, true, "year divisible by 400: leap year"},
}
