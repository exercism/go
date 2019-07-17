package main

import (
	"encoding/json"
	"log"
	"strings"
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
		Description string
		Input       struct {
			Hour, Minute int
		}
		Expected string
	} `property:"create"`
	AddCases []struct {
		Description string
		Input       struct {
			Hour, Minute, Value int
		}
		Expected string
	} `property:"add"`
	SubtractCases []struct {
		Description string
		Input       struct {
			Hour, Minute, Value int
		}
		Expected string
	} `property:"subtract"`
	EqCases []struct {
		Description string
		Input       struct {
			Clock1, Clock2 struct{ Hour, Minute int }
		}
		Expected bool
	} `property:"equal"`
}

func (groups testGroup) GroupShortName() string {
	return strings.ToLower(strings.Fields(groups.Description)[0])
}

var tmpl = `package clock

{{.Header}}

{{range .J.Groups}}
	// {{ .Description }}

{{- if .CreateCases }}
	var timeTests = []struct {
			n string
			h, m int
			want string
		}{
			{{- range .CreateCases }}
				{
				  {{ printf "%q" .Description }},
				  {{.Input.Hour}}, {{.Input.Minute}}, {{.Expected | printf "%#v"}},
				},
			{{- end }}
		}
	{{- end }}

	{{- if .AddCases }}
		var {{ .GroupShortName }}Tests = []struct {
			n string
			h, m, a int
			want string
		}{
			{{- range .AddCases }}
				{
				  {{ printf "%q" .Description}},
				  {{.Input.Hour}}, {{.Input.Minute}}, {{.Input.Value}}, {{.Expected | printf "%#v"}},
				 },
			{{- end }}
		}
	{{- end }}

	{{- if .SubtractCases }}
		var {{ .GroupShortName }}Tests = []struct {
			n string
			h, m, a int
			want string
		}{
			{{- range .SubtractCases }}
			    {
				  {{ printf "%q" .Description}},
				  {{.Input.Hour}}, {{.Input.Minute}}, {{.Input.Value}}, {{.Expected | printf "%#v"}},
				},
			{{- end }}
		}
	{{- end }}

	{{- if .EqCases }}
		type hm struct{ h, m int }
		var eqTests = []struct {
			n string
			c1, c2 hm
			want   bool
		}{
			{{- range .EqCases }}
				{
				  {{ printf "%q" .Description}},
					hm{ {{.Input.Clock1.Hour}}, {{.Input.Clock1.Minute}}},
					hm{ {{.Input.Clock2.Hour}}, {{.Input.Clock2.Minute}}},
					{{.Expected}},
				},
			{{- end }}
		}
	{{- end }}
{{end}}
`
