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
	if err := gen.Gen("isogram", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Cases    []struct {
		Description string
		Cases       []oneCase
	}
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Phrase string
	}
	Expected bool
}

// Template to generate test cases.
var tmpl = `package isogram

{{.Header}}

var testCases = []struct {
	description	string
	input		string
	expected	bool
}{ {{range .J.Cases}} {{range .Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%q"  .Input.Phrase}},
	expected:	{{printf "%t"  .Expected}},
},{{end}}{{end}}
}
`
