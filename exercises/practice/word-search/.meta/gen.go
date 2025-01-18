package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"search": &[]testCase{},
	}
	if err := gen.Gen("word-search", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Grid             []string `json:"grid"`
		WordsToSearchFor []string `json:"wordsToSearchFor"`
	} `json:"input"`
	Expected map[string]entry `json:"expected"`
}

type entry struct {
	Start point
	End   point
}

type point struct {
	Column position
	Row    position
}

type position int

func (p position) Minus1() position {
	return p - 1
}

func (c testCase) ErrorExpected() bool {
	// When any of the word positions have a null position, expect an error.
	for _, p := range c.Expected {
		if p.Start.Row == 0 && p.Start.Column == 0 && p.End.Row == 0 && p.End.Column == 0 {
			return true
		}
	}
	return false
}

var tmpl = `package wordsearch

{{.Header}}

var testCases = []struct {
	description string
	puzzle      []string
	words       []string
	expectError bool
	expected    map[string][2][2]int
}{ {{range .J.search}}
{
	description: 	{{printf "%q"  .Description}},
	puzzle: 		{{printf "%#v" .Input.Grid}},
	words:			{{printf "%#v"  .Input.WordsToSearchFor}},
	expectError: 	{{printf "%t"  .ErrorExpected}},
	expected:		map[string][2][2]int{ {{ range $key, $value := .Expected }} "{{ $key }}": { { {{ $value.Start.Column.Minus1 }}, {{ $value.Start.Row.Minus1 }}, }, { {{ $value.End.Column.Minus1 }}, {{ $value.End.Row.Minus1 }} } }, {{ end }} },
},{{end}}
}
`
