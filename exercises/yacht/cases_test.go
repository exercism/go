package yacht

/*
categories
	0 ones
	1 twos
	2 threes
	3 fours
	4 fives
	5 sixes
	6 Full House
	7 Four Of A Kind
	8 Little Straight
	9 Big Straight
	10 Choice
	11 Yacht
*/

var testCases = []struct {
	description string
	dice        []int
	category    int
	expected    int
}{
	{
		description: "Yacht",
		dice:        []int{5, 5, 5, 5, 5},
		category:    11,
		expected:    50,
	},
	{
		description: "Not Yacht",
		dice:        []int{1, 3, 3, 2, 5},
		category:    11,
		expected:    0,
	},
	{
		description: "Ones",
		dice:        []int{1, 1, 1, 3, 5},
		category:    0,
		expected:    3,
	},
	{
		description: "Ones, out of order",
		dice:        []int{3, 1, 1, 5, 1},
		category:    0,
		expected:    3,
	},
	{
		description: "No ones",
		dice:        []int{4, 3, 6, 5, 5},
		category:    0,
		expected:    0,
	},
	{
		description: "Fours",
		dice:        []int{1, 4, 1, 4, 1},
		category:    3,
		expected:    8,
	},
	{
		description: "Yacht counted as threes",
		dice:        []int{3, 3, 3, 3, 3},
		category:    2,
		expected:    15,
	},
	{
		description: "Yacht of 3s counted as fives",
		dice:        []int{3, 3, 3, 3, 3},
		category:    4,
		expected:    0,
	},
	{
		description: "Full house",
		dice:        []int{2, 2, 4, 4, 4},
		category:    6,
		expected:    16,
	},
	{
		description: "Full house unsorted",
		dice:        []int{5, 3, 3, 5, 3},
		category:    6,
		expected:    19,
	},
	{
		description: "Not a full house",
		dice:        []int{2, 2, 4, 4, 5},
		category:    6,
		expected:    0,
	},
	{
		description: "Yacht is not a full house",
		dice:        []int{2, 2, 2, 2, 2},
		category:    6,
		expected:    0,
	},
	{
		description: "Four of a Kind",
		dice:        []int{1, 2, 2, 2, 2},
		category:    7,
		expected:    8,
	},
	{
		description: "Four of a Kind unsorted",
		dice:        []int{6, 6, 4, 6, 6},
		category:    7,
		expected:    24,
	},
	{
		description: "Yacht can be scored as Four of a Kind",
		dice:        []int{3, 3, 3, 3, 3},
		category:    7,
		expected:    12,
	},
	{
		description: "Not Four of a Kind",
		dice:        []int{3, 3, 3, 5, 5},
		category:    7,
		expected:    0,
	},
	{
		description: "Little Straight",
		dice:        []int{1, 2, 3, 4, 5},
		category:    8,
		expected:    30,
	},
	{
		description: "Little Straight unsorted",
		dice:        []int{3, 5, 4, 1, 2},
		category:    8,
		expected:    30,
	},
	{
		description: "Little Straight as Big Straight",
		dice:        []int{1, 2, 3, 4, 5},
		category:    9,
		expected:    0,
	},
	{
		description: "Not a little straight",
		dice:        []int{1, 2, 3, 4, 6},
		category:    8,
		expected:    0,
	},
	{
		description: "Big Straight",
		dice:        []int{2, 3, 4, 5, 6},
		category:    9,
		expected:    30,
	},
	{
		description: "Big Straight unsorted",
		dice:        []int{4, 6, 2, 5, 3},
		category:    9,
		expected:    30,
	},
	{
		description: "Big Straight as little straight",
		dice:        []int{6, 5, 4, 3, 2},
		category:    8,
		expected:    0,
	},
	{
		description: "Choice",
		dice:        []int{3, 3, 5, 6, 6},
		category:    10,
		expected:    23,
	},
	{
		description: "Yacht as choice",
		dice:        []int{2, 2, 2, 2, 2},
		category:    10,
		expected:    10,
	},
	{
		description: "Big Straight as Choice",
		dice:        []int{2, 3, 4, 5, 6},
		category:    10,
		expected:    20,
	},
	{
		description: "Full House as choice",
		dice:        []int{1, 1, 1, 6, 6},
		category:    10,
		expected:    15,
	},
}
