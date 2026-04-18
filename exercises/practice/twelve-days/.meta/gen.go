package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"recite": &[]testCase{},
	}
	if err := gen.Gen("twelve-days", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		StartVerse int `json:"startVerse"`
		EndVerse   int `json:"endVerse"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	verse       int
	expected    string
}

var testCases = []testCase { {{range .J.recite}} {{ if eq .Input.StartVerse .Input.EndVerse }}
	{
		description: {{printf "%q" .Description}},
		verse:       {{.Input.StartVerse}},
		expected:    {{index .Expected 0 | printf "%q"}},
	},{{end}}{{end}}
}

var song = {{range .J.recite}}{{if eq .Input.StartVerse 1}}{{ if eq .Input.EndVerse 12 }}[]string{ {{range .Expected}}
	{{printf "%q" .}},{{end}}
}{{end}}{{end}}{{end}}
`
