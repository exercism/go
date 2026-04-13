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
		"proteins": &[]testCase{},
	}
	if err := gen.Gen("protein-translation", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Strand string `json:"strand"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (tc testCase) ExpectedProteins() string {
	if gen.IsError(tc.Expected) {
		return "nil"
	}
	rawValues, ok := tc.Expected.([]any)
	if !ok {
		return "nil"
	}
	var stringValues []string
	for _, raw := range rawValues {
		stringValues = append(stringValues, raw.(string))
	}
	return fmt.Sprintf("%#v", stringValues)
}

func (tc testCase) ExpectedError() string {
	if gen.IsError(tc.Expected) {
		return "ErrInvalidBase"
	}
	return "nil"
}

var tmpl = `{{.Header}}

type testCase struct {
	description    string
	input          string
	expected       []string
	expectedError  error
}

var testCases = []testCase { {{range .J.proteins}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%q" .Input.Strand}},
		expected:    	{{.ExpectedProteins}},
		expectedError:  {{.ExpectedError}},
	},{{end}}
}`
