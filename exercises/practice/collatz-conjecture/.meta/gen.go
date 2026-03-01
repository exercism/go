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
		"steps": &[]testCase{},
	}
	if err := gen.Gen("collatz-conjecture", j, t); err != nil {
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

func (t testCase) ExpectedValue() int {
	val, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(val)
}

func (t testCase) ExpectError() bool {
	_, hasExpectedNumber := t.Expected.(float64)
	return !hasExpectedNumber
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       int
	expectError bool
	expected    int
}{
{{range .J.steps}}{
	description:	{{printf "%q"  .Description}},
	input:			{{printf "%d"  .Input.Number}},
	expectError: 	{{printf "%v"  .ExpectError}},
	expected: 		{{printf "%d"  .ExpectedValue}},
},
{{end}}
}
`
