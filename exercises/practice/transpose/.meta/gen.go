package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"transpose": &[]testCase{},
	}
	if err := gen.Gen("transpose", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Lines []string `json:"lines"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package transpose

{{.Header}}

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{
{{range .J.transpose}}{
	description: 	{{printf "%q" .Description}},
	input:			{{printf "%#v" .Input.Lines}},
	expected:		{{printf "%#v" .Expected}},
},
{{end}}}
`
