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
	if err := gen.Gen("dominoes", &j, t); err != nil {
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

type Dominoe [2]int

func (d Dominoe) String() string {
	return fmt.Sprintf("Dominoe{%d, %d}", d[0], d[1])
}

// template applied to above data structure generates the Go test cases

type OneCase struct {
	Description string
	Dominoes    []Dominoe `json:"input"`
	Expected    bool
}

//func (c OneCase)
// Template to generate list of test cases.
var tmpl = `package dominoes

{{.Header}}

var testCases = []struct {
	description    string
	dominoes  []Dominoe
	valid          bool     // true => can chain, false => cannot chain
}{ {{range .J.Cases}}
{
	{{printf "%q"  .Description}},
	[]Dominoe{ {{range .Dominoes}} {{printf "%v" .}}, {{end}} },
	{{printf "%v"  .Expected}},
}, {{end}}
}
`
