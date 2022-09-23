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
	j := map[string]interface{}{
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
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectedValue() int {
	v, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(v)
}

func (t testCase) ExpectedError() string {
	v, ok := t.Expected.(map[string]interface{})
	if ok {
		e, ok := v["error"].(string)
		if ok {
			return e
		}
	}
	return ""
}

var tmpl = `package lsproduct

{{.Header}}

var testCases = []struct {
	description string
	digits      string
	span        int
	expected    int64
	error       string
}{
{{range .J.largestProduct}}{
		description: {{printf "%q" .Description}},
		digits:      {{printf "%q" .Input.Digits}},
		span:        {{printf "%d" .Input.Span}},
		expected:    {{printf "%d" .ExpectedValue}},
		error:       {{printf "%q" .ExpectedError}},
	},
{{end}}
}
`
