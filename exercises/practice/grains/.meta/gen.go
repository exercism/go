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
		"square": &[]testCase{},
		"total":  &[]testCase{}, // expected value for Total() not used from `canonical-data.json`
	}
	if err := gen.Gen("grains", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Square int `json:"square"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) ExpectedValue() uint64 {
	v, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return uint64(v)
}

func (t testCase) ExpectError() bool {
	_, ok := t.Expected.(float64)
	return !ok
}

var tmpl = `package grains

{{.Header}}

// returns the number of grains on the square
var squareTests = []struct {
	description string
	input       int
	expectedVal uint64
	expectError bool
}{
	{{range .J.square}}{
			description: {{printf "%q" .Description}},
			input:       {{printf "%d" .Input.Square}},
			expectedVal: {{printf "%d" .ExpectedValue}},
			expectError: {{printf "%t" .ExpectError}},
		},
	{{end}}
}
`
