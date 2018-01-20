package anagram

// Source: exercism/problem-specifications
// Commit: b8359f9 Anagram canonical-data: Update version
// Problem Specifications Version: 1.2.0

var testCases = []struct {
	description string
	subject     string
	candidates  []string
	expected    []string
}{
	{
		description: "no matches",
		subject:     "diaper",
		candidates: []string{
			"hello",
			"world",
			"zombies",
			"pants"},
		expected: []string{},
	},
	{
		description: "detects two anagrams",
		subject:     "master",
		candidates: []string{
			"stream",
			"pigeon",
			"maters"},
		expected: []string{
			"stream",
			"maters"},
	},
	{
		description: "does not detect anagram subsets",
		subject:     "good",
		candidates: []string{
			"dog",
			"goody"},
		expected: []string{},
	},
	{
		description: "detects anagram",
		subject:     "listen",
		candidates: []string{
			"enlists",
			"google",
			"inlets",
			"banana"},
		expected: []string{
			"inlets"},
	},
	{
		description: "detects three anagrams",
		subject:     "allergy",
		candidates: []string{
			"gallery",
			"ballerina",
			"regally",
			"clergy",
			"largely",
			"leading"},
		expected: []string{
			"gallery",
			"regally",
			"largely"},
	},
	{
		description: "does not detect non-anagrams with identical checksum",
		subject:     "mass",
		candidates: []string{
			"last"},
		expected: []string{},
	},
	{
		description: "detects anagrams case-insensitively",
		subject:     "Orchestra",
		candidates: []string{
			"cashregister",
			"Carthorse",
			"radishes"},
		expected: []string{
			"Carthorse"},
	},
	{
		description: "detects anagrams using case-insensitive subject",
		subject:     "Orchestra",
		candidates: []string{
			"cashregister",
			"carthorse",
			"radishes"},
		expected: []string{
			"carthorse"},
	},
	{
		description: "detects anagrams using case-insensitive possible matches",
		subject:     "orchestra",
		candidates: []string{
			"cashregister",
			"Carthorse",
			"radishes"},
		expected: []string{
			"Carthorse"},
	},
	{
		description: "does not detect a anagram if the original word is repeated",
		subject:     "go",
		candidates: []string{
			"go Go GO"},
		expected: []string{},
	},
	{
		description: "anagrams must use all letters exactly once",
		subject:     "tapper",
		candidates: []string{
			"patter"},
		expected: []string{},
	},
	{
		description: "capital word is not own anagram",
		subject:     "BANANA",
		candidates: []string{
			"Banana"},
		expected: []string{},
	},
}
