// +build ignore

package main

import (
	"log"
	"text/template"

	"../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("roman-numerals.json", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Number   int
		Expected string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package romannumerals

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

type romanNumeralTest struct {
	arabic   int
	roman    string
	hasError bool
}

var romanNumeralTests = []romanNumeralTest {
{{range .J.Cases}}{ {{.Number}}, "{{.Expected}}", false},
{{end}}}
`
