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
	var j = map[string]interface{}{
		"abbreviate": &[]testCase{},
	}
	if err := gen.Gen("acronym", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Phrase string `json:"phrase"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package acronym

{{.Header}}

type acronymTest struct {
	description string
	input       string
	expected    string
}

var testCases = []acronymTest {
{{range .J.abbreviate}}{
	description: {{printf "%q" .Description }},
	input: {{printf "%q" .Input.Phrase }},
	expected: {{printf "%q" .Expected }},
},
{{end}}
}
`
