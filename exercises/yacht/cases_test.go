package yacht

// Source: exercism/problem-specifications
// Commit: 6eb19d2 yacht: Add "Four of a kind is not a full house" test case (#1232)
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	dice        []int
	category    string
	expected    int
}{
	{
		description: "Yacht",
		dice:        []int{5, 5, 5, 5, 5},
		category:    "yacht",
		expected:    50,
	},
	{
		description: "Not Yacht",
		dice:        []int{1, 3, 3, 2, 5},
		category:    "yacht",
		expected:    0,
	},
	{
		description: "Ones",
		dice:        []int{1, 1, 1, 3, 5},
		category:    "ones",
		expected:    3,
	},
	{
		description: "Ones, out of order",
		dice:        []int{3, 1, 1, 5, 1},
		category:    "ones",
		expected:    3,
	},
	{
		description: "No ones",
		dice:        []int{4, 3, 6, 5, 5},
		category:    "ones",
		expected:    0,
	},
	{
		description: "Twos",
		dice:        []int{2, 3, 4, 5, 6},
		category:    "twos",
		expected:    2,
	},
	{
		description: "Fours",
		dice:        []int{1, 4, 1, 4, 1},
		category:    "fours",
		expected:    8,
	},
	{
		description: "Yacht counted as threes",
		dice:        []int{3, 3, 3, 3, 3},
		category:    "threes",
		expected:    15,
	},
	{
		description: "Yacht of 3s counted as fives",
		dice:        []int{3, 3, 3, 3, 3},
		category:    "fives",
		expected:    0,
	},
	{
		description: "Sixes",
		dice:        []int{2, 3, 4, 5, 6},
		category:    "sixes",
		expected:    6,
	},
	{
		description: "Full house two small, three big",
		dice:        []int{2, 2, 4, 4, 4},
		category:    "full house",
		expected:    16,
	},
	{
		description: "Full house three small, two big",
		dice:        []int{5, 3, 3, 5, 3},
		category:    "full house",
		expected:    19,
	},
	{
		description: "Two pair is not a full house",
		dice:        []int{2, 2, 4, 4, 5},
		category:    "full house",
		expected:    0,
	},
	{
		description: "Four of a kind is not a full house",
		dice:        []int{1, 4, 4, 4, 4},
		category:    "full house",
		expected:    0,
	},
	{
		description: "Yacht is not a full house",
		dice:        []int{2, 2, 2, 2, 2},
		category:    "full house",
		expected:    0,
	},
	{
		description: "Four of a Kind",
		dice:        []int{6, 6, 4, 6, 6},
		category:    "four of a kind",
		expected:    24,
	},
	{
		description: "Yacht can be scored as Four of a Kind",
		dice:        []int{3, 3, 3, 3, 3},
		category:    "four of a kind",
		expected:    12,
	},
	{
		description: "Full house is not Four of a Kind",
		dice:        []int{3, 3, 3, 5, 5},
		category:    "four of a kind",
		expected:    0,
	},
	{
		description: "Little Straight",
		dice:        []int{3, 5, 4, 1, 2},
		category:    "little straight",
		expected:    30,
	},
	{
		description: "Little Straight as Big Straight",
		dice:        []int{1, 2, 3, 4, 5},
		category:    "big straight",
		expected:    0,
	},
	{
		description: "Four in order but not a little straight",
		dice:        []int{1, 1, 2, 3, 4},
		category:    "little straight",
		expected:    0,
	},
	{
		description: "No pairs but not a little straight",
		dice:        []int{1, 2, 3, 4, 6},
		category:    "little straight",
		expected:    0,
	},
	{
		description: "Minimum is 1, maximum is 5, but not a little straight",
		dice:        []int{1, 1, 3, 4, 5},
		category:    "little straight",
		expected:    0,
	},
	{
		description: "Big Straight",
		dice:        []int{4, 6, 2, 5, 3},
		category:    "big straight",
		expected:    30,
	},
	{
		description: "Big Straight as little straight",
		dice:        []int{6, 5, 4, 3, 2},
		category:    "little straight",
		expected:    0,
	},
	{
		description: "Choice",
		dice:        []int{3, 3, 5, 6, 6},
		category:    "choice",
		expected:    23,
	},
	{
		description: "Yacht as choice",
		dice:        []int{2, 2, 2, 2, 2},
		category:    "choice",
		expected:    10,
	},
}
