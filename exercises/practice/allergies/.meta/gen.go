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
			Score uint
		}
		Expected []struct {
			Substance string
			Result    bool
		}
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

{{range .J.Groups}}
	// {{ .Description }}

	{{- if .AllergicToCases }}
		type allergicResult struct{
			substance string
			result bool
		}
		var allergicToTests = []struct {
			description string
			score uint
			expected []allergicResult
		}{
			{{- range .AllergicToCases }}
				{
					description: {{.Description | printf "%q"}},
					score: {{.Input.Score}},
					expected: []allergicResult{ {{range .Expected}}
					  { {{.Substance | printf "%q"}}, {{.Result}} },{{end}}
					},
				},
			{{- end }}
		}
	{{- end }}

	{{- if .ListCases }}
		var listTests = []struct {
			description string
			score uint
			expected []string
		}{
			{{- range .ListCases }}
				{ {{.Description | printf "%q"}}, {{.Input.Score}}, {{.Expected | printf "%#v"}}},
			{{- end }}
		}
	{{- end }}
{{end}}
`
