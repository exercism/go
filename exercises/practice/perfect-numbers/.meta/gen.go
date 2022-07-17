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
	var j = map[string]interface{}{
		"classify": &[]testCase{},
	}
	if err := gen.Gen("perfect-numbers", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Number int `json:"number"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) Valid() bool {
	_, ok := t.Expected.(string)
	return ok
}

func (t testCase) ExpectedClassification() string {
	e, ok := t.Expected.(string)
	if ok {
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
	return ""
}

// Template to generate test cases.
var tmpl = `package perfect

{{.Header}}

var classificationTestCases = []struct {
	description	string
	input		int64
	ok		bool
	expected	Classification
}{
{{range .J.classify}}{
	description:	{{printf "%q" .Description}},
	input:			{{printf "%d" .Input.Number}},
	ok: 			{{printf "%t" .Valid}},
	{{if .Valid}} expected: 		{{printf "%s" .ExpectedClassification}},{{- end}}
},
{{end}}
}
`
