package dominoes

// Source: exercism/problem-specifications
// Commit: b5bc74d dominoes: Apply new "input" policy
// Problem Specifications Version: 2.1.0

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
		//This meets the requirement of being possibly-Euclidean.
		//All vertices have even degree.
		//Nevertheless, there is no chain here, as there's no way to get from 1 to 2.
		//This test (and the two following) prevent solutions from using the even-degree test as the sole criterion,
		//as that is not a sufficient condition.
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
		//Some solutions may make a chain out of (1, 2), (2, 3), (3, 1)
		//then decide that since there are no more dominoes containing a 1,
		//there is no chain possible.
		//There is indeed a chain here, so this test checks for this line of reasoning.
		//You need to place the (2, 4) after the (1, 2) rather than the (2, 3).
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
