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
	if err := gen.Gen("transpose", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Lines []string
		}
		Expected []string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package transpose

{{.Header}}

var testCases = []struct {
	description string
	input       []string
	expected    []string
}{
{{range .J.Cases}}{
	{{printf "%q" .Description}},
	[]string{
{{range .Input.Lines}}  {{printf "%q" .}},
{{end}}},
	[]string{
{{range .Expected}}  {{printf "%q" .}},
{{end}}},
},
{{end}}}
`
