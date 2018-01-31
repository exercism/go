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
	if err := gen.Gen("rail-fence-cipher", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Groups []testGroup `json:"cases"`
}

type testGroup struct {
	Description string
	Cases       []json.RawMessage `property:"RAW"`
	EncodeCases []oneCase         `property:"encode"`
	DecodeCases []oneCase         `property:"decode"`
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	Input       struct {
		Message string `json:"msg"`
		Rails   int
	}
	Expected string
}

// Template to generate test cases.
var tmpl = `package railfence

{{.Header}}

type testCase struct {
	description string
	message     string
	rails       int
	expected    string
}

{{range .J.Groups}}
	// {{ .Description }}
	{{- if .EncodeCases }}
		var encodeTests = []testCase{
			{{- range .EncodeCases }}

				{ {{.Description | printf "%q"}},
				{{.Input.Message | printf "%q"}},
				{{.Input.Rails}},
				{{.Expected | printf "%q"}}},
			{{- end }}
		}
	{{- end }}

	{{- if .DecodeCases }}
		var decodeTests = []testCase{
			{{- range .DecodeCases }}

				{ {{.Description | printf "%q"}},
				{{.Input.Message | printf "%q"}},
				{{.Input.Rails}},
				{{.Expected | printf "%q"}}},
			{{- end }}
		}
	{{- end }}

{{end}}
`
