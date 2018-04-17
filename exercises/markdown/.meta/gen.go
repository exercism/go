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
	if err := gen.Gen("markdown", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Comments []string
	Cases    []oneCase
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Markdown string
	}
	Expected string
}

// Template to generate test cases.
var tmpl = `package markdown

{{.Header}}

{{range .J.Comments}}
// {{ . }}
{{end}}

var testCases = []struct {
	description	  string
	input		  string
	expected	  string
}{ {{range .J.Cases}} 
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%#v"  .Input.Markdown}},
	expected:	{{printf "%#v"  .Expected}},
},{{end}}
}
`
