// This file contains tests from the shared problem specifications repo.
package listops

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 7a8722a Reorder keys (#1960)

type testCaseAppend struct {
	description  string
	list1, list2 []int
	expected     []int
}

var testCasesAppend = []testCaseAppend{
	{
		description: "empty lists",
		list1:       []int{},
		list2:       []int{},
		expected:    []int{},
	},
	{
		description: "list to empty list",
		list1:       []int{},
		list2:       []int{1, 2, 3, 4},
		expected:    []int{1, 2, 3, 4},
	},
	{
		description: "empty list to list",
		list1:       []int{1, 2, 3, 4},
		list2:       []int{},
		expected:    []int{1, 2, 3, 4},
	},
	{
		description: "non-empty lists",
		list1:       []int{1, 2},
		list2:       []int{2, 3, 4, 5},
		expected:    []int{1, 2, 2, 3, 4, 5},
	},
}

type testCaseConcat struct {
	description string
	lists       [][]int
	expected    []int
}

var testCasesConcat = []testCaseConcat{
	{
		description: "empty list",
		lists:       [][]int{},
		expected:    []int{},
	},
	{
		description: "list of lists",
		lists:       [][]int{[]int{1, 2}, []int{3}, []int{}, []int{4, 5, 6}},
		expected:    []int{1, 2, 3, 4, 5, 6},
	},
}

type testCaseFilter struct {
	description string
	list        []int
	function    func(int) bool
	functionStr string
	expected    []int
}

var testCasesFilter = []testCaseFilter{
	{
		description: "empty list",
		list:        []int{},
		function:    func(x int) bool { return x%2 == 1 },
		functionStr: "(x) -> x modulo 2 == 1",
		expected:    []int{},
	},
	{
		description: "non-empty list",
		list:        []int{1, 2, 3, 5},
		function:    func(x int) bool { return x%2 == 1 },
		functionStr: "(x) -> x modulo 2 == 1",
		expected:    []int{1, 3, 5},
	},
}

type testCaseFoldl struct {
	description string
	list        []int
	initial     int
	function    func(int, int) int
	functionStr string
	expected    int
}

var testCasesFoldl = []testCaseFoldl{
	{
		description: "direction dependent function applied to non-empty list",
		list:        []int{2, 5},
		initial:     5,
		function:    func(x, y int) int { return x / y },
		functionStr: "(x, y) -> x / y",
		expected:    0,
	},
	{
		description: "empty list",
		list:        []int{},
		initial:     2,
		function:    func(acc, el int) int { return el * acc },
		functionStr: "(acc, el) -> el * acc",
		expected:    2,
	},
	{
		description: "direction independent function applied to non-empty list",
		list:        []int{1, 2, 3, 4},
		initial:     5,
		function:    func(acc, el int) int { return el + acc },
		functionStr: "(acc, el) -> el + acc",
		expected:    15,
	},
}

type testCaseFoldr struct {
	description string
	list        []int
	initial     int
	function    func(int, int) int
	functionStr string
	expected    int
}

var testCasesFoldr = []testCaseFoldr{
	{
		description: "direction dependent function applied to non-empty list",
		list:        []int{2, 5},
		initial:     5,
		function:    func(x, y int) int { return x / y },
		functionStr: "(x, y) -> x / y",
		expected:    2,
	},
	{
		description: "empty list",
		list:        []int{},
		initial:     2,
		function:    func(acc, el int) int { return el * acc },
		functionStr: "(acc, el) -> el * acc",
		expected:    2,
	},
	{
		description: "direction independent function applied to non-empty list",
		list:        []int{1, 2, 3, 4},
		initial:     5,
		function:    func(acc, el int) int { return el + acc },
		functionStr: "(acc, el) -> el + acc",
		expected:    15,
	},
}

type testCaseLength struct {
	description string
	list        []int
	expected    int
}

var testCasesLength = []testCaseLength{
	{
		description: "empty list",
		list:        []int{},
		expected:    0,
	},
	{
		description: "non-empty list",
		list:        []int{1, 2, 3, 4},
		expected:    4,
	},
}

type testCaseMap struct {
	description string
	list        []int
	function    func(int) int
	functionStr string
	expected    []int
}

var testCasesMap = []testCaseMap{
	{
		description: "empty list",
		list:        []int{},
		function:    func(x int) int { return x + 1 },
		functionStr: "(x) -> x + 1",
		expected:    []int{},
	},
	{
		description: "non-empty list",
		list:        []int{1, 3, 5, 7},
		function:    func(x int) int { return x + 1 },
		functionStr: "(x) -> x + 1",
		expected:    []int{2, 4, 6, 8},
	},
}

type testCaseReverse struct {
	description string
	list        []int
	expected    []int
}

var testCasesReverse = []testCaseReverse{
	{
		description: "empty list",
		list:        []int{},
		expected:    []int{},
	},
	{
		description: "non-empty list",
		list:        []int{1, 3, 5, 7},
		expected:    []int{7, 5, 3, 1},
	},
}
