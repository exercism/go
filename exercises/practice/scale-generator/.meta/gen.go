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
		"chromatic": &[]testCase{},
		"interval":  &[]testCase{},
	}
	if err := gen.Gen("scale-generator", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Uuid        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Tonic     string `json:"tonic"`
		Intervals string `json:"intervals"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package scale

{{.Header}}

type scaleTest struct {
	description string
	tonic    string
	interval string
	expected []string
}

var scaleTestCases = []scaleTest {
{{range .J.chromatic}}{
	description: {{printf "%q" .Description }},
	tonic: {{printf "%q" .Input.Tonic }},
	interval: {{printf "%q" .Input.Intervals }},
	expected: {{printf "%#v" .Expected }},
},
{{end}}
{{range .J.interval}}{
	description: {{printf "%q" .Description }},
	tonic: {{printf "%q" .Input.Tonic }},
	interval: {{printf "%q" .Input.Intervals }},
	expected: {{printf "%#v" .Expected }},
},
{{end}}
}
`
