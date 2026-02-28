package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Funcs(template.FuncMap{"fmtArray": FormatArray}).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"flatten": &[]testCase{},
	}
	if err := gen.Gen("flatten-array", j, t); err != nil {
		log.Fatal(err)
	}
}

func FormatArray(data []any) string {
	return strings.ReplaceAll(fmt.Sprintf("%#v", data), "interface {}", "any")
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Array any `json:"array"`
	} `json:"input"`
	Expected []any `json:"expected"`
}

// Template to generate test cases.
var tmpl = `{{.Header}}

var testCases = []struct {
	description	string
	input		any
	expected	[]any
}{ {{range .J.flatten}}
	{
		description:	{{printf "%q"  .Description}},
		input:			{{fmtArray .Input.Array}},
		expected:		{{fmtArray .Expected}},
	},{{end}}
}
`
