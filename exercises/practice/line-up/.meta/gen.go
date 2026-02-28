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
		"format": &[]testCase{},
	}
	if err := gen.Gen("line-up", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Name string `json:"name"`
		Number int `json:"number"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `package lineup

{{.Header}}

var testCases = []struct {
	description string
	name        string
	number      int
	expected    string
}{
{{range .J.format}}{
	description: 	{{printf "%q" .Description}},
	name: 		{{printf "%q" .Input.Name}},
	number:		{{.Input.Number}},
	expected:    	{{printf "%q" .Expected}},
},
{{end}}
}`
