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
	if err := gen.Gen("nth-prime", &j, t); err != nil {
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

func (c OneCase) HasPrimeAnswer() bool {
	hasPrimeAnswer, _ := determineExpected(c.Expected)
	return hasPrimeAnswer
}

func (c OneCase) PrimeAnswer() int {
	_, answer := determineExpected(c.Expected)
	return answer
}

// determineExpected examines an .Expected interface{} object and determines
// whether a test case has a Prime answer or expects an error,
// returning true and the answer, or false and zero.
func determineExpected(expected interface{}) (bool, int) {
	value, ok := expected.(float64)
	if ok {
		return true, int(value)
	}
	return false, 0
}

// template applied to above data structure generates the Go test cases
var tmpl = `package prime

{{.Header}}

var tests = []struct {
	description string
	n   int
	p   int
	ok  bool
}{
{{range .J.Cases}}{
	"{{.Description}}",
	{{.Input.Number}},
{{- if .HasPrimeAnswer}}
	{{.PrimeAnswer}},
	true,
{{- else}}
	0,
	false,
{{- end}}
},
{{end}}}
`
