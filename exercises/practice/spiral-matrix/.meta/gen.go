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
		"spiralMatrix": &[]testCase{},
	}
	if err := gen.Gen("spiral-matrix", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Size int `json:"size"`
	} `json:"input"`
	Expected [][]int `json:"expected"`
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	input       int
	expected    [][]int
}{ {{range .J.spiralMatrix}}
	{
		description: {{printf "%q" .Description}},
		input:       {{.Input.Size}},
		expected:    [][]int{ {{range .Expected}}
			{{printf "%#v" .}}, {{end}}
		},
	},{{end}}
}`
