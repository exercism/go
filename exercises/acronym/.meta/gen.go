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
	if err := gen.Gen("acronym", &j, t); err != nil {
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
			Phrase      string
			Expected    string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package acronym

{{.Header}}

type acronymTest struct {
	input    string
	expected string
}

var stringTestCases = []acronymTest {
{{range .J.Cases}} {{range .Cases}}{
	input: {{printf "%q" .Phrase }},
	expected: {{printf "%q" .Expected }},
},
{{end}}{{end}}
}
`
