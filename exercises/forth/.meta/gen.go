// +build ignore

package main

import (
	"log"
	"text/template"

	"../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("forth", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	ParsingAndNumbers  section `json:"parsing and numbers"`
	Addition           section
	Subtraction        section
	Multiplication     section
	Division           section
	CombinedArithmetic section `json:"combined arithmetic"`
	Dup                section
	Drop               section
	Swap               section
	Over               section
	UserDefinedWords   section `json:"user-defined words"`
}

// A grouped section of test cases.
type section struct {
	Cases []struct {
		Description string
		Input       []string
		Expected    []int
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package forth

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

{{/* template for repeated caseSection */}}

{{define "caseSection"}} = []testCase{
{{range $y, $c := .}}{
	{{printf "%q" $c.Description}},
	{{printf "%#v" $c.Input}},
	{{printf "%#v" $c.Expected}},
},
{{end}}}{{end}}

var parsingGroup{{template "caseSection" .J.ParsingAndNumbers.Cases}}

var additionGroup{{template "caseSection" .J.Addition.Cases}}

var subtractionGroup{{template "caseSection" .J.Subtraction.Cases}}

var multiplicationGroup{{template "caseSection" .J.Multiplication.Cases}}

var divisionGroup{{template "caseSection" .J.Division.Cases}}

var arithmeticGroup{{template "caseSection" .J.CombinedArithmetic.Cases}}

var dupGroup{{template "caseSection" .J.Dup.Cases}}

var dropGroup{{template "caseSection" .J.Drop.Cases}}

var swapGroup{{template "caseSection" .J.Swap.Cases}}

var overGroup{{template "caseSection" .J.Over.Cases}}

var userdefinedGroup{{template "caseSection" .J.UserDefinedWords.Cases}}

var testSections = []testcaseSection{
	{"parsing", parsingGroup},
	{"addition(+)", additionGroup},
	{"subtraction(-)", subtractionGroup},
	{"multiplication(*)", multiplicationGroup},
	{"division(/)", divisionGroup},
	{"arithmetic", arithmeticGroup},
	{"dup", dupGroup},
	{"drop", dropGroup},
	{"swap", swapGroup},
	{"over", overGroup},
	{"user-defined", userdefinedGroup},
}
`
