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
	var j js
	if err := gen.Gen("state-of-tic-tac-toe", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []oneCase
	}
}

type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Board []string
	}
	Expected interface{}
}

func (o oneCase) Result() string {
	s, ok := o.Expected.(string)
	if !ok {
		return ""
	}
	return s
}

func (o oneCase) Err() string {
	m, ok := o.Expected.(map[string]interface{})
	if !ok {
		return ""
	}
	return m["error"].(string)
}

// template applied to above data structure generates the Go test cases
var tmpl = `package stateoftictactoe

{{.Header}}

var testCases = []struct{
	description string
	board []string
	expected string
	expectedErr string
} {
{{range .J.Cases}} {{range .Cases}} {
	description: {{printf "%q" .Description}},
	board: {{printf "%#v" .Input.Board }},
	expected: {{printf "%q" .Result }},
	expectedErr:  {{printf "%q" .Err }},
},
{{end}}{{end}}
}
`
