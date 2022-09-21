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
	j := map[string]interface{}{
		"parse": &[]testCase{},
	}
	if err := gen.Gen("markdown", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Markdown string `json:"markdown"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package markdown

{{.Header}}

var testCases = []struct {
	description	  string
	input		  string
	expected	  string
}{ {{range .J.parse}} 
{
	description:	{{printf "%q"  .Description}},
	input:		    {{printf "%q"  .Input.Markdown}},
	expected:	    {{printf "%q"  .Expected}},
},{{end}}
}
`
