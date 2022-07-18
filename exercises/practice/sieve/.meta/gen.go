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
		"primes": &[]testCase{},
	}
	if err := gen.Gen("sieve", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Limit int `json:"limit"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package sieve

{{.Header}}

var testCases = []struct {
	description string
	limit int
	expected []int
}{
{{range .J.primes}}{
	description: 	{{printf "%q"  .Description}},
	limit: 			{{printf "%d"  .Input.Limit}},
	expected: 		{{printf "%#v" .Expected}},
},
{{end}}}
`
