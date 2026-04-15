package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
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
	Parents     []string
	Input       struct {
		Instructions []string `json:"instructions"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (tc testCase) ExtendedDescription() string {
	return fmt.Sprintf("%q", strings.Join(append(tc.Parents, tc.Description), " -> "))
}

func (tc testCase) ExpectedNumbers() []int {
	return gen.FloatSliceToInts(tc.Expected)
}

func (tc testCase) ExplainText() string {
	return gen.ErrorMessage(tc.Expected)
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
		description: {{.ExtendedDescription}},
		input: {{printf "%#v" .Input.Instructions}},
		expected: {{printf "%#v" .ExpectedNumbers}},
		explainText: {{printf "%q"  .ExplainText}},
	},{{end}}
}
`
