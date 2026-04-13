package main

import (
	"../../../../gen"
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
		"sumOfSquares": &[]testCase{},
		"squareOfSum": &[]testCase{},
		"differenceOfSquares": &[]testCase{},
	}
	if err := gen.Gen("difference-of-squares", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected int `json:"expected"`
}

func (tc testCase) Func() string {
	n := strings.ToUpper(tc.Property[0:1]) + tc.Property[1:]
	if n == "DifferenceOfSquares" {
		n = "Difference"
	}
	return n
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	fn          func(int) int
	fnName      string
	input       int
	expected    int
}

var testCases = []testCase { {{range .J.sumOfSquares}}
	{
		description: {{printf "%q" .Description}},
		fn:          {{.Func}},
		fnName:      {{printf "%q" .Func}},
		input:       {{.Input.Number}},
		expected:    {{.Expected}},
	},{{end}} {{range .J.squareOfSum}}
	{
		description: {{printf "%q" .Description}},
		fn:          {{.Func}},
		fnName:      {{printf "%q" .Func}},
		input:       {{.Input.Number}},
		expected:    {{.Expected}},
	},{{end}} {{range .J.differenceOfSquares}}
	{
		description: {{printf "%q" .Description}},
		fn:          {{.Func}},
		fnName:      {{printf "%q" .Func}},
		input:       {{.Input.Number}},
		expected:    {{.Expected}},
	},{{end}}
}`
