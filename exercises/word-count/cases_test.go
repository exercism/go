package wordcount

// Source: exercism/problem-specifications
// Commit: 77623ec word-count: Use camel-case for property name word-count: remove newline from eof
// Problem Specifications Version: 1.2.0

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
		"handles cramped lists",
		"one,two,three",
		Frequency{"one": 1, "three": 1, "two": 1},
	},
	{
		"handles expanded lists",
		"one,\ntwo,\nthree",
		Frequency{"one": 1, "three": 1, "two": 1},
	},
	{
		"ignore punctuation",
		"car: carpet as java: javascript!!&@$%^&",
		Frequency{"as": 1, "car": 1, "carpet": 1, "java": 1, "javascript": 1},
	},
	{
		"include numbers",
		"testing, 1, 2 testing",
		Frequency{"1": 1, "2": 1, "testing": 2},
	},
	{
		"normalize case",
		"go Go GO Stop stop",
		Frequency{"go": 3, "stop": 2},
	},
	{
		"with apostrophes",
		"First: don't laugh. Then: don't cry.",
		Frequency{"cry": 1, "don't": 2, "first": 1, "laugh": 1, "then": 1},
	},
	{
		"with quotations",
		"Joe can't tell between 'large' and large.",
		Frequency{"and": 1, "between": 1, "can't": 1, "joe": 1, "large": 2, "tell": 1},
	},
	{
		"multiple spaces not detected as a word",
		" multiple   whitespaces",
		Frequency{"multiple": 1, "whitespaces": 1},
	},
}
