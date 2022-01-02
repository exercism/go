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
	if err := gen.Gen("isbn-verifier", &j, t); err != nil {
		log.Fatal(err)
	}
}

type OneCase struct {
	Description string
	Input       struct {
		ISBN string
	}
	Expected bool
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []OneCase
}

// template applied to above data structure generates the Go test cases
var tmpl = `package isbn

{{.Header}}

var testCases = []struct {
	isbn string
	expected bool
	description string
}{
{{range .J.Cases}}{ "{{.Input.ISBN}}", {{.Expected}}, "{{.Description}}"},
{{end}}}
`
