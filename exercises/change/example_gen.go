// +build ignore

package main

import (
	"log"
	"text/template"

	"../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("change", &j, t); err != nil {
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
	Property    string
	Coins       []int
	Target      int
	Expected    interface{}
}

var tmpl = `package change

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

var testCases = []struct {
	description string
	coins       []int
	target	    int
	expected    interface{}
}{
{{range .J.Cases}}{
	{{printf "%q" .Description}},
	[]int{
{{range .Coins}}  {{printf "%v" .}},
{{end}}},
	{{printf "%d" .Target}},
	{{printf "%#v" .Expected}},
},
{{end}}}
`
