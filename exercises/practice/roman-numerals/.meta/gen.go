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
	var j = map[string]interface{}{
		"roman": &[]testCase{},
	}
	if err := gen.Gen("roman-numerals", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package romannumerals

{{.Header}}

type romanNumeralTest struct {
	description string
	input       int
	expected    string
}

var validRomanNumeralTests = []romanNumeralTest{
	{{range .J.roman}}{
		description: {{printf "%q" .Description}},
		input:       {{printf "%d" .Input.Number}},
		expected:    {{printf "%q" .Expected}},
	},
	{{end}}
}
`
