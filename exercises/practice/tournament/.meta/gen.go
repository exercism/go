package main

import (
	"../../../../gen"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Funcs(template.FuncMap{"join": strings.Join}).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"tally": &[]testCase{},
	}
	if err := gen.Gen("tournament", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Rows []string `json:"rows"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `{{.Header}}

type TestCase = struct{
	description string
	input       string
	expected    string
}

var testCases = []TestCase{ {{range .J.tally}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        ` + "`" + `
{{join .Input.Rows "\n"}}
` + "`" + `,
		expected:	        ` + "`" + `
{{join .Expected "\n"}}
` + "`" + `,
	},{{end}}
}`
