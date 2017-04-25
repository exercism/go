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
	if err := gen.Gen("hamming", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Strand1     string
		Strand2     string
		Expected    int
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package hamming

{{.Header}}

var testCases = []struct {
	s1   string
	s2   string
	want int
}{
{{range .J.Cases}}{ // {{.Description}}
	{{printf "%q" .Strand1}},
	{{printf "%q" .Strand2}},
	{{.Expected}},
},
{{end}}}
`
