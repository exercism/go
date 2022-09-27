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
		"commands": &[]testCase{},
	}
	if err := gen.Gen("secret-handshake", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package secret

{{.Header}}

type secretHandshakeTest struct {
	description string
	input       uint
	expected    []string
}

var tests = []secretHandshakeTest {
{{range .J.commands}}{ 
	description: 	{{printf "%q" .Description }},
	input: 			{{printf "%d" .Input.Number }}, 
	expected: 		{{printf "%#v" .Expected }},
},
{{end}}
}
`
