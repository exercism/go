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

type Domino [2]int

func (d Domino) String() string {
	return fmt.Sprintf("{%d, %d}", d[0], d[1])
}

// template applied to above data structure generates the Go test cases

type OneCase struct {
	Description string
	Comments    []string
	Input       struct {
		Dominoes []Domino
	}
	Expected bool
}

//func (c OneCase)
// Template to generate list of test cases.
var tmpl = `package dominoes

{{.Header}}

var testCases = []struct {
	description    string
	dominoes  []Domino
	valid          bool     // true => can chain, false => cannot chain
}{ {{range .J.Cases}}
{
	{{printf "%q"  .Description}},
	{{range .Comments}}//{{.}}
	{{end}}[]Domino{ {{range .Input.Dominoes}} {{printf "%v" .}}, {{end}} },
	{{printf "%v"  .Expected}},
}, {{end}}
}
`
