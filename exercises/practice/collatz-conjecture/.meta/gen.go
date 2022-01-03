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
	var j js
	if err := gen.Gen("collatz-conjecture", &j, t); err != nil {
		log.Fatal(err)
	}
}

type OneCase struct {
	Description string
	Input       struct {
		Number int
	}
	Expected interface{}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []OneCase
}

func (c OneCase) Valid() bool {
	valid, _ := determineExpected(c.Expected)
	return valid
}

func (c OneCase) Answer() int {
	_, answer := determineExpected(c.Expected)
	return answer
}

// determineExpected examines an .Expected interface{} object and determines
// whether a test case is valid(bool) and has an answer or expects an error.
// returning valid and answer.
func determineExpected(expected interface{}) (bool, int) {
	ans, ok := expected.(float64)
	if ok {
		return ok, int(ans)
	}
	return false, 0
}

// template applied to above data structure generates the Go test cases
var tmpl = `package collatzconjecture

{{.Header}}

var testCases = []struct {
	description string
	input       int
	expectError bool
	expected    int
}{
{{range .J.Cases}}{
	description:	"{{.Description}}",
	input:		{{.Input.Number}},
{{if .Valid}} expected:	{{.Answer}},
{{- else}} expectError: true,
{{- end}}
},
{{end}}}
`
