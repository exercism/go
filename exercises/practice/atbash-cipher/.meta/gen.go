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
		"encode": &[]testCase{},
		"decode": &[]testCase{},
	}
	if err := gen.Gen("atbash-cipher", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Phrase string `json:"phrase"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package atbash

{{.Header}}

var testCases = []struct {
	description string
	phrase      string
	expected    string
}{
{{range .J.encode}}{
		description: {{printf "%q"  .Description}},
		phrase:      {{printf "%q"  .Input.Phrase}},
		expected:    {{printf "%q"  .Expected}},
	},{{end}}
}
`
