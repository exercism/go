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
	if err := gen.Gen("sum-of-multiples", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Factors []int
			Limit   int
		}
		Expected int
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package summultiples

{{.Header}}

var varTests = []struct {
	divisors     []int
	limit        int
	sum          int
}{
{{range .J.Cases}}{ {{.Input.Factors | printf "%#v"}}, {{.Input.Limit}}, {{.Expected}}}, // {{.Description}}
{{end}}}
`
