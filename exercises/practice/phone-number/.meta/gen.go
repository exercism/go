package main

import (
	"fmt"
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
		"clean": &[]testCase{},
	}
	if err := gen.Gen("phone-number", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Phrase string `json:"phrase"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectError() bool {
	v, ok := t.Expected.(map[string]interface{})
	if ok {
		_, gotError := v["error"]
		return gotError
	}
	return false
}

func (t testCase) ExpectedNumber() string {
	v, ok := t.Expected.(string)
	if ok {
		return v
	}
	return ""
}

func (t testCase) AreaCode() string {
	if expectedNumber := t.ExpectedNumber(); expectedNumber != "" {
		return expectedNumber[:3]
	}
	return ""
}

func (t testCase) Formatted() string {
	expectedNumber := t.ExpectedNumber()
	if expectedNumber == "" {
		return ""
	}
	return fmt.Sprintf("(%s) %s-%s", t.AreaCode(), expectedNumber[3:6], expectedNumber[6:10])
}

var tmpl = `package phonenumber

{{.Header}}

type testCase struct {
	description       string
	input             string
	expectErr         bool
	expectedNumber    string
	expectedAreaCode  string
	expectedFormatted string
}

var testCases = []testCase{
{{range .J.clean}}{
		description:      {{printf "%q"  .Description}},
		input:            {{printf "%q"  .Input.Phrase}},
		expectErr:        {{printf "%t"  .ExpectError}},
		expectedNumber:   {{printf "%q"  .ExpectedNumber}},
		expectedAreaCode:         {{printf "%q"  .AreaCode}},
		expectedFormatted:        {{printf "%q"  .Formatted}},
},
{{end}}
}
`
