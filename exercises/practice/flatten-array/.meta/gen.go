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
	if err := gen.Gen("flatten-array", &j, t); err != nil {
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
		Array interface{}
	}
	Expected []interface{}
}

// Template to generate test cases.
var tmpl = `package flatten

{{.Header}}

var testCases = []struct {
	description	string
	input		interface{}
	expected	[]interface{}
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%#v"  .Input.Array}},
	expected:	{{printf "%#v"  .Expected}},
},{{end}}
}
`
