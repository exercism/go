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
		"age": &[]testCase{},
	}
	if err := gen.Gen("space-age", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Planet  string  `json:"planet"`
		Seconds float64 `json:"seconds"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectedValue() float64 {
	v, ok := t.Expected.(float64)
	if ok {
		return v
	}
	return -1
}

// Template to generate test cases.
var tmpl = `package space

{{.Header}}

var testCases = []struct {
	description		string
	planet			Planet
	seconds			float64
	expected		float64
}{ {{range .J.age}}
{
	description:	{{printf "%q"    .Description}},
	planet:			{{printf "%q"   .Input.Planet}},
	seconds:		{{printf "%.1f"  .Input.Seconds}},
	expected:		{{printf "%.2f"  .ExpectedValue}},
},{{end}}
}
`
