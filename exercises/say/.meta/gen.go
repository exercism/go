package main

import (
	"fmt"
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
	if err := gen.Gen("say", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Comments []string
	Cases    []OneCase
}

// Test cases
type OneCase struct {
	Description string
	Property    string
	Input       struct {
		Number int64
	}
	Expected interface{}
}

func (c OneCase) ErrorExpected() bool {
	switch value := c.Expected.(type) {
	case float64: // you'd think int but no, JSON uses float for numbers
		if value == -1 {
			return true
		}
	case string:
		return false
	}

	panic(fmt.Sprintf("Unexpected error value: %T => %v", c.Expected, c.Expected))
}

// Template to generate test cases.
var tmpl = `package say

{{.Header}}

var testCases = []struct {
	description	string
	input		int64
	expected	string
	expectError	bool
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%v"  .Input.Number}},
	{{if .ErrorExpected}}expectError:	true,
	{{else}}expected:	{{printf "%q"  .Expected}},
	{{- end}}
},{{end}}
}
`
