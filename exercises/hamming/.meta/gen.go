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
	if err := gen.Gen("hamming", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []oneCase
}

type oneCase struct {
	Description string
	Input       struct {
		Strand1 string
		Strand2 string
	}
	Expected interface{}
}

func (o oneCase) Want() int {
	d, ok := o.Expected.(float64)
	if ok {
		return int(d)
	}
	return -1
}

// template applied to above data structure generates the Go test cases
var tmpl = `package hamming

{{.Header}}

var testCases = []struct {
	s1   string
	s2   string
	want int
}{
{{range .J.Cases}}{ // {{.Description}}
	{{printf "%q" .Input.Strand1}},
	{{printf "%q" .Input.Strand2}},
	{{.Want}},
},
{{end}}}
`
