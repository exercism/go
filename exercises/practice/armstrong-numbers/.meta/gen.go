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
		"isArmstrongNumber": &[]testCase{},
	}
	if err := gen.Gen("armstrong-numbers", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

// Template to generate test cases.
var tmpl = `{{.Header}}

var testCases = []struct {
	description	string
	input		int
	expected	bool
}{ {{range .J.isArmstrongNumber}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%#v"  .Input.Number}},
	expected:	{{printf "%#v"  .Expected}},
},{{end}}
}
`
