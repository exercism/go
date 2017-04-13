// +build ignore

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
	if err := gen.Gen("gigasecond", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []struct {
			Description string
			Input       string
			Expected    string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package gigasecond

{{.Header}}

{{range .J.Cases}}// {{.Description}}
var addCases = []struct {
	description string
	in   string
	want string
}{
{{range .Cases}}{
	{{printf "%q" .Description}},
	{{printf "%q" .Input}},
	{{printf "%q" .Expected}},
},
{{end}}{{end}}
}
`
