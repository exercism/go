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
	if err := gen.Gen("anagram", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Comments []string
	Cases    []OneCase
}

// The JSON structure we expect to be able to unmarshal into
type OneCase struct {
	Description string
	Subject     string
	Candidates  []string
	Expected    []string
}

// template applied to above data structure generates the Go test cases
var tmpl = `package anagram

{{.Header}}

var testCases = []struct {
	description string
	subject     string
	candidates  []string
	expected    []string
}{ {{range .J.Cases}}
{
	description: {{printf "%q"   .Description}},
	subject: {{printf "%q"   .Subject}},
	candidates: []string { {{range $line := .Candidates}}{{printf "\n%q," $line}}{{end}}},
	expected: []string { {{range $line := .Expected}}{{printf "\n%q," $line}}{{end}}},
	},{{end}}
}
`
