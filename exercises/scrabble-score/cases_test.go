package scrabble

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleTest struct {
	input    string
	expected int
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1},                           // lowercase letter
	{"A", 1},                           // uppercase letter
	{"f", 4},                           // valuable letter
	{"at", 2},                          // short word
	{"zoo", 12},                        // short, valuable word
	{"street", 6},                      // medium word
	{"quirky", 22},                     // medium, valuable word
	{"OxyphenButazone", 41},            // long, mixed-case word
	{"pinata", 8},                      // english-like word
	{"", 0},                            // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available
}
