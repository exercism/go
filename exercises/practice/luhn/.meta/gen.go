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
	if err := gen.Gen("luhn", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Value string
		}
		Expected bool
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package luhn

{{.Header}}

var testCases = []struct {
	description string
	input       string
	ok			bool
}{
{{range .J.Cases}}{
	{{printf "%q" .Description}},
	{{printf "%q" .Input.Value}},
	{{.Expected}},
},
{{end}}
}`
