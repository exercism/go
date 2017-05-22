package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("connect", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Board       []string
		Expected    string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package connect

{{.Header}}

var testCases = []struct {
    description string
    board       []string
    expected    string
}{
	{{range .J.Cases}}{
		description: {{printf "%q" .Description}},
		board: []string{ {{range $line := .Board}}{{printf "\n    %q," $line}}{{end}}},
		expected: {{printf "%q" .Expected}},
},
{{end}}}
`
