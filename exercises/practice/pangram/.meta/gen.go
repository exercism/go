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
	if err := gen.Gen("pangram", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Cases    []struct {
		Cases []oneCase
	}
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Sentence string
	}
	Expected bool
}

// Template to generate test cases.
var tmpl = `package pangram

{{.Header}}

var testCases = []struct {
	description	string
	input		string
	expected	bool
}{ {{range .J.Cases}}{{range .Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%q"  .Input.Sentence}},
	expected:	{{printf "%t"  .Expected}},
},{{end}}{{end}}
}
`
