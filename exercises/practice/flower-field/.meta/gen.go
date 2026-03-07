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
		"annotate": &[]testCase{},
	}
	if err := gen.Gen("flower-field", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Garden []string `json:"garden"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{ {{range .J.annotate}}
	{
		description: {{printf "%q" .Description}},
		input:       []string{
			{{range .Input.Garden}} {{printf "%q" .}},
			{{end}}
		},
		expected:       []string{
			{{range .Expected}} {{printf "%q" .}},
			{{end}}
		},
	},
{{end}}
}`
