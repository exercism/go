package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
		"toRna": &[]testCase{},
	}
	if err := gen.Gen("rna-transcription", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		DNA string `json:"dna"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// Template applied to above data structure generates the Go test cases.
var tmpl = `package strand

{{.Header}}

var testCases = []struct {
	description string
	input       string
	expected    string
}{
{{range .J.toRna}}{ 
	description: 	{{printf "%q" .Description}},
	input: 		 	{{printf "%q" .Input.DNA}},
	expected: 		{{printf "%q" .Expected}},
},
{{end}}}
`
