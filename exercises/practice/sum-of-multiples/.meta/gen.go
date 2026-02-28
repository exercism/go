package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"sum": &[]testCase{},
	}
	if err := gen.Gen("sum-of-multiples", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Factors []int `json:"factors"`
		Limit   int   `json:"limit"`
	} `json:"input"`
	Expected int `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	divisors     []int
	limit        int
	expected          int
}{
{{range .J.sum}}{ 
	description: 	{{printf "%q"  .Description}},
	divisors: 		{{printf "%#v" .Input.Factors}}, 
	limit: 			{{printf "%d"  .Input.Limit}}, 
	expected: 		{{printf "%d"  .Expected}},
},
{{end}}
}
`
