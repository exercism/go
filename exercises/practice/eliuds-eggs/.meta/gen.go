package main

import (
	"log"
	"text/template"

	"go/gen" // "../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
		"eggCount": &[]testCase{},
	}
	if err := gen.Gen("eliuds-eggs", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected int `json:"expected"`
}

// Template to generate various test cases.
var tmpl = `package eliudseggs
{{.Header}}
type testCase struct {
	description    string
	input          int
	expected       int
}

var testCases = []testCase{
{{range .J.eggCount}}{
		description:         {{printf "%q" .Description}},
		input:               {{printf "%d" .Input.Number}},
		expected:            {{printf "%d" .Expected}},
	},
{{end}}
}
`
