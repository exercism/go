package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("proverb", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []struct {
			Description string
			Property    string
			Input       struct {
				Phrase string
			}
			Expected string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package proverb

{{.Header}}

type proverbTest struct {
	input    string
	expected string
}

var stringTestCases = []proverbTest {
{{range .J.Cases}} {{range .Cases}}{
	input: {{printf "%q" .Input.Phrase }},
	expected: {{printf "%q" .Expected }},
},
{{end}}{{end}}
}
`
