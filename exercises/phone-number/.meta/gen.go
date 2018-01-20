package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"areacode":  areacode,
		"expectErr": expectErr,
		"formatted": formatted,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("phone-number", &j, t); err != nil {
		log.Fatal(err)
	}
}

func expectErr(expected string) bool {
	return len(expected) == 0
}

func areacode(expected string) string {
	if expectErr(expected) {
		return ""
	}
	return expected[0:3]
}

func formatted(expected string) string {
	if expectErr(expected) {
		return ""
	}
	return "(" + areacode(expected) + ") " + expected[3:6] + "-" + expected[6:10]
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []struct {
			Description string
			Input       struct {
				Phrase string
			}
			Expected string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package phonenumber

{{.Header}}

{{range .J.Cases}}// {{.Description}}
var numberTests = []struct {
	description string
	input       string
	expectErr   bool
	number      string
	areaCode    string
	formatted   string
}{
	{{range .Cases}}{
		description: {{printf "%q" .Description}},
		input: {{printf "%q" .Input.Phrase}},
		{{if expectErr .Expected}} expectErr: {{expectErr .Expected | printf "%v" }},
		{{else}} number: {{printf "%q" .Expected}},
		areaCode: {{areacode .Expected  | printf "%q" }},
		formatted: {{formatted .Expected  | printf "%q" }},{{- end}}
},
{{end}}{{end}}
}
`
