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
		"rows": &[]testCase{},
	}
	if err := gen.Gen("pascals-triangle", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Count int `json:"count"`
	} `json:"input"`
	Expected [][]int `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	rows        int
	expected    [][]int
}

var testCases = []testCase { {{range .J.rows}}
	{
		description: {{printf "%q" .Description}},
		rows:        {{.Input.Count}},
		expected:    [][]int{ {{range .Expected}}
			{{printf "%#v" .}},{{end}}
		},
	},{{end}}
}`
