// +build ignore

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
	if err := gen.Gen("scrabble-score", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       string
		Expected    int
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package scrabble

{{.Header}}

type scrabbleTest struct {
	input       string
	expected    int
}

var scrabbleScoreTests = []scrabbleTest {
{{range .J.Cases}}{ "{{.Input}}", {{.Expected}}}, // {{.Description}}
{{end}}}
`
