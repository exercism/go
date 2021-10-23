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
	if err := gen.Gen("prime-factors", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Groups []TestGroup `json:"cases"`
}

type TestGroup struct {
	Description string
	Cases       []OneCase
}

type OneCase struct {
	Description string
	Input       struct {
		Value int64
	}
	Expected []int64
}

// template applied to above data structure generates the Go test cases
var tmpl = `package prime

{{.Header}}

var tests = []struct {
	description string
	input int64
	expected []int64
}{
{{range .J.Groups}}
{{range .Cases}} {
	"{{.Description}}",
	{{.Input.Value}},
	{{printf "%#v" .Expected}},
},
{{end}}{{end}}}
`
