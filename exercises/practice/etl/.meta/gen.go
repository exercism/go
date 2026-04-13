package main

import (
	"../../../../gen"
	"log"
	"strconv"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"transform": &[]testCase{},
	}
	if err := gen.Gen("etl", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Legacy map[string][]string `json:"Legacy"`
	} `json:"input"`
	Expected map[string]int `json:"expected"`
}

func (tc *testCase) Legacy() map[int][]string {
	legacy := make(map[int][]string)
	for key, values := range tc.Input.Legacy {
		num, err := strconv.Atoi(key)
		if err != nil {
			panic("Could not parse the data")
		}
		legacy[num] = values
	}
	return legacy
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	input       map[int][]string
	expected    map[string]int
}

var testCases = []testCase { {{range .J.transform}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%#v" .Legacy}},
		expected:    {{printf "%#v" .Expected}},
	},{{end}}
}`
