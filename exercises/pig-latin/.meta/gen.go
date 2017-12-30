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
	if err := gen.Gen("pig-latin", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Cases    []oneLevel
}

// Test levels
type oneLevel struct {
  Description string
  Cases []oneCase
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       string
	Expected    string
}

// Template to generate test cases.
var tmpl = `package igpay

{{.Header}}

var testCases = []struct {
	description	string
	input		    string
	expected	  string
}{ {{range .J.Cases}} {{range .Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%#v"  .Input}},
	expected:	{{printf "%#v"  .Expected}},
},{{end}}{{end}}
}
`
