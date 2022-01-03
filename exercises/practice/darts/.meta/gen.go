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
	if err := gen.Gen("darts", &j, t); err != nil {
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
		X float64
		Y float64
	}
	Expected int
}

// Template to generate test cases.
var tmpl = `package darts

{{.Header}}

var testCases = []struct {
	description	string
	x		float64
	y		float64
	expected	int
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	x:		{{printf "%.1f"  .Input.X}},
	y:		{{printf "%.1f"  .Input.Y}},
	expected:	{{printf "%d"  .Expected}},
},{{end}}
}
`
