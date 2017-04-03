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
	if err := gen.Gen("rna-transcription", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Dna         string
		Expected    string
	}
}

// Template applied to above data structure generates the Go test cases.
//
// Note x-common test cases have some test cases that are not used here.
//
// It has some with invalid DNA inputs indicated with null expected values.
// The current API does not have an error or ok return.  It could be changed
// but it's simple for now, leaving the result undefined for invalid input.
// So these test cases are skipped here.
//
// Also the x-common test cases include some test cases in the reverse
// direction, from rna to dna.  These are not called for by the exercise
// readme and have no biological basis and so are not converted here.
var tmpl = `package strand

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

var rnaTests = []struct {
	input    string
	expected string
}{
{{range .J.Cases}} {{if and .Dna .Expected}} // {{.Description}}
	{"{{.Dna}}", "{{.Expected}}"},

{{end}}{{end}}}
`
