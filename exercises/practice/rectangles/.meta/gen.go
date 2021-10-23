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
	if err := gen.Gen("rectangles", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Cases    []oneCase
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Strings []string
	}
	Expected int
}

// Template to generate test cases.
var tmpl = `package rectangles

{{.Header}}

var testCases = []struct {
	description	string
	input       []string
	expected	  int
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		[]string{
		{{range .Input.Strings}}{{printf "%q" .}},
		{{end}}
	},
	expected:	{{printf "%d"  .Expected}},
},{{end}}
}
`
