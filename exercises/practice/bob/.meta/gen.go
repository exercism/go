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
	j := map[string]interface{}{
		"response": &[]testCase{},
	}
	if err := gen.Gen("bob", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		HeyBob string `json:"heyBob"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package bob

{{.Header}}

var testCases = []struct {
	description string
	input       string
	expected    string
}{
{{range .J.response}}{
	description: {{printf "%q" .Description}},
	input: {{printf "%q" .Input.HeyBob}},
	expected: {{printf "%q" .Expected}},
},
{{end}}}
`
