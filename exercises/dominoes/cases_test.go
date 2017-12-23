package dominoes

// Source: exercism/problem-specifications
// Commit: b4ceaf4 Update dominoes canonical-data expected key (#823)
// Problem Specifications Version: 2.0.0

var testCases = []struct {
	description string
	dominoes    []Dominoe
	valid       bool // true => can chain, false => cannot chain
}{
	{
		"empty input = empty output",
		[]Dominoe{},
		true,
	},
	{
		"singleton input = singleton output",
		[]Dominoe{Dominoe{1, 1}},
		true,
	},
	{
		"singleton that can't be chained",
		[]Dominoe{Dominoe{1, 2}},
		false,
	},
	{
		"three elements",
		[]Dominoe{Dominoe{1, 2}, Dominoe{3, 1}, Dominoe{2, 3}},
		true,
	},
	{
		"can reverse dominoes",
		[]Dominoe{Dominoe{1, 2}, Dominoe{1, 3}, Dominoe{2, 3}},
		true,
	},
	{
		"can't be chained",
		[]Dominoe{Dominoe{1, 2}, Dominoe{4, 1}, Dominoe{2, 3}},
		false,
	},
	{
		"disconnected - simple",
		[]Dominoe{Dominoe{1, 1}, Dominoe{2, 2}},
		false,
	},
	{
		"disconnected - double loop",
		[]Dominoe{Dominoe{1, 2}, Dominoe{2, 1}, Dominoe{3, 4}, Dominoe{4, 3}},
		false,
	},
	{
		"disconnected - single isolated",
		[]Dominoe{Dominoe{1, 2}, Dominoe{2, 3}, Dominoe{3, 1}, Dominoe{4, 4}},
		false,
	},
	{
		"need backtrack",
		[]Dominoe{Dominoe{1, 2}, Dominoe{2, 3}, Dominoe{3, 1}, Dominoe{2, 4}, Dominoe{2, 4}},
		true,
	},
	{
		"separate loops",
		[]Dominoe{Dominoe{1, 2}, Dominoe{2, 3}, Dominoe{3, 1}, Dominoe{1, 1}, Dominoe{2, 2}, Dominoe{3, 3}},
		true,
	},
	{
		"nine elements",
		[]Dominoe{Dominoe{1, 2}, Dominoe{5, 3}, Dominoe{3, 1}, Dominoe{1, 2}, Dominoe{2, 4}, Dominoe{1, 6}, Dominoe{2, 3}, Dominoe{3, 4}, Dominoe{5, 6}},
		true,
	},
}
