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
	if err := gen.Gen("clock", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to umarshal into
type js struct {
	Groups []testGroup `json:"Cases"`
}

type testGroup struct {
	Description string
	Cases       []json.RawMessage `property:"RAW"`
	CreateCases []struct {
		Description  string
		Hour, Minute int
		Expected     string
	} `property:"create"`
	AddCases []struct {
		Description       string
		Hour, Minute, Add int
		Expected          string
	} `property:"add"`
	EqCases []struct {
		Description    string
		Clock1, Clock2 struct{ Hour, Minute int }
		Expected       bool
	} `property:"equal"`
}

var tmpl = `package clock

{{.Header}}

{{range .J.Groups}}
	// {{ .Description }}

	{{- if .CreateCases }}
		var timeTests = []struct {
			h, m int
			want string
		}{
			{{- range .CreateCases }}
				{ {{.Hour}}, {{.Minute}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
			{{- end }}
		}
	{{- end }}

	{{- if .AddCases }}
		var addTests = []struct {
			h, m, a int
			want string
		}{
			{{- range .AddCases }}
				{ {{.Hour}}, {{.Minute}}, {{.Add}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
			{{- end }}
		}
	{{- end }}

	{{- if .EqCases }}
		type hm struct{ h, m int }
		var eqTests = []struct {
			c1, c2 hm
			want   bool
		}{
			{{- range .EqCases }}
				// {{.Description}}
				{
					hm{ {{.Clock1.Hour}}, {{.Clock1.Minute}}},
					hm{ {{.Clock2.Hour}}, {{.Clock2.Minute}}},
					{{.Expected}},
				},
			{{- end }}
		}
	{{- end }}
{{end}}
`
