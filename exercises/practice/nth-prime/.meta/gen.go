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
		"prime": &[]testCase{},
	}
	if err := gen.Gen("nth-prime", j, t); err != nil {
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

func (t testCase) GetExpectedValue() int {
	v, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(v)
}

func (t testCase) GetError() string {
	return gen.ErrorMessage(t.Expected)
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       int
	expected    int
	err         string
}

var tests = []testCase { {{range .J.prime}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%d" .Input.Number}},
		expected:    {{printf "%d" .GetExpectedValue}},
		err:         {{printf "%q" .GetError}},
	},{{end}}
}
`
