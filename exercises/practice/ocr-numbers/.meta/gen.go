package main

import (
	"../../../../gen"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Funcs(template.FuncMap{"tolower": strings.ToLower}).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"convert": &[]testCase{},
	}
	if err := gen.Gen("ocr-numbers", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Rows []string `json:"rows"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (tc testCase) IsError() bool {
	return gen.IsError(tc.Expected)
}

func (tc testCase) ErrorMessage() string {
	return gen.ErrorMessage(tc.Expected)
}

func (tc testCase) ParseInput() string {
	return "\n" + strings.Join(tc.Input.Rows, "\n")
}

func (tc testCase) ParseExpected() []string {
	return strings.Split(tc.Expected.(string), ",")
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       string
	expected    []string
	errorMsg    string
}

var testCases = []testCase { {{range .J.convert}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .ParseInput}},{{if .IsError}}
		errorMsg:    {{tolower .ErrorMessage | printf "%q"}},{{else}}
		expected:    {{printf "%#v" .ParseExpected}},{{end}}
	},{{end}}
}`
