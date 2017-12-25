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
	if err := gen.Gen("binary-search", &j, t); err != nil {
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
		Array []int
		Value int
	}
	Expected int
}

// Template to generate test cases.
var tmpl = `package binarysearch

{{.Header}}

var testCases = []struct {
	description	string
	slice		[]int
	key		int
	x	int
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	slice:		{{printf "%#v" .Input.Array}},
	key:		{{printf "%d"  .Input.Value}},
	x:	{{printf "%d"  .Expected}},
},{{end}}
}
`
