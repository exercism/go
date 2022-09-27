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
		"total": &[]testCase{},
	}
	if err := gen.Gen("book-store", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Basket []int `json:"basket"`
	} `json:"input"`
	Expected int `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package bookstore

{{.Header}}

var testCases = []struct {
	description string
	basket      []int
	expected    int
}{
{{range .J.total}} {
	description: {{printf "%q"  .Description}},
	basket:      {{printf "%#v" .Input.Basket}},
	expected:    {{printf "%d"  .Expected}},
},
{{end}}}
`
