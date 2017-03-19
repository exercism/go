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
	Find_fewest_coins struct {
		Description []string
		Cases       []struct {
			Description string
			Coins       []int
			Target      int
			Expected    []int
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package change

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

var testCases = []struct {
	description string
	coins       []int
	target	    int
	expected    []int
}{
{{range .J.Find_fewest_coins.Cases}}{
	{{printf "%q" .Description}},
	[]int{
{{range .Coins}}  {{printf "%v" .}},
{{end}}},
	{{printf "%d" .Target}},
	[]int{
{{range .Expected}}  {{printf "%v" .}},
{{end}}},
},
{{end}}}
`
