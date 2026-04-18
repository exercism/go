package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"slices": &[]testCase{},
	}
	if err := gen.Gen("series", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Series      string `json:"series"`
		SliceLength int    `json:"sliceLength"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (tc testCase) IsError() bool {
	return gen.IsError(tc.Expected)
}

func (tc testCase) ErrMsg() string {
	return gen.ErrorMessage(tc.Expected)
}

func (tc testCase) GetExpected() []string {
	if gen.IsError(tc.Expected) {
		return nil
	}
	expected, ok := tc.Expected.([]any)
	if !ok {
		panic(fmt.Sprintf("cannot parse expected for %s", tc.Description))
	}
	out := make([]string, len(expected))
	for i, v := range expected {
		if s, ok := v.(string); !ok {
			panic(fmt.Sprintf("cannot parse expected for %s", tc.Description))
		} else {
			out[i] = s
		}
	}
	return out
}

var tmpl = `{{.Header}}

type simpleTestCase struct {
	description string
	digits      int
	input       string
	expected    string
}

var simpleTestCases = []simpleTestCase { {{range .J.slices}}{{if not .IsError}}
	{
		description: {{printf "%q" .Description}},
		digits:      {{.Input.SliceLength}},
		input:       {{printf "%q" .Input.Series}},
		expected:    "{{index .GetExpected 0}}",
	},{{end}}{{end}}
}

type testCase struct {
	description string
	digits      int
	input       string
	expected    []string
	errorMsg    string
}

var testCases = []testCase { {{range .J.slices}}
	{
		description: {{printf "%q" .Description}},
		digits:      {{.Input.SliceLength}},
		input:       {{printf "%q" .Input.Series}},
		expected:    {{printf "%#v" .GetExpected}},{{if .IsError}}
		errorMsg:    {{printf "%q" .ErrMsg}},{{end}}
	},{{end}}
}`
