package etl

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 65e86ee Rework etl documentation (#2250)

type testCase struct {
	description string
	input       map[int][]string
	expected    map[string]int
}

var testCases = []testCase{
	{
		description: "single letter",
		input:       map[int][]string{1: []string{"A"}},
		expected:    map[string]int{"a": 1},
	},
	{
		description: "single score with multiple letters",
		input:       map[int][]string{1: []string{"A", "E", "I", "O", "U"}},
		expected:    map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1},
	},
	{
		description: "multiple scores with multiple letters",
		input:       map[int][]string{1: []string{"A", "E"}, 2: []string{"D", "G"}},
		expected:    map[string]int{"a": 1, "d": 2, "e": 1, "g": 2},
	},
	{
		description: "multiple scores with differing numbers of letters",
		input:       map[int][]string{1: []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}, 2: []string{"D", "G"}, 3: []string{"B", "C", "M", "P"}, 4: []string{"F", "H", "V", "W", "Y"}, 5: []string{"K"}, 8: []string{"J", "X"}, 10: []string{"Q", "Z"}},
		expected:    map[string]int{"a": 1, "b": 3, "c": 3, "d": 2, "e": 1, "f": 4, "g": 2, "h": 4, "i": 1, "j": 8, "k": 5, "l": 1, "m": 3, "n": 1, "o": 1, "p": 3, "q": 10, "r": 1, "s": 1, "t": 1, "u": 1, "v": 4, "w": 4, "x": 8, "y": 4, "z": 10},
	},
}
