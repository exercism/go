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
		"tick": &[]testCase{},
	}
	if err := gen.Gen("game-of-life", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Matrix [][]int `json:"matrix"`
	} `json:"input"`
	Expected [][]int `json:"expected"`
}

var tmpl = `package gameoflife

{{.Header}}

var testCases = []struct {
	description string
	input       [][]int
	expected    [][]int
}{
{{range .J.tick}}{
	description: {{printf "%q" .Description}},
	input:       [][]int{
		{{range .Input.Matrix}} { {{range .}} {{.}}, {{end}}},
		{{end}}
	},
	expected:       [][]int{
		{{range .Expected}} { {{range .}} {{.}}, {{end}}},
		{{end}}
	},
},
{{end}}
}`
