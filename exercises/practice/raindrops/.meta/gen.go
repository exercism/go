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
		"convert": &[]testCase{},
	}
	if err := gen.Gen("raindrops", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       int
	expected    string
}{
{{range .J.convert}}{ 
	description: 	{{printf "%q" .Description}},
	input: 			{{printf "%d" .Input.Number}},
	expected: 		{{printf "%q" .Expected}},
},
{{end}}
}
`
