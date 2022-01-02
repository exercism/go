package main

import (
	"encoding/json"
	"errors"
	"log"
	"text/template"

	"../../../../gen"
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

func expectErr(expected ExpectedType) bool {
	return len(expected.Error) > 0
}

func areacode(expected ExpectedType) string {
	if expectErr(expected) {
		return ""
	}
	return expected.Value[0:3]
}

func formatted(expected ExpectedType) string {
	if expectErr(expected) {
		return ""
	}
	return "(" + areacode(expected) + ") " + expected.Value[3:6] + "-" + expected.Value[6:10]
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
			Expected ExpectedType
		}
	}
}

type ExpectedType struct {
	Value string
	Error string
}

// UnmarshalJSON for custom type (allow for having multiple types in the same JSON field)
func (ex *ExpectedType) UnmarshalJSON(b []byte) error {
	switch b[0] {
	case '"': // a plain string
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		ex.Value = s
		ex.Error = ""
		return nil
	case '{': // a map
		var s map[string]string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		ex.Error = s["error"]
		ex.Value = ""
		return nil
	}
	return errors.New("Expected type not recognized")
}

// template applied to above data structure generates the Go test cases
var tmpl = `package phonenumber

{{.Header}}

{{range .J.Cases}}// {{.Description}}
var numberTests = []struct {
	description 		string
	input       		string
	expectErr   		bool
	errorDescription	string
	number      		string
	areaCode    		string
	formatted   		string
}{
	{{range .Cases}}{
		description: {{printf "%q" .Description}},
		input: {{printf "%q" .Input.Phrase}},
		{{if expectErr .Expected}} expectErr: {{expectErr .Expected | printf "%v" }},
		errorDescription: {{printf "%q" .Expected.Error}},
		{{else}} number: {{printf "%q" .Expected.Value}},
		areaCode: {{areacode .Expected | printf "%q" }},
		formatted: {{formatted .Expected | printf "%q" }},{{- end}}
},
{{end}}{{end}}
}
`
