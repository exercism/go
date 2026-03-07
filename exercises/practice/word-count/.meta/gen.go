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
		"countWords": &[]testCase{},
	}
	if err := gen.Gen("word-count", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Sentence string `json:"sentence"`
	} `json:"input"`
	Expected map[string]int `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

var testCases = []struct {
    description string
	input       string
	expected      Frequency
}{
{{range .J.countWords}}{
	description: 	{{printf "%q" .Description}},
	input: 			{{printf "%q" .Input.Sentence}},
	expected: 		Frequency{ {{range $key, $val := .Expected}} {{printf "%q: %d, " $key $val}} {{end}} },
},
{{end}}}
`
