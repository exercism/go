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
		"factors": &[]testCase{},
	}
	if err := gen.Gen("prime-factors", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Value int64 `json:"value"`
	} `json:"input"`
	Expected []int64 `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package prime

{{.Header}}

var testCases = []struct {
	description string
	input       int64
	expected    []int64
}{
{{range .J.factors}} {
	description: 	{{printf "%q" .Description}},
	input: 			{{printf "%#v" .Input.Value}},
	expected: 		{{printf "%#v" .Expected}},
},
{{end}}
}
`
