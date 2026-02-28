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
		"rows": &[]testCase{},
	}
	if err := gen.Gen("diamond", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Letter string `json:"letter"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       string
	expected    []string
	expectedError error
}{
{{range .J.rows}}{
	description: {{printf "%q"   .Description}},
	input: {{printf "%q"   .Input.Letter}},
	expected: []string{
		{{range .Expected}} "{{printf "%v" .}}", 
		{{end}}
	},
	expectedError: nil,
},
{{end}}
}`
