// +build ignore

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
	if err := gen.Gen("largest-series-product", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Digits      string
		Span        int
		Expected    int64
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package lsproduct

{{.Header}}

var tests = []struct {
	digits  string
	span    int
	product int64
	ok      bool
}{
{{range .J.Cases}}{ "{{.Digits}}", {{.Span}}, {{.Expected}}, {{ge .Expected 0}}},
{{end}}}
`
