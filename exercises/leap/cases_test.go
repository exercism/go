package leap

// Source: exercism/problem-specifications
// Commit: e348053 leap: Apply new "input" policy
// Problem Specifications Version: 1.3.0

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{2015, false, "year not divisible by 4: common year"},
	{1996, true, "year divisible by 4, not divisible by 100: leap year"},
	{2100, false, "year divisible by 100, not divisible by 400: common year"},
	{2000, true, "year divisible by 400: leap year"},
}
