package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

func sliceIntToStr(ints []int) string {
	parts := make([]string, len(ints))
	for i, v := range ints {
		parts[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("%s", strings.Join(parts, ", "))
}

func asIntList(list []int) string {
	return fmt.Sprintf("IntList{%s}", sliceIntToStr(list))
}

func asIntListSlice(list [][]int) string {
	parts := make([]string, len(list))
	for i, p := range list {
		parts[i] = fmt.Sprintf("{%s}", sliceIntToStr(p))
	}
	return fmt.Sprintf("[]IntList{%s}", strings.Join(parts, ", "))
}

func main() {
	funcMap := template.FuncMap{
		"IntList":      asIntList,
		"IntListSlice": asIntListSlice,
	}
	t, err := template.New("").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"append":  &[]testCaseAppend{},
		"concat":  &[]testCaseConcat{},
		"filter":  &[]testCaseFilter{},
		"foldl":   &[]testCaseFoldl{},
		"foldr":   &[]testCaseFoldr{},
		"length":  &[]testCaseLength{},
		"map":     &[]testCaseMap{},
		"reverse": &[]testCaseReverse{},
	}
	if err := gen.Gen("list-ops", j, t); err != nil {
		log.Fatal(err)
	}
}

type Operation struct {
	Operation string `json:"operation"`
	Value     int    `json:"value"`
	Expected  int    `json:"expected"`
}

func (op Operation) String() string {
	hasInput := []string{"push", "unshift", "delete"}
	opFunc := strings.ToUpper(op.Operation[0:1]) + op.Operation[1:]
	if slices.Contains(hasInput, op.Operation) {
		return fmt.Sprintf("{operation: %q, value: %d}", opFunc, op.Value)
	}
	return fmt.Sprintf("{operation: %q, expected: %d}", opFunc, op.Expected)
}

type testCaseAppend struct {
	Description string `json:"description"`
	Input       struct {
		List1 []int `json:"list1"`
		List2 []int `json:"list2"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

type testCaseConcat struct {
	Description string `json:"description"`
	Input       struct {
		Lists [][]int `json:"lists"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

type testCaseFilter struct {
	Description string `json:"description"`
	Input       struct {
		List     []int  `json:"list"`
		Function string `json:"function"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

type testCaseFoldl struct {
	Description string `json:"description"`
	Input       struct {
		List     []int  `json:"list"`
		Initial  int    `json:"initial"`
		Function string `json:"function"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type testCaseFoldr struct {
	Description string `json:"description"`
	Input       struct {
		List     []int  `json:"list"`
		Initial  int    `json:"initial"`
		Function string `json:"function"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type testCaseLength struct {
	Description string `json:"description"`
	Input       struct {
		List []int `json:"list"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type testCaseMap struct {
	Description string `json:"description"`
	Input       struct {
		List     []int  `json:"list"`
		Function string `json:"function"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

type testCaseReverse struct {
	Description string `json:"description"`
	Input       struct {
		List []int `json:"list"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

func (tc testCaseFilter) Func() string {
	before, after, _ := strings.Cut(tc.Input.Function[1:], ") -> ")
	after = strings.ReplaceAll(after, "modulo", "%")
	return fmt.Sprintf("func(%s int) bool {return %s}", before, after)
}

func (tc testCaseMap) Func() string {
	before, after, _ := strings.Cut(tc.Input.Function[1:], ") -> ")
	after = strings.ReplaceAll(after, "modulo", "%")
	return fmt.Sprintf("func(%s int) int {return %s}", before, after)
}

func FoldFunc(function string) string {
	before, after, _ := strings.Cut(function[1:], ") -> ")
	return fmt.Sprintf("func(%s int) int {return %s}", before, after)
}

func (tc testCaseFoldl) Func() string {
	return FoldFunc(tc.Input.Function)
}

func (tc testCaseFoldr) Func() string {
	return FoldFunc(tc.Input.Function)
}

var tmpl = `// This file contains tests from the shared problem specifications repo.
{{.Header}}

type testCaseAppend struct {
	description   string
	list1, list2  IntList
	expected      IntList
}

var testCasesAppend = []testCaseAppend { {{range .J.append}}
	{
		description: {{printf "%q" .Description}},
		list1:       {{IntList .Input.List1}},
		list2:       {{IntList .Input.List2}},
		expected:    {{IntList .Expected}},
	},{{end}}
}

type testCaseConcat struct {
	description string
	lists       []IntList
	expected    IntList
}

var testCasesConcat = []testCaseConcat { {{range .J.concat}}
	{
		description: {{printf "%q" .Description}},
		lists:       {{IntListSlice .Input.Lists}},
		expected:    {{IntList .Expected}},
	},{{end}}
}

type testCaseFilter struct {
	description string
	list        IntList
	function    func(int) bool
	functionStr string
	expected    IntList
}

var testCasesFilter = []testCaseFilter { {{range .J.filter}}
	{
		description: {{printf "%q" .Description}},
		list:        {{IntList .Input.List}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{IntList .Expected}},
	},{{end}}
}

type testCaseFoldl struct {
	description string
	list        IntList
	initial     int
	function    func(int, int) int
	functionStr string
	expected    int
}

var testCasesFoldl = []testCaseFoldl { {{range .J.foldl}}
	{
		description: {{printf "%q" .Description}},
		list:        {{IntList .Input.List}},
		initial:     {{.Input.Initial}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{.Expected}},
	},{{end}}
}

type testCaseFoldr struct {
	description string
	list        IntList
	initial     int
	function    func(int, int) int
	functionStr string
	expected    int
}

var testCasesFoldr = []testCaseFoldr { {{range .J.foldr}}
	{
		description: {{printf "%q" .Description}},
		list:        {{IntList .Input.List}},
		initial:     {{.Input.Initial}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{.Expected}},
	},{{end}}
}

type testCaseLength struct {
	description string
	list        IntList
	expected    int
}

var testCasesLength = []testCaseLength { {{range .J.length}}
	{
		description: {{printf "%q" .Description}},
		list:        {{IntList .Input.List}},
		expected:    {{.Expected}},
	},{{end}}
}

type testCaseMap struct {
	description string
	list        IntList
	function    func(int) int
	functionStr string
	expected    IntList
}

var testCasesMap = []testCaseMap { {{range .J.map}}
	{
		description: {{printf "%q" .Description}},
		list:        {{IntList .Input.List}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{IntList .Expected}},
	},{{end}}
}

type testCaseReverse struct {
	description string
	list        IntList
	expected    IntList
}

var testCasesReverse = []testCaseReverse { {{range .J.reverse}}
	{
		description: {{printf "%q" .Description}},
		list:        {{IntList .Input.List}},
		expected:    {{IntList .Expected}},
	},{{end}}
}`
