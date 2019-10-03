package allergies

// Source: exercism/problem-specifications
// Commit: ca9c09ed Allergies: Modified canonical description to make it generatable (#1494)
// Problem Specifications Version: 2.0.0

type aTest struct {
	description string
	item        string
	score       uint
	expected    bool
}

type lTest struct {
	description string
	score       uint
	expected    []string
}

func getTests() ([]aTest, []lTest) {

	var allergicToTests []aTest
	var listTests []lTest

	// testing for eggs allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "eggs",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to eggs",
			item:        "eggs",
			score:       1,
			expected:    true,
		},
		aTest{
			description: "allergic to eggs and something else",
			item:        "eggs",
			score:       3,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not eggs",
			item:        "eggs",
			score:       2,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "eggs",
			score:       255,
			expected:    true,
		},
	)

	// testing for peanuts allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "peanuts",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to peanuts",
			item:        "peanuts",
			score:       2,
			expected:    true,
		},
		aTest{
			description: "allergic to peanuts and something else",
			item:        "peanuts",
			score:       7,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not peanuts",
			item:        "peanuts",
			score:       5,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "peanuts",
			score:       255,
			expected:    true,
		},
	)

	// testing for shellfish allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "shellfish",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to shellfish",
			item:        "shellfish",
			score:       4,
			expected:    true,
		},
		aTest{
			description: "allergic to shellfish and something else",
			item:        "shellfish",
			score:       14,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not shellfish",
			item:        "shellfish",
			score:       10,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "shellfish",
			score:       255,
			expected:    true,
		},
	)

	// testing for strawberries allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "strawberries",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to strawberries",
			item:        "strawberries",
			score:       8,
			expected:    true,
		},
		aTest{
			description: "allergic to strawberries and something else",
			item:        "strawberries",
			score:       28,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not strawberries",
			item:        "strawberries",
			score:       20,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "strawberries",
			score:       255,
			expected:    true,
		},
	)

	// testing for tomatoes allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "tomatoes",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to tomatoes",
			item:        "tomatoes",
			score:       16,
			expected:    true,
		},
		aTest{
			description: "allergic to tomatoes and something else",
			item:        "tomatoes",
			score:       56,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not tomatoes",
			item:        "tomatoes",
			score:       40,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "tomatoes",
			score:       255,
			expected:    true,
		},
	)

	// testing for chocolate allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "chocolate",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to chocolate",
			item:        "chocolate",
			score:       32,
			expected:    true,
		},
		aTest{
			description: "allergic to chocolate and something else",
			item:        "chocolate",
			score:       112,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not chocolate",
			item:        "chocolate",
			score:       80,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "chocolate",
			score:       255,
			expected:    true,
		},
	)

	// testing for pollen allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "pollen",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to pollen",
			item:        "pollen",
			score:       64,
			expected:    true,
		},
		aTest{
			description: "allergic to pollen and something else",
			item:        "pollen",
			score:       224,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not pollen",
			item:        "pollen",
			score:       160,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "pollen",
			score:       255,
			expected:    true,
		},
	)

	// testing for cats allergy
	allergicToTests = append(allergicToTests,
		aTest{
			description: "not allergic to anything",
			item:        "cats",
			score:       0,
			expected:    false,
		},
		aTest{
			description: "allergic only to cats",
			item:        "cats",
			score:       128,
			expected:    true,
		},
		aTest{
			description: "allergic to cats and something else",
			item:        "cats",
			score:       192,
			expected:    true,
		},
		aTest{
			description: "allergic to something, but not cats",
			item:        "cats",
			score:       64,
			expected:    false,
		},
		aTest{
			description: "allergic to everything",
			item:        "cats",
			score:       255,
			expected:    true,
		},
	)

	// list when:
	listTests = append(listTests,
		lTest{"no allergies", 0, []string{}},
		lTest{"just eggs", 1, []string{"eggs"}},
		lTest{"just peanuts", 2, []string{"peanuts"}},
		lTest{"just strawberries", 8, []string{"strawberries"}},
		lTest{"eggs and peanuts", 3, []string{"eggs", "peanuts"}},
		lTest{"more than eggs but not peanuts", 5, []string{"eggs", "shellfish"}},
		lTest{"lots of stuff", 248, []string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
		lTest{"everything", 255, []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
		lTest{"no allergen score parts", 509, []string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
	)

	return allergicToTests, listTests
}
