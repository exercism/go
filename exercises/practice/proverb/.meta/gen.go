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
		"recite": &[]testCase{},
	}
	if err := gen.Gen("proverb", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Strings []string `json:"strings"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package proverb

{{.Header}}

type proverbTest struct {
	description string
	input    []string
	expected []string
}

var testCases = []proverbTest {
{{range .J.recite}} {
	description: 	{{printf "%q" .Description }},
	input: 			{{printf "%#v" .Input.Strings }},
	expected: 		{{printf "%#v" .Expected }},
},
{{end}}
}
`
