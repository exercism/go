package etl

import "testing"

var transformTests = []struct {
	description string
	input       map[int][]string
	expect      map[string]int
}{
	{
		description: "single letter for one score",
		input:       map[int][]string{1: {"A"}},
		expect:      map[string]int{"a": 1},
	},
	{
		description: "multiple letters for one score",
		input:       map[int][]string{1: {"A", "E", "I", "O", "U"}},
		expect:      map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1},
	},
	{
		description: "multiple letters for multiple scores",
		input: map[int][]string{
			1: {"A", "E"},
			2: {"D", "G"},
		},
		expect: map[string]int{
			"a": 1,
			"e": 1,
			"d": 2,
			"g": 2,
		},
	},
	{
		description: "all letters",
		input: map[int][]string{
			1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
			2:  {"D", "G"},
			3:  {"B", "C", "M", "P"},
			4:  {"F", "H", "V", "W", "Y"},
			5:  {"K"},
			8:  {"J", "X"},
			10: {"Q", "Z"},
		},
		expect: map[string]int{
			"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
			"d": 2, "g": 2,
			"b": 3, "c": 3, "m": 3, "p": 3,
			"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
			"k": 5,
			"j": 8, "x": 8,
			"q": 10, "z": 10,
		},
	},
}

func equal(actual, expectation map[string]int) bool {
	if len(actual) != len(expectation) {
		return false
	}

	for k, actualVal := range actual {
		expectationVal, present := expectation[k]

		if !present || actualVal != expectationVal {
			return false
		}
	}

	return true
}

func TestTransform(t *testing.T) {
	for _, tt := range transformTests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := Transform(tt.input); !equal(actual, tt.expect) {
				t.Fatalf("Transform(%v)\n got:%v\nwant:%v", tt.input, actual, tt.expect)
			}
		})
	}
}

func BenchmarkTransform(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {

		for _, tt := range transformTests {
			Transform(tt.input)
		}

	}
}
