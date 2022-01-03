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
	var j js
	if err := gen.Gen("proverb", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Strings []string
		}
		Expected []string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package proverb

{{.Header}}

type proverbTest struct {
	description string
	input    []string
	expected []string
}

var stringTestCases = []proverbTest {
{{range .J.Cases}} {
	description: {{printf "%q" .Description }},
	input: {{printf "%#v" .Input.Strings }},
	expected: {{printf "%#v" .Expected }},
},
{{end}}
}
`
