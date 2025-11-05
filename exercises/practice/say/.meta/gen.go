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
		"say": &[]testCase{},
	}
	if err := gen.Gen("say", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) ExpectedValue() string {
	v, ok := t.Expected.(string)
	if ok {
		return v
	}
	return ""
}

func (t testCase) ExpectError() bool {
	v, ok := t.Expected.(map[string]any)
	if ok {
		_, ok := v["error"]
		return ok
	}
	return false
}

// Template to generate test cases.
var tmpl = `package say

{{.Header}}

var testCases = []struct {
	description	string
	input		int64
	expected	string
	expectError	bool
}{ {{range .J.say}}
{
	description:	{{printf "%q"  .Description}},
	input:		{{printf "%v"  .Input.Number}},
	expected: 	{{printf "%q"  .ExpectedValue}},
	expectError: {{printf "%t"  .ExpectError}},
},{{end}}
}
`
