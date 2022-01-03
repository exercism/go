package main

import (
	"encoding/json"
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
	if err := gen.Gen("run-length-encoding", &j, t); err != nil {
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
	EncodeCases []struct {
		Description string
		Input       struct {
			String string
		}
		Expected string
	} `property:"encode"`
	DecodeCases []struct {
		Description string
		Input       struct {
			String string
		}
		Expected string
	} `property:"decode"`
	EncodeDecodeCases []struct {
		Description string
		Input       struct {
			String string
		}
		Expected string
	} `property:"consistency"`
}

var tmpl = `package encode

{{.Header}}

{{range .J.Groups}}
	// {{ .Description }}

	{{- if .EncodeCases }}
		var encodeTests = []struct {
			input string
			expected string
			description string
		}{
			{{- range .EncodeCases }}
				{ {{.Input.String | printf "%q"}}, {{.Expected | printf "%q"}}, {{.Description | printf "%q"}} },
			{{- end }}
		}
	{{- end }}

	{{- if .DecodeCases }}
		var decodeTests = []struct {
			input string
			expected string
			description string
		}{
			{{- range .DecodeCases }}
				{ {{.Input.String | printf "%q"}}, {{.Expected | printf "%q"}}, {{.Description | printf "%q"}} },
			{{- end }}
		}
	{{- end }}

	{{- if .EncodeDecodeCases }}
		var encodeDecodeTests = []struct {
			input string
			expected string
			description string
		}{
			{{- range .EncodeDecodeCases }}
				{ {{.Input.String | printf "%q"}}, {{.Expected | printf "%q"}}, {{.Description | printf "%q"}} },
			{{- end }}
		}
	{{- end }}
{{end}}
`
