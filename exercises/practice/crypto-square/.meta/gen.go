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
		"ciphertext": &[]testCase{},
	}
	if err := gen.Gen("crypto-square", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Plaintext string `json:"plaintext"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       string
	expected    string
}
var testCases = []testCase { {{range .J.ciphertext}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%q" .Input.Plaintext}},
		expected:    	{{printf "%q" .Expected}},
	},{{end}}
}`
