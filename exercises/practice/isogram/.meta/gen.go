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
		"isIsogram": &[]testCase{},
	}
	if err := gen.Gen("isogram", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Phrase string `json:"phrase"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

var tmpl = `package isogram

{{.Header}}

var testCases = []struct {
	description	string
	input		string
	expected	bool
}{ {{range .J.isIsogram}} 
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%q"  .Input.Phrase}},
	expected:	{{printf "%t"  .Expected}},
},{{end}}
}
`
