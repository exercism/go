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
		"evaluate": &[]testCase{},
		// TODO: add test with property `evaluateBoth` to forth_test.go and generate cases in cases_test.go
	}
	if err := gen.Gen("forth", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Instructions []string `json:"instructions"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) ExpectedNumbers() []int {
	return gen.FloatSliceToInts(t.Expected)
}

func (t testCase) ExplainText() string {
	return gen.ErrorMessage(t.Expected)
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       []string
	expected    []int // nil slice indicates error expected.
	explainText string   // error explanation text
}

var testCases = []testCase { {{range .J.evaluate}}
	{
		description: {{printf "%q"  .Description}},
		input: {{printf "%#v" .Input.Instructions}},
		expected: {{printf "%#v" .ExpectedNumbers}},
		explainText: {{printf "%q"  .ExplainText}},
	},{{end}}
}
`
