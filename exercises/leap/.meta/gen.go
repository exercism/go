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
	if err := gen.Gen("leap", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Year int
		}
		Expected bool
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package leap

{{.Header}}

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
{{range .J.Cases}}{ {{.Input.Year}}, {{.Expected}}, {{printf "%q" .Description}}},
{{end}}}
`
