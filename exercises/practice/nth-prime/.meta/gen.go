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
		"prime": &[]testCase{},
	}
	if err := gen.Gen("nth-prime", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) GetExpectedValue() int {
	v, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(v)
}

func (t testCase) GetError() string {
	v, ok := t.Expected.(map[string]any)
	if ok {
		e, ok := v["error"].(string)
		if ok {
			return e
		}
	}
	return ""
}

// template applied to above data structure generates the Go test cases
var tmpl = `package nthprime

{{.Header}}

var tests = []struct {
	description string
	input       int
	expected    int
	err         string
}{
{{range .J.prime}}{
		description: {{printf "%q" .Description}},
		input:       {{printf "%d" .Input.Number}},
		expected:    {{printf "%d" .GetExpectedValue}},
		err:         {{printf "%q" .GetError}},
	},
{{end}}
}
`
