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
	if err := gen.Gen("reverse-string", &j, t); err != nil {
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
	Input       string
	Expected    string
}

// Template to generate test cases.
var tmpl = `package reverse

{{.Header}}

var testCases = []struct {
	description	string
	input		string
	expected	string
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%q"  .Input}},
	expected:	{{printf "%q"  .Expected}},
},{{end}}
}
`
