package main

import (
	"log"
	"text/template"

	space ".."
	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("space-age", &j, t); err != nil {
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
		Planet  space.Planet
		Seconds float64
	}
	Expected float64
}

// Template to generate test cases.
var tmpl = `package space

{{.Header}}

var testCases = []struct {
	description	string
	planet		Planet
	seconds		float64
	expected	float64
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	planet:		{{printf "%#v"  .Input.Planet}},
	seconds:		{{printf "%.0f"  .Input.Seconds}},
	expected:	{{printf "%.2f"  .Expected}},
},{{end}}
}
`
