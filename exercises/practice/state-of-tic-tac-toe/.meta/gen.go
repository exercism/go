package main

import (
	"log"
	"strings"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
		"gamestate": &[]testCase{},
	}
	if err := gen.Gen("state-of-tic-tac-toe", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Board []string `json:"board"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (o testCase) ExpectedResult() string {
	s, ok := o.Expected.(string)
	if !ok {
		return "\"\""
	}
	return strings.Title(s)
}

func (o testCase) ExpectError() bool {
	_, ok := o.Expected.(map[string]interface{})
	return ok
}

// template applied to above data structure generates the Go test cases
var tmpl = `package stateoftictactoe

{{.Header}}

var testCases = []struct{
	description string
	board []string
	expected State
	wantErr bool
} {
{{range .J.gamestate}} {
	description: 	{{printf "%q" 	.Description}},
	board: 			{{printf "%#v" 	.Input.Board }},
	expected: 		{{printf "%s" 	.ExpectedResult }},
	wantErr:  		{{printf "%t" 	.ExpectError }},
},
{{end}}
}
`
