package allergies

// Source: exercism/problem-specifications
// Commit: 85dcb59 allergies: Update json for new "input" policy (#1033)
// Problem Specifications Version: 1.1.0

// allergicTo
type allergicResult struct {
	substance string
	result    bool
}

var allergicToTests = []struct {
	description string
	score       uint
	expected    []allergicResult
}{
	{
		description: "no allergies means not allergic",
		score:       0,
		expected: []allergicResult{
			{"peanuts", false},
			{"cats", false},
			{"strawberries", false},
		},
	},
	{
		description: "is allergic to eggs",
		score:       1,
		expected: []allergicResult{
			{"eggs", true},
		},
	},
	{
		description: "allergic to eggs in addition to other stuff",
		score:       5,
		expected: []allergicResult{
			{"eggs", true},
			{"shellfish", true},
			{"strawberries", false},
		},
	},
}

// list
var listTests = []struct {
	description string
	score       uint
	expected    []string
}{
	{"no allergies at all", 0, []string{}},
	{"allergic to just eggs", 1, []string{"eggs"}},
	{"allergic to just peanuts", 2, []string{"peanuts"}},
	{"allergic to just strawberries", 8, []string{"strawberries"}},
	{"allergic to eggs and peanuts", 3, []string{"eggs", "peanuts"}},
	{"allergic to more than eggs but not peanuts", 5, []string{"eggs", "shellfish"}},
	{"allergic to lots of stuff", 248, []string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
	{"allergic to everything", 255, []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
	{"ignore non allergen score parts", 509, []string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
}
