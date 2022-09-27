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
		"winner": &[]testCase{},
	}
	if err := gen.Gen("connect", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Board []string `json:"board"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package connect

{{.Header}}

var testCases = []struct {
    description string
    board       []string
    expected    string
}{
	{{range .J.winner}}{
		description: {{printf "%q" .Description}},
		board: []string{
			{{range .Input.Board}} {{printf "%q" .}},
			{{end}} 
		},
		expected: {{printf "%q" .Expected}},
},
{{end}}
}
`
