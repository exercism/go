package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"slices"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
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

var testCasesAppend = []struct{
	description   string
	list1, list2  []int
	expected      []int
}{ {{range .J.append}}
	{
		description: {{printf "%q" .Description}},
		list1:       {{printf "%#v" .Input.List1}},
		list2:       {{printf "%#v" .Input.List2}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}

var testCasesConcat = []struct{
	description string
	lists       [][]int
	expected    []int
}{ {{range .J.concat}}
	{
		description: {{printf "%q" .Description}},
		lists:       {{printf "%#v" .Input.Lists}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}

var testCasesFilter = []struct{
	description string
	list        []int
	function    func(int) bool
	functionStr string
	expected    []int
}{ {{range .J.filter}}
	{
		description: {{printf "%q" .Description}},
		list:        {{printf "%#v" .Input.List}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}

var testCasesFoldl = []struct{
	description string
	list []int
	initial int
	function func(int, int) int
	functionStr string
	expected int
}{ {{range .J.foldl}}
	{
		description: {{printf "%q" .Description}},
		list:        {{printf "%#v" .Input.List}},
		initial:     {{.Input.Initial}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}

var testCasesFoldr = []struct{
	description string
	list []int
	initial int
	function func(int, int) int
	functionStr string
	expected int
}{ {{range .J.foldr}}
	{
		description: {{printf "%q" .Description}},
		list:        {{printf "%#v" .Input.List}},
		initial:     {{.Input.Initial}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}

var testCasesLength = []struct{
	description string
	list        []int
	expected    int
}{ {{range .J.length}}
	{
		description: {{printf "%q" .Description}},
		list:        {{printf "%#v" .Input.List}},
		expected:    {{.Expected}},
	},{{end}}
}

var testCasesMap = []struct{
	description string
	list []int
	function func(int) int
	functionStr string
	expected []int
}{ {{range .J.map}}
	{
		description: {{printf "%q" .Description}},
		list:        {{printf "%#v" .Input.List}},
		function:    {{.Func}},
		functionStr: {{printf "%q" .Input.Function}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}

var testCasesReverse = []struct{
	description string
	list        []int
	expected    []int
}{ {{range .J.reverse}}
	{
		description: {{printf "%q" .Description}},
		list:        {{printf "%#v" .Input.List}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}`
