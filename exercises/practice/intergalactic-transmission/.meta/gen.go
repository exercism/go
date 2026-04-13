package main

import (
	"fmt"
	"log"
	"strings"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"transmitSequence": &[]testCase{},
		"decodeMessage":    &[]testCase{},
	}
	if err := gen.Gen("intergalactic-transmission", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Message []string `json:"message"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func toCommaSeparatedString(values []interface{}) string {
	asStrings := make([]string, len(values))
	for i, value := range values {
		asStrings[i] = fmt.Sprintf("%v", value)
	}

	return strings.Join(asStrings, ",")
}

func (t testCase) FormatInput() string {
	return fmt.Sprintf("[]byte{%s}", strings.Join(t.Input.Message, ", "))
}

func (t testCase) FormatExpectedValue() string {
	if expected, ok := t.Expected.([]interface{}); ok {
		return fmt.Sprintf("[]byte{%s}", toCommaSeparatedString(expected))
	}
	return "[]byte{}"
}

func (t testCase) FormatExpectedError() string {
	return gen.ErrorMessage(t.Expected)
}

var tmpl = `{{.Header}}

type testCaseTransmit struct {
	description string
	input       []byte
	expected    []byte
}

var transmitCases = []testCaseTransmit { {{range .J.transmitSequence}}
	{
		description: 	{{printf "%q" .Description}},
		input: 			{{.FormatInput}},
		expected: 	    {{.FormatExpectedValue}},
	}, {{end}}
}

type testCaseDecode struct {
	description string
	input       []byte
	expectedValue []byte
	expectedError string
}

var decodeCases = []testCaseDecode { {{range .J.decodeMessage}}
	{
		description: 	{{printf "%q" .Description}},
		input: 			{{.FormatInput}},
		expectedValue: 	{{.FormatExpectedValue}},
		expectedError:  {{printf "%q" .FormatExpectedError}},
	},{{end}}
}
`
