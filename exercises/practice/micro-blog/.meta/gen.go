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
		"truncate": &[]testCase{},
	}
	if err := gen.Gen("micro-blog", j, t); err != nil {
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

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       string
	expected    string
}{ {{range .J.truncate}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .Input.Phrase}},
		expected:    {{printf "%q" .Expected}},
	},{{end}}
}`
