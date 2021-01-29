package leap

// Source: exercism/problem-specifications
// Commit: 18875ec Leap: Improve the specification so that code generation is more readable - … (#1468)
// Problem Specifications Version: 1.5.1

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{2015, false, "year not divisible by 4 in common year"},
	{1970, false, "year divisible by 2, not divisible by 4 in common year"},
	{1996, true, "year divisible by 4, not divisible by 100 in leap year"},
	{2100, false, "year divisible by 100, not divisible by 400 in common year"},
	{2000, true, "year divisible by 400 in leap year"},
	{1800, false, "year divisible by 200, not divisible by 400 in common year"},
}
