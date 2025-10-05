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
		"isPaired": &[]testCase{},
	}
	if err := gen.Gen("matching-brackets", j, t); err != nil {
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

// template applied to above data structure generates the Go test cases
var tmpl = `package brackets

{{.Header}}

type bracketTest struct {
	description string
	input       string
	expected    bool
}

var testCases = []bracketTest {
{{range .J.isPaired}}{
	description:  	{{printf "%q" .Description}},
	input: 			{{printf "%q" .Input.Value}},
	expected: 		{{.Expected}},
},
{{end}}}
`
