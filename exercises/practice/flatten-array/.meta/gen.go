package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
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

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Array any `json:"array"`
	} `json:"input"`
	Expected []any `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package flatten

{{.Header}}

var testCases = []struct {
	description	string
	input		any
	expected	[]any
}{ {{range .J.flatten}}
{
	description:	{{printf "%q"  .Description}},
	input:			{{printf "%#v"  .Input.Array}},
	expected:		{{printf "%#v"  .Expected}},
},{{end}}
}
`
