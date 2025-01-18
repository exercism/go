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
		"leapYear": &[]testCase{},
	}
	if err := gen.Gen("leap", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Year int `json:"year"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package leap

{{.Header}}

var testCases = []struct {
	description string
	year        int
	expected    bool
}{
{{range .J.leapYear}}{
		description: {{printf "%q" .Description}},
		year:        {{printf "%d" .Input.Year}},
		expected:    {{printf "%t" .Expected}},
	},
{{end}}
}
`
