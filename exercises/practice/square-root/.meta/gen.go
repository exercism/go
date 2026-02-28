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
		"squareRoot": &[]testCase{},
	}
	if err := gen.Gen("square-root", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Radicand int `json:"radicand"`
	} `json:"input"`
	Expected int `json:"expected"`
}

var tmpl = `package squareroot

{{.Header}}

var testCases = []struct {
	description string
	input       int
	expected    int
}{
{{range .J.squareRoot}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{printf "%d" .Input.Radicand}},
	expected:    	{{.Expected}},
},
{{end}}
}`
