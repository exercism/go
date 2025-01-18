package anagram

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 6373ab6 anagram: reimplement cases where `candidates` isn't a set (#1943)

type anagramTest struct {
	description string
	subject     string
	candidates  []string
	expected    []string
}

var testCases = []anagramTest{
	{
		description: "no matches",
		subject:     "diaper",
		candidates:  []string{"hello", "world", "zombies", "pants"},
		expected:    []string{},
	},
	{
		description: "detects two anagrams",
		subject:     "master",
		candidates:  []string{"stream", "pigeon", "maters"},
		expected:    []string{"stream", "maters"},
	},
	{
		description: "detects two anagrams",
		subject:     "solemn",
		candidates:  []string{"lemons", "cherry", "melons"},
		expected:    []string{"lemons", "melons"},
	},
	{
		description: "does not detect anagram subsets",
		subject:     "good",
		candidates:  []string{"dog", "goody"},
		expected:    []string{},
	},
	{
		description: "detects anagram",
		subject:     "listen",
		candidates:  []string{"enlists", "google", "inlets", "banana"},
		expected:    []string{"inlets"},
	},
	{
		description: "detects three anagrams",
		subject:     "allergy",
		candidates:  []string{"gallery", "ballerina", "regally", "clergy", "largely", "leading"},
		expected:    []string{"gallery", "regally", "largely"},
	},
	{
		description: "detects multiple anagrams with different case",
		subject:     "nose",
		candidates:  []string{"Eons", "ONES"},
		expected:    []string{"Eons", "ONES"},
	},
	{
		description: "does not detect non-anagrams with identical checksum",
		subject:     "mass",
		candidates:  []string{"last"},
		expected:    []string{},
	},
	{
		description: "detects anagrams case-insensitively",
		subject:     "Orchestra",
		candidates:  []string{"cashregister", "Carthorse", "radishes"},
		expected:    []string{"Carthorse"},
	},
	{
		description: "detects anagrams using case-insensitive subject",
		subject:     "Orchestra",
		candidates:  []string{"cashregister", "carthorse", "radishes"},
		expected:    []string{"carthorse"},
	},
	{
		description: "detects anagrams using case-insensitive possible matches",
		subject:     "orchestra",
		candidates:  []string{"cashregister", "Carthorse", "radishes"},
		expected:    []string{"Carthorse"},
	},
	{
		description: "does not detect an anagram if the original word is repeated",
		subject:     "go",
		candidates:  []string{"go Go GO"},
		expected:    []string{},
	},
	{
		description: "anagrams must use all letters exactly once",
		subject:     "tapper",
		candidates:  []string{"patter"},
		expected:    []string{},
	},
	{
		description: "words are not anagrams of themselves",
		subject:     "BANANA",
		candidates:  []string{"BANANA"},
		expected:    []string{},
	},
	{
		description: "words are not anagrams of themselves even if letter case is partially different",
		subject:     "BANANA",
		candidates:  []string{"Banana"},
		expected:    []string{},
	},
	{
		description: "words are not anagrams of themselves even if letter case is completely different",
		subject:     "BANANA",
		candidates:  []string{"banana"},
		expected:    []string{},
	},
	{
		description: "words other than themselves can be anagrams",
		subject:     "LISTEN",
		candidates:  []string{"LISTEN", "Silent"},
		expected:    []string{"Silent"},
	},
}
