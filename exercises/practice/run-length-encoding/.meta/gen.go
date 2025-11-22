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
		"encode":      &[]testCase{},
		"decode":      &[]testCase{},
		"consistency": &[]testCase{},
	}
	if err := gen.Gen("run-length-encoding", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		String string `json:"string"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `package encode

{{.Header}}

type testCase struct {
	description string
	input       string
	expected    string
}

// run-length encode a string
var encodeTests = []testCase{
	{{range .J.encode}}{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .Input.String}},
		expected:    {{printf "%q" .Expected}},
	},
	{{end}}
}

// run-length decode a string
var decodeTests = []testCase{
	{{range .J.decode}}{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .Input.String}},
		expected:    {{printf "%q" .Expected}},
	},
	{{end}}
}

// encode and then decode
var encodeDecodeTests = []testCase{
	{{range .J.consistency}}{
		description: {{printf "%q" .Description}},
		input:       {{printf "%q" .Input.String}},
		expected:    {{printf "%q" .Expected}},
	},
	{{end}}
}
`
