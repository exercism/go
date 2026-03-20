package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

const (
	Unknown = "Unknown"
	Int     = "Int"
	String  = "String"
	Slice   = "Slice"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"keep": &[]testCase{},
		"discard": &[]testCase{},
	}
	if err := gen.Gen("strain", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		List []any `json:"list"`
		Predicate string `json:"predicate"`
	} `json:"input"`
	Expected []any `json:"expected"`
}

func (tc testCase) Type() string {
	if len(tc.Input.List) == 0 {
		return Int
	}
	switch tc.Input.List[0].(type) {
	case float64:
		return Int
	case string:
		return String
	case []any:
		return Slice
	}
	return Unknown
}

func (tc testCase) Filter() string {
	f := strings.TrimPrefix(tc.Input.Predicate, "fn(x) ->")
	f = strings.ReplaceAll(f, "'", "\"")
	f = strings.ReplaceAll(f, "contains", "slices.Contains")
	f = strings.ReplaceAll(f, "starts_with", "strings.HasPrefix")
	switch tc.Type() {
	case Int:
		return fmt.Sprintf("func(x int) bool { return %s }", f)
	case String:
		return fmt.Sprintf("func(x string) bool { return %s }", f)
	case Slice:
		return fmt.Sprintf("func(x []int) bool { return %s }", f)
	}
	panic("Unknown")
}

func (tc testCase) fmtList(list []any) string {
	switch tc.Type() {
	case Int:
		var data []int
		for _, i := range list {
			data = append(data, int(i.(float64)))
		}
		return fmt.Sprintf("%#v", data)
	case String:
		var data []string
		for _, i := range list {
			data = append(data, i.(string))
		}
		return fmt.Sprintf("%#v", data)
	case Slice:
		var data [][]int
		for _, i := range list {
			i := i.([]any)
			var inner []int
			for _, j := range i {
				inner = append(inner, int(j.(float64)))
			}
			data = append(data, inner)
		}
		return fmt.Sprintf("%#v", data)
	}
	panic("Unknown")
}

func (tc testCase) FmtInput() string {
	return tc.fmtList(tc.Input.List)
}

func (tc testCase) FmtExpected() string {
	return tc.fmtList(tc.Expected)
}

var tmpl = `{{.Header}}

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

var intTestCases = []testCase[int] { {{range .J.keep}} {{if eq .Type "Int"}}
	{
		description: {{printf "%q" .Description}},
		keep:        true,
		input:       {{.FmtInput}},
		filter:      {{.Filter}},
		filterStr:   {{printf "%q" .Filter}},
		expected:    {{.FmtExpected}},
	},{{end}}{{end}}{{range .J.discard}}{{if eq .Type "Int"}}
	{
		description: {{printf "%q" .Description}},
		keep:        false,
		input:       {{.FmtInput}},
		filter:      {{.Filter}},
		filterStr:   {{printf "%q" .Filter}},
		expected:    {{.FmtExpected}},
	},{{end}}{{end}}
}

var stringTestCases = []testCase[string] { {{range .J.keep}} {{if eq .Type "String"}}
	{
		description: {{printf "%q" .Description}},
		keep:        true,
		input:       {{.FmtInput}},
		filter:      {{.Filter}},
		filterStr:   {{printf "%q" .Filter}},
		expected:    {{.FmtExpected}},
	},{{end}}{{end}}{{range .J.discard}}{{if eq .Type "String"}}
	{
		description: {{printf "%q" .Description}},
		keep:        false,
		input:       {{.FmtInput}},
		filter:      {{.Filter}},
		filterStr:   {{printf "%q" .Filter}},
		expected:    {{.FmtExpected}},
	},{{end}}{{end}}
}

var sliceTestCases = []testCase[[]int] { {{range .J.keep}} {{if eq .Type "Slice"}}
	{
		description: {{printf "%q" .Description}},
		keep:        true,
		input:       {{.FmtInput}},
		filter:      {{.Filter}},
		filterStr:   {{printf "%q" .Filter}},
		expected:    {{.FmtExpected}},
	},{{end}}{{end}}{{range .J.discard}}{{if eq .Type "Slice"}}
	{
		description: {{printf "%q" .Description}},
		keep:        false,
		input:       {{.FmtInput}},
		filter:      {{.Filter}},
		filterStr:   {{printf "%q" .Filter}},
		expected:    {{.FmtExpected}},
	},{{end}}{{end}}
}
`
