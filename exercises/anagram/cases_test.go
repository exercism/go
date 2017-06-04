package anagram

// Source: exercism/x-common
// Commit: 196fc1a anagram: Rename duplicated test case description (#671)
// x-common version: 1.0.1

type testcase struct {
	subject     string
	candidates  []string
	expected    []string
	description string
}

var testCases = []testcase{

	{
		subject:     "diaper",
		candidates:  []string{"hello", "world", "zombies", "pants"},
		expected:    []string{},
		description: "no matches",
	},

	{
		subject:     "ant",
		candidates:  []string{"tan", "stand", "at"},
		expected:    []string{"tan"},
		description: "detects simple anagram",
	},

	{
		subject:     "galea",
		candidates:  []string{"eagle"},
		expected:    []string{},
		description: "does not detect false positives",
	},

	{
		subject:     "master",
		candidates:  []string{"stream", "pigeon", "maters"},
		expected:    []string{"stream", "maters"},
		description: "detects two anagrams",
	},

	{
		subject:     "good",
		candidates:  []string{"dog", "goody"},
		expected:    []string{},
		description: "does not detect anagram subsets",
	},

	{
		subject:     "listen",
		candidates:  []string{"enlists", "google", "inlets", "banana"},
		expected:    []string{"inlets"},
		description: "detects anagram",
	},

	{
		subject:     "allergy",
		candidates:  []string{"gallery", "ballerina", "regally", "clergy", "largely", "leading"},
		expected:    []string{"gallery", "regally", "largely"},
		description: "detects three anagrams",
	},

	{
		subject:     "corn",
		candidates:  []string{"corn", "dark", "Corn", "rank", "CORN", "cron", "park"},
		expected:    []string{"cron"},
		description: "does not detect identical words",
	},

	{
		subject:     "mass",
		candidates:  []string{"last"},
		expected:    []string{},
		description: "does not detect non-anagrams with identical checksum",
	},

	{
		subject:     "Orchestra",
		candidates:  []string{"cashregister", "Carthorse", "radishes"},
		expected:    []string{"Carthorse"},
		description: "detects anagrams case-insensitively",
	},

	{
		subject:     "Orchestra",
		candidates:  []string{"cashregister", "carthorse", "radishes"},
		expected:    []string{"carthorse"},
		description: "detects anagrams using case-insensitive subject",
	},

	{
		subject:     "orchestra",
		candidates:  []string{"cashregister", "Carthorse", "radishes"},
		expected:    []string{"Carthorse"},
		description: "detects anagrams using case-insensitive possible matches",
	},

	{
		subject:     "banana",
		candidates:  []string{"Banana"},
		expected:    []string{},
		description: "does not detect a word as its own anagram",
	},

	{
		subject:     "go",
		candidates:  []string{"go Go GO"},
		expected:    []string{},
		description: "does not detect a anagram if the original word is repeated",
	},

	{
		subject:     "tapper",
		candidates:  []string{"patter"},
		expected:    []string{},
		description: "anagrams must use all letters exactly once",
	},

	{
		subject:     "BANANA",
		candidates:  []string{"Banana"},
		expected:    []string{},
		description: "capital word is not own anagram",
	},
}
