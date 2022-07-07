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
		return "\"\""
	}
	var res string
	switch s {
	case "win":
		res = "state_win"
	case "ongoing":
		res = "state_ongoing"
	case "draw":
		res = "state_draw"
	}
	return res
}

func (o oneCase) Err() string {
	_, ok := o.Expected.(map[string]interface{})
	if !ok {
		return "false"
	}
	return "true"
}

// template applied to above data structure generates the Go test cases
var tmpl = `package stateoftictactoe

{{.Header}}

var testCases = []struct{
	description string
	board []string
	expected state
	wantErr bool
} {
{{range .J.Cases}} {{range .Cases}} {
	description: {{printf "%q" .Description}},
	board: {{printf "%#v" .Input.Board }},
	expected: {{ .Result }},
	wantErr:  {{ .Err }},
},
{{end}}{{end}}
}
`
