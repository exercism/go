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
		"evaluate": &[]testCase{},
		// TODO: add test with property `evaluateBoth` to forth_test.go and generate cases in cases_test.go
	}
	if err := gen.Gen("forth", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Instructions []string `json:"instructions"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectedNumbers() []int {
	numbers, ok := t.Expected.([]interface{})
	if !ok {
		return nil
	}
	result := make([]int, 0)
	for _, number := range numbers {
		x, ok := number.(float64)
		if !ok {
			return nil
		}
		result = append(result, int(x))
	}
	return result
}

func (t testCase) ExplainText() string {
	if t.ExpectedNumbers() != nil {
		return ""
	}
	m, ok := t.Expected.(map[string]interface{})
	if !ok {
		return ""
	}
	errText, ok := m["error"].(string)
	if !ok {
		return ""
	}
	return errText
}

// template applied to above data structure generates the Go test cases
var tmpl = `package forth

{{.Header}}

var testCases = []struct {
	description    string
	input   	   []string
	expected 	   []int // nil slice indicates error expected.
	explainText    string   // error explanation text
}{ {{range .J.evaluate}}
{
	description: {{printf "%q"  .Description}},
	input: {{printf "%#v" .Input.Instructions}},
	expected: {{printf "%#v" .ExpectedNumbers}},
	explainText: {{printf "%q"  .ExplainText}},
},{{end}}
}
`
