package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"isEncode": isEncode,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("atbash-cipher", &j, t); err != nil {
		log.Fatal(err)
	}
}

func isEncode(property string) bool {
	return property == "encode"
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []struct {
			Description string
			Property    string
			Phrase      string
			Expected    string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package atbash

{{.Header}}

type atbashTest struct {
	s        string
	expected string
}

var tests = []atbashTest {
{{range .J.Cases}} {{range .Cases}}{
	{{if isEncode .Property}} s: {{printf "%q" .Phrase }},
		expected: {{printf "%q" .Expected }}, {{else}} s: {{printf "%q" .Expected }},
		expected:  {{printf "%q" .Phrase }}, {{- end}}
},
{{end}}{{end}}
}
`
