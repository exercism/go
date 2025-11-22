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
		"findAnagrams": &[]testCase{},
	}
	if err := gen.Gen("anagram", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Subject    string   `json:"subject"`
		Candidates []string `json:"candidates"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package anagram

{{.Header}}

type anagramTest struct {
	description string
	subject     string
	candidates  []string
	expected    []string
}

var testCases = []anagramTest{ 
	{{range .J.findAnagrams}}{
		description: {{printf "%q"   .Description}},
		subject: {{printf "%q"       .Input.Subject}},
		candidates: {{printf "%#v"   .Input.Candidates}},
		expected: {{printf "%#v"     .Expected}},
	},
	{{end}}
}
`
