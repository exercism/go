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
		"rotate": &[]testCase{},
	}
	if err := gen.Gen("rotational-cipher", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Text string `json:"text"`
		Shift int `json:"shiftKey"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       string
	shift       int
	expected    string
}

var testCases = []testCase { {{range .J.rotate}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .Input.Text}},
		shift:       {{.Input.Shift}},
		expected:    {{printf "%q" .Expected}},
	},{{end}}
}`
