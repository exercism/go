package main

import (
	"encoding/json"
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
	if err := gen.Gen("allergies", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Groups   []testGroup `json:"Cases"`
}

type testGroup struct {
	Description     string
	Cases           []json.RawMessage `property:"RAW"`
	AllergicToCases []struct {
		Description string
		Input       struct {
			Item  string
			Score uint
		}
		Expected bool
	} `property:"allergicTo"`
	ListCases []struct {
		Description string
		Input       struct {
			Score uint
		}
		Expected []string
	} `property:"list"`
}

var tmpl = `package allergies

{{.Header}}

type aTest struct {
	description string
	item string
	score uint
	expected bool
}

type lTest struct {
	description string
	score uint
	expected []string
}


func getTests() ([]aTest,[]lTest) {

	var allergicToTests []aTest
	var listTests []lTest

{{range .J.Groups}}
	// {{ .Description }}
	{{- if .AllergicToCases }}
	allergicToTests = append(allergicToTests,
			{{- range .AllergicToCases }}
			aTest{
					description: {{.Description | printf "%q"}},
					item: {{.Input.Item | printf "%q" }},
					score: {{.Input.Score}},
					expected: {{.Expected}},
			},	
			{{- end }}
	){{- end }}
	{{- if .ListCases }}
		listTests = append(listTests,
			{{- range .ListCases }}
				lTest{ {{.Description | printf "%q"}}, {{.Input.Score}}, {{.Expected | printf "%#v"}}},
			{{- end }}
		){{- end }}
{{end}}
return allergicToTests,listTests
}`
