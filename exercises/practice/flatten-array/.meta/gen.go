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
		"flatten": &[]testCase{},
	}
	if err := gen.Gen("flatten-array", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Array interface{} `json:"array"`
	} `json:"input"`
	Expected []interface{} `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package flatten

{{.Header}}

var testCases = []struct {
	description	string
	input		interface{}
	expected	[]interface{}
}{ {{range .J.flatten}}
{
	description:	{{printf "%q"  .Description}},
	input:			{{printf "%#v"  .Input.Array}},
	expected:		{{printf "%#v"  .Expected}},
},{{end}}
}
`
