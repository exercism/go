package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"reflect"
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
	if reflect.ValueOf(tc.Expected).Kind() != reflect.Slice {
		return "nil"
	}
	rawValues := reflect.ValueOf(tc.Expected).Interface().([]any)
	var stringValues []string
	for _, raw := range rawValues {
		stringValues = append(stringValues, raw.(string))
	}
	return fmt.Sprintf("%#v", stringValues)
}

func (tc testCase) ExpectedError() string {
	if reflect.ValueOf(tc.Expected).Kind() != reflect.Map {
		return "nil"
	}
	return "ErrInvalidBase"
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description    string
	input          string
	expected       []string
	expectedError  error
}{
{{range .J.proteins}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{printf "%q" .Input.Strand}},
	expected:    	{{.ExpectedProteins}},
	expectedError:  {{.ExpectedError}},
},
{{end}}
}`
