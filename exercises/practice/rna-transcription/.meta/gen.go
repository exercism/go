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
	if err := gen.Gen("rna-transcription", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Dna string
		}
		Expected string
	}
}

// Template applied to above data structure generates the Go test cases.
var tmpl = `package strand

{{.Header}}

var rnaTests = []struct {
	description string
	input       string
	expected    string
}{
{{range .J.Cases}}{ "{{.Description}}",
	"{{.Input.Dna}}", "{{.Expected}}",
},
{{end}}}
`
