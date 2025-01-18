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
	j := map[string]interface{}{
		"rectangles": &[]testCase{},
	}
	if err := gen.Gen("rectangles", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Strings []string `json:"strings"`
	} `json:"input"`
	Expected int `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package rectangles

{{.Header}}

var testCases = []struct {
	description	string
	input       []string
	expected	  int
}{
{{range .J.rectangles}}{
	description:	{{printf "%q"  .Description}},
	input:			[]string{
		{{range .Input.Strings}}{{ printf "%q" .}},
		{{end}}
	},
	expected:		{{printf "%d"  .Expected}},
},
{{end}}
}
`
