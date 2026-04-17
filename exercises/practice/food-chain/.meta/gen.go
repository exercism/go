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
	if err := gen.Gen("food-chain", j, t); err != nil {
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

type verseTestCase struct {
	description string
	verse       int
	expected    []string
}

var verseTestCases = []verseTestCase { {{range .J.recite}}{{ if eq .Input.StartVerse .Input.EndVerse }}
	{
		description: {{printf "%q" .Description}},
		verse:       {{.Input.StartVerse}},
		expected:    []string{ {{range .Expected}}
			{{printf "%q" .}},{{end}}
		},
	},{{end}}{{end}}
}

type multiVerseTestCase struct {
	description string
	start       int
	end         int
	expected    []string
}

var multiVerseTestCases = []multiVerseTestCase { {{range .J.recite}}{{ if ne .Input.StartVerse .Input.EndVerse }}
	{
		description: {{printf "%q" .Description}},
		start:       {{.Input.StartVerse}},
		end:         {{.Input.EndVerse}},
		expected:    []string{ {{range .Expected}}
			{{printf "%q" .}},{{end}}
		},
	},{{end}}{{end}}
}

var song = {{range .J.recite}}{{if eq .Input.StartVerse 1}}{{ if eq .Input.EndVerse 8 }}[]string{ {{range .Expected}}
	{{printf "%q" .}},{{end}}
}{{end}}{{end}}{{end}}
`
