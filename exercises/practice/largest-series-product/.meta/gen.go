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
		"largestProduct": &[]testCase{},
	}
	if err := gen.Gen("largest-series-product", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Digits string `json:"digits"`
		Span   int    `json:"span"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) ExpectedValue() int {
	v, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(v)
}

func (t testCase) ExpectedError() string {
	return gen.ErrorMessage(t.Expected)
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	digits      string
	span        int
	expected    int64
	error       string
}

var testCases = []testCase { {{range .J.largestProduct}}
	{
		description: {{printf "%q" .Description}},
		digits:      {{printf "%q" .Input.Digits}},
		span:        {{printf "%d" .Input.Span}},
		expected:    {{printf "%d" .ExpectedValue}},
		error:       {{printf "%q" .ExpectedError}},
	},{{end}}
}
`
