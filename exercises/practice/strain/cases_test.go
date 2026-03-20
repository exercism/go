package strain

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 40815e2 Strain: Add canonical data (#2214)

import (
	"slices"
	"strings"
)

type testCase[T any] struct {
	description string
	keep        bool
	input       []T
	filter      func(T) bool
	filterStr   string
	expected    []T
}

var intTestCases = []testCase[int]{
	{
		description: "keep on empty list returns empty list",
		keep:        true,
		input:       []int(nil),
		filter:      func(x int) bool { return true },
		filterStr:   "func(x int) bool { return  true }",
		expected:    []int(nil),
	},
	{
		description: "keeps everything",
		keep:        true,
		input:       []int{1, 3, 5},
		filter:      func(x int) bool { return true },
		filterStr:   "func(x int) bool { return  true }",
		expected:    []int{1, 3, 5},
	},
	{
		description: "keeps nothing",
		keep:        true,
		input:       []int{1, 3, 5},
		filter:      func(x int) bool { return false },
		filterStr:   "func(x int) bool { return  false }",
		expected:    []int(nil),
	},
	{
		description: "keeps first and last",
		keep:        true,
		input:       []int{1, 2, 3},
		filter:      func(x int) bool { return x%2 == 1 },
		filterStr:   "func(x int) bool { return  x % 2 == 1 }",
		expected:    []int{1, 3},
	},
	{
		description: "keeps neither first nor last",
		keep:        true,
		input:       []int{1, 2, 3},
		filter:      func(x int) bool { return x%2 == 0 },
		filterStr:   "func(x int) bool { return  x % 2 == 0 }",
		expected:    []int{2},
	},
	{
		description: "discard on empty list returns empty list",
		keep:        false,
		input:       []int(nil),
		filter:      func(x int) bool { return true },
		filterStr:   "func(x int) bool { return  true }",
		expected:    []int(nil),
	},
	{
		description: "discards everything",
		keep:        false,
		input:       []int{1, 3, 5},
		filter:      func(x int) bool { return true },
		filterStr:   "func(x int) bool { return  true }",
		expected:    []int(nil),
	},
	{
		description: "discards nothing",
		keep:        false,
		input:       []int{1, 3, 5},
		filter:      func(x int) bool { return false },
		filterStr:   "func(x int) bool { return  false }",
		expected:    []int{1, 3, 5},
	},
	{
		description: "discards first and last",
		keep:        false,
		input:       []int{1, 2, 3},
		filter:      func(x int) bool { return x%2 == 1 },
		filterStr:   "func(x int) bool { return  x % 2 == 1 }",
		expected:    []int{2},
	},
	{
		description: "discards neither first nor last",
		keep:        false,
		input:       []int{1, 2, 3},
		filter:      func(x int) bool { return x%2 == 0 },
		filterStr:   "func(x int) bool { return  x % 2 == 0 }",
		expected:    []int{1, 3},
	},
}

var stringTestCases = []testCase[string]{
	{
		description: "keeps strings",
		keep:        true,
		input:       []string{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"},
		filter:      func(x string) bool { return strings.HasPrefix(x, "z") },
		filterStr:   "func(x string) bool { return  strings.HasPrefix(x, \"z\") }",
		expected:    []string{"zebra", "zombies", "zealot"},
	},
	{
		description: "discards strings",
		keep:        false,
		input:       []string{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"},
		filter:      func(x string) bool { return strings.HasPrefix(x, "z") },
		filterStr:   "func(x string) bool { return  strings.HasPrefix(x, \"z\") }",
		expected:    []string{"apple", "banana", "cherimoya"},
	},
}

var sliceTestCases = []testCase[[]int]{
	{
		description: "keeps lists",
		keep:        true,
		input:       [][]int{[]int{1, 2, 3}, []int{5, 5, 5}, []int{5, 1, 2}, []int{2, 1, 2}, []int{1, 5, 2}, []int{2, 2, 1}, []int{1, 2, 5}},
		filter:      func(x []int) bool { return slices.Contains(x, 5) },
		filterStr:   "func(x []int) bool { return  slices.Contains(x, 5) }",
		expected:    [][]int{[]int{5, 5, 5}, []int{5, 1, 2}, []int{1, 5, 2}, []int{1, 2, 5}},
	},
	{
		description: "discards lists",
		keep:        false,
		input:       [][]int{[]int{1, 2, 3}, []int{5, 5, 5}, []int{5, 1, 2}, []int{2, 1, 2}, []int{1, 5, 2}, []int{2, 2, 1}, []int{1, 2, 5}},
		filter:      func(x []int) bool { return slices.Contains(x, 5) },
		filterStr:   "func(x []int) bool { return  slices.Contains(x, 5) }",
		expected:    [][]int{[]int{1, 2, 3}, []int{2, 1, 2}, []int{2, 2, 1}},
	},
}
