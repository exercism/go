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
