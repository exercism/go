package main

import (
	"../../../../gen"
	"log"
	"strings"
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
	if err := gen.Gen("house", j, t); err != nil {
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

// wrapVerse inserts newlines before the "that" section separators.
func wrapVerse(verse string) string {
	var sb strings.Builder
	for part := range strings.SplitSeq(verse, " that ") {
		if strings.HasPrefix(part, "Jack") || strings.HasPrefix(part, "crowed") {
			sb.WriteString(" that ")
		} else if ! strings.HasPrefix(part, "This") {
			sb.WriteString("\nthat ")
		}
		sb.WriteString(part)
	}
	return sb.String()
}

func (tc testCase) WrappedVerses() []string {
	out := make([]string, len(tc.Expected))
	for i, verse := range tc.Expected {
		out[i] = wrapVerse(verse)
	}
	return out
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
		expected:    {{index .WrappedVerses 0 | printf "%q"}},
	},{{end}}{{end}}
}

var song = {{range .J.recite}}{{if eq .Input.StartVerse 1}}{{ if eq .Input.EndVerse 12 }}[]string{ {{range .WrappedVerses}}
	{{printf "%q" .}},{{end}}
}{{end}}{{end}}{{end}}
`
