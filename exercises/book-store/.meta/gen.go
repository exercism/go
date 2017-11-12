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
	if err := gen.Gen("book-store", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases  []OneCase
	Groups []TestGroup `json:"cases"`
}

type TestGroup struct {
	Description string
	Cases       []OneCase
}

type OneCase struct {
	Description string
	Basket      []int
	Expected    float64
}

// template applied to above data structure generates the Go test cases
var tmpl = `package bookstore

{{.Header}}

var testCases = []struct {
	description string
	basket      []int
	expected    float64
}{{range .J.Groups}}{
{{range .Cases}} {
	description: "{{.Description}}",
	basket:      {{printf "%#v" .Basket}},
	expected:    {{printf "%.2f" .Expected}},
},
{{end}}}{{end}}
`
