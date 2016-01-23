// +build ignore

package main

import (
	"log"
	"text/template"

	"../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("gigasecond.json", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Add struct {
		Description []string
		Cases       []struct {
			Input    string
			Expected string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package gigasecond

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}
{{range .J.Add.Description}}// {{.}}
{{end}}var addCases = []struct {
	in   string
	want string
}{
{{range .J.Add.Cases}}{
	{{printf "%q" .Input}},
	{{printf "%q" .Expected}},
},
{{end}}}
`
