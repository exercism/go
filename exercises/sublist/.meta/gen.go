package main

import (
	"log"
	"text/template"

	".."
	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("sublist", &j, t); err != nil {
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
	ListOne     []int `json:"listOne"`
	ListTwo     []int `json:"listTwo"`
	Expected    sublist.Relation
}

// Template to generate test cases.
var tmpl = `package sublist

{{.Header}}

var testCases = []struct {
	description	string
	listOne		[]int
	listTwo		[]int
	expected	Relation
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	listOne:		{{printf "%#v"  .ListOne}},
	listTwo:		{{printf "%#v"  .ListTwo}},
	expected:	{{printf "%#v"  .Expected}},
},{{end}}
}
`
