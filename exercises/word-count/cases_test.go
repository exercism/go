package wordcount

// Source: exercism/x-common
// Commit: 3b07e53 Merge pull request #117 from mikeyjcat/add-raindrops-json

var testCases = []struct {
	description string
	input       string
	output      Frequency
}{
	{
		"count one word",
		"word",
		Frequency{"word": 1},
	},
	{
		"count one of each word",
		"one of each",
		Frequency{"each": 1, "of": 1, "one": 1},
	},
	{
		"multiple occurrences of a word",
		"one fish two fish red fish blue fish",
		Frequency{"blue": 1, "fish": 4, "one": 1, "red": 1, "two": 1},
	},
	{
		"ignore punctuation",
		"car : carpet as java : javascript!!&@$%^&",
		Frequency{"as": 1, "car": 1, "carpet": 1, "java": 1, "javascript": 1},
	},
	{
		"include numbers",
		"testing, 1, 2 testing",
		Frequency{"1": 1, "2": 1, "testing": 2},
	},
	{
		"normalize case",
		"go Go GO",
		Frequency{"go": 3},
	},
}
