package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"isPangram": &[]testCase{},
	}
	if err := gen.Gen("pangram", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Sentence string `json:"sentence"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package pangram

{{.Header}}

var testCases = []struct {
	description	string
	input		string
	expected	bool
}{ {{range .J.isPangram}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%q"  .Input.Sentence}},
	expected:	{{printf "%t"  .Expected}},
},{{end}}
}
`
