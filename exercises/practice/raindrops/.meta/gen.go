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
	if err := gen.Gen("raindrops", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Input struct {
			Number int
		}
		Expected string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package raindrops

{{.Header}}

var tests = []struct {
	input    int
	expected string
}{
{{range .J.Cases}}{ {{.Input.Number}}, "{{.Expected}}"},
{{end}}}
`
