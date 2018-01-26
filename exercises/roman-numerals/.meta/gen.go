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
	if err := gen.Gen("roman-numerals", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Input struct {
			Number int
		}
		Expected string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package romannumerals

{{.Header}}

type romanNumeralTest struct {
	arabic   int
	roman    string
	hasError bool
}

var romanNumeralTests = []romanNumeralTest {
{{range .J.Cases}}{ {{.Input.Number}}, "{{.Expected}}", false},
{{end}}}
`
