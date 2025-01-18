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
		"score": &[]testCase{},
	}
	if err := gen.Gen("scrabble-score", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Word string `json:"word"`
	} `json:"input"`
	Expected int `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package scrabble

{{.Header}}

type scrabbleTest struct {
	description string
	input       string
	expected    int
}

var scrabbleScoreTests = []scrabbleTest {
{{range .J.score}}{
	description: 	{{printf "%q" .Description}},
	input:  		{{printf "%q" .Input.Word}},
	expected: 		{{printf "%d" .Expected}},
},
{{end}}}
`
