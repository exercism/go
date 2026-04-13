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
		"twoFer": &[]testCase{},
	}
	if err := gen.Gen("two-fer", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Name string `json:"name"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase { {{range .J.twoFer}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .Input.Name}},
		expected:    {{printf "%q" .Expected}},
	},{{end}}
}`
