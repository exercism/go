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
		"isValid": &[]testCase{},
	}
	if err := gen.Gen("isbn-verifier", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		ISBN string `json:"isbn"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

var tmpl = `package isbn

{{.Header}}

var testCases = []struct {
	description string
	isbn        string
	expected    bool
}{
{{range .J.isValid}}{
			description: 	{{printf "%q"  .Description}},
			isbn: 			{{printf "%q"  .Input.ISBN}},
			expected:   	{{printf "%t"  .Expected}},
	},
{{end}}
}
`
