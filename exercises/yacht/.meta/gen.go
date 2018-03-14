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
	if err := gen.Gen("yacht", &j, t); err != nil {
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

// template applied to above data structure generates the Go test cases

type OneCase struct {
	Description string
	Input       struct {
		Dice     []int
		Category string
	}
	Expected int
}

//func (c OneCase)
// Template to generate list of test cases.
var tmpl = `package yacht

{{.Header}}

var testCases = []struct {
	description    string
	dice           []int
	category       string
	expected          int 
}{ {{range .J.Cases}}
{
	description: {{printf "%q"  .Description}},
	dice: {{.Input.Dice | printf "%#v"}} ,
	category: {{printf "%q" .Input.Category}},
	expected: {{printf "%d" .Expected}},
}, 
{{end}}
}
`
