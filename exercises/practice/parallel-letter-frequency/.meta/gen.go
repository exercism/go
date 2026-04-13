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
		"calculateFrequencies": &[]testCase{},
	}
	if err := gen.Gen("parallel-letter-frequency", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Texts []string `json:"texts"`
	} `json:"input"`
	Expected map[string]int `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       []string
	expected    map[rune]int
}

var testCases = []testCase { {{range .J.calculateFrequencies}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%#v" .Input.Texts}},
		expected:    map[rune]int{ {{range $k, $v := .Expected}} '{{$k}}': {{$v}}, {{end}} },
	},{{end}}
}`
