package sublist

// Source: exercism/problem-specifications
// Commit: 1854cd4 sublist: Apply new "input" policy
// Problem Specifications Version: 1.1.0

var testCases = []struct {
	description string
	listOne     []int
	listTwo     []int
	expected    Relation
}{
	{
		description: "empty lists",
		listOne:     []int{},
		listTwo:     []int{},
		expected:    "equal",
	},
	{
		description: "empty list within non empty list",
		listOne:     []int{},
		listTwo:     []int{1, 2, 3},
		expected:    "sublist",
	},
	{
		description: "non empty list contains empty list",
		listOne:     []int{1, 2, 3},
		listTwo:     []int{},
		expected:    "superlist",
	},
	{
		description: "list equals itself",
		listOne:     []int{1, 2, 3},
		listTwo:     []int{1, 2, 3},
		expected:    "equal",
	},
	{
		description: "different lists",
		listOne:     []int{1, 2, 3},
		listTwo:     []int{2, 3, 4},
		expected:    "unequal",
	},
	{
		description: "false start",
		listOne:     []int{1, 2, 5},
		listTwo:     []int{0, 1, 2, 3, 1, 2, 5, 6},
		expected:    "sublist",
	},
	{
		description: "consecutive",
		listOne:     []int{1, 1, 2},
		listTwo:     []int{0, 1, 1, 1, 2, 1, 2},
		expected:    "sublist",
	},
	{
		description: "sublist at start",
		listOne:     []int{0, 1, 2},
		listTwo:     []int{0, 1, 2, 3, 4, 5},
		expected:    "sublist",
	},
	{
		description: "sublist in middle",
		listOne:     []int{2, 3, 4},
		listTwo:     []int{0, 1, 2, 3, 4, 5},
		expected:    "sublist",
	},
	{
		description: "sublist at end",
		listOne:     []int{3, 4, 5},
		listTwo:     []int{0, 1, 2, 3, 4, 5},
		expected:    "sublist",
	},
	{
		description: "at start of superlist",
		listOne:     []int{0, 1, 2, 3, 4, 5},
		listTwo:     []int{0, 1, 2},
		expected:    "superlist",
	},
	{
		description: "in middle of superlist",
		listOne:     []int{0, 1, 2, 3, 4, 5},
		listTwo:     []int{2, 3},
		expected:    "superlist",
	},
	{
		description: "at end of superlist",
		listOne:     []int{0, 1, 2, 3, 4, 5},
		listTwo:     []int{3, 4, 5},
		expected:    "superlist",
	},
	{
		description: "first list missing element from second list",
		listOne:     []int{1, 3},
		listTwo:     []int{1, 2, 3},
		expected:    "unequal",
	},
	{
		description: "second list missing element from first list",
		listOne:     []int{1, 2, 3},
		listTwo:     []int{1, 3},
		expected:    "unequal",
	},
	{
		description: "order matters to a list",
		listOne:     []int{1, 2, 3},
		listTwo:     []int{3, 2, 1},
		expected:    "unequal",
	},
	{
		description: "same digits but different numbers",
		listOne:     []int{1, 0, 1},
		listTwo:     []int{10, 1},
		expected:    "unequal",
	},
}
