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
		"valid": &[]testCase{},
	}
	if err := gen.Gen("luhn", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Value string `json:"value"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

var tmpl = `package luhn

{{.Header}}

var testCases = []struct {
	description string
	input       string
	expected	bool
}{
{{range .J.valid}}{
	description: 	{{printf "%q" .Description}},
	input: 			{{printf "%q" .Input.Value}},
	expected:    	{{.Expected}},
},
{{end}}
}`
