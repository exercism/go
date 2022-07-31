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
		"isIsogram" : &[]testCase{},
	}
	if err := gen.Gen("isogram", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string   `json:"description"`
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
