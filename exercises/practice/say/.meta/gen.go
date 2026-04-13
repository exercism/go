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
	return gen.IsError(t.Expected)
}

// Template to generate test cases.
var tmpl = `{{.Header}}

type testCase struct {
	description	string
	input		int64
	expected	string
	expectError	bool
}

var testCases = []testCase { {{range .J.say}}
	{
		description:	{{printf "%q"  .Description}},
		input:		{{printf "%v"  .Input.Number}},
		expected: 	{{printf "%q"  .ExpectedValue}},
		expectError: {{printf "%t"  .ExpectError}},
	},{{end}}
}
`
