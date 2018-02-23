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
	if err := gen.Gen("scale-generator", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Comments    []string
		Description string
		Cases       []struct {
			Description string
			Property    string
			Input       struct {
				Tonic     string
				Intervals string
			}
			Expected []string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package scale

{{.Header}}

type scaleTest struct {
	description string
	tonic    string
	interval string
	expected []string
}

var scaleTestCases = []scaleTest {
{{range .J.Cases}} {{range .Cases}}{
	description: {{printf "%q" .Description }},
	tonic: {{printf "%q" .Input.Tonic }},
	interval: {{printf "%q" .Input.Intervals }},
	expected: {{printf "%#v" .Expected }},
},
{{end}}{{end}}
}
`
