package resistorcolor

import "testing"

func TestColorCode(t *testing.T) {
	for _, tc := range colorCodeTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := ColorCode(tc.input)

			if actual != tc.expected {
				t.Fatalf("ColorCode(%q): expected %d, actual %d", tc.input, tc.expected, actual)
			}
		})
	}
}

func TestColors(t *testing.T) {
	for _, tc := range colorsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Colors()

			if len(actual) != len(tc.expected) {
				t.Fatalf("Colors(): expected %+v, actual %+v", tc.expected, actual)
			}

			expectedMap := makeMap(tc.expected)
			actualMap := makeMap(actual)

			if !mapsEqual(expectedMap, actualMap) {
				t.Fatalf("Colors(): expected %+v, actual %+v", tc.expected, actual)
			}
		})
	}
}

// colorCodeBench is intended to be used in BenchmarkColorCode to avoid compiler optimizations.
var colorCodeBench int

func BenchmarkColorCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range colorCodeTestCases {
			colorCodeBench = ColorCode(tc.input)
		}
	}
}

// colorsBench is intended to be used in BenchmarkColors to avoid compiler optimizations.
var colorsBench []string

func BenchmarkColors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		colorsBench = Colors()
	}
}

func makeMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}

func mapsEqual(a, b map[string]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}
