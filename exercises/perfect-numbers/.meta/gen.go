package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("perfect-numbers", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Cases    []struct {
		Description string
		Cases       []oneCase
	}
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Number int64
	}
	Expected interface{}
}

func (c oneCase) Valid() bool {
	valid, _ := determineExpected(c.Expected)
	return valid
}

func (c oneCase) ExpectedClassification() string {
	_, e := determineExpected(c.Expected)
	switch e {
	case "perfect":
		return "ClassificationPerfect"
	case "abundant":
		return "ClassificationAbundant"
	case "deficient":
		return "ClassificationDeficient"
	}
	return e
}

// determineExpected examines an .Expected interface{} object and determines
// whether a test case is valid(bool) and has a classification or expects an error,
// returning valid and classification.
func determineExpected(expected interface{}) (bool, string) {
	exp, ok := expected.(string)
	if ok {
		return ok, exp
	}
	return false, ""
}

// Template to generate test cases.
var tmpl = `package perfect

{{.Header}}

var classificationTestCases = []struct {
	description	string
	input		int64
	ok		bool
	expected	Classification
}{ {{range .J.Cases}} {{range .Cases}}
{
	description:	"{{.Description}}",
	input:		{{.Input.Number}},
{{if .Valid}} ok: true,
	expected: {{.ExpectedClassification}},
{{- else}} ok: false,
{{- end}}
},{{end}}{{end}}
}
`
