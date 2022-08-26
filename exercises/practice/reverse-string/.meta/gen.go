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
	var j = map[string]interface{}{
		"reverse" : &[]testCase{},
	}
	if err := gen.Gen("reverse-string", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Value string `json:"value"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package reverse

{{.Header}}

type reverseTestCase struct {
	description	string
	input		string
	expected	string
}

var testCases = []reverseTestCase{ {{range .J.reverse}}
{
	description:	{{printf "%q"  .Description}},
	input:			{{printf "%q"  .Input.Value}},
	expected:		{{printf "%q"  .Expected}},
},{{end}}
}
`
