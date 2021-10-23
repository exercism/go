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
	if err := gen.Gen("sieve", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Limit int
		}
		Expected []int
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package sieve

{{.Header}}

var testCases = []struct {
	description string
	limit int
	expected []int
}{
{{range .J.Cases}}{
	"{{.Description}}",
	{{.Input.Limit}},
	{{printf "%#v" .Expected}},
},
{{end}}}
`
