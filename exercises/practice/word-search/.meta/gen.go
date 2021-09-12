package main

import (
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
	if err := gen.Gen("word-search", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Comments []string
	Cases    []OneCase
}

// template applied to above data structure generates the Go test cases

type Int int

type Coordinates struct {
	Column Int
	Row    Int
}
type Position struct {
	Start Coordinates
	End   Coordinates
}

type OneCase struct {
	Description string
	Property    string
	Input       struct {
		Grid             []string
		WordsToSearchFor []string
	}
	Expected map[string]Position
}

func (p Position) NullPosition() bool {
	if p.Start.Column == 0 && p.Start.Row == 0 &&
		p.End.Column == 0 && p.End.Row == 0 {
		return true
	}
	return false
}

func (c OneCase) ErrorExpected() bool {
	// When any of the word positions have an null position, expect an error.
	for _, p := range c.Expected {
		if p.NullPosition() {
			return true
		}
	}
	return false
}

func (v Int) Minus1() int {
	return int(v) - 1
}

// Template to generate test cases

var tmpl = `package wordsearch

{{.Header}}

var testCases = []struct {
	description    string
	puzzle  []string	// puzzle strings
	words   []string        // words to search for
	expected map[string][2][2]int // expected coordinates
	expectError bool
}{ {{range .J.Cases}}
{
	{{printf "%q"  .Description}},
	{{printf "%#v" .Input.Grid}},
	{{printf "%#v"  .Input.WordsToSearchFor}},
	{{if .ErrorExpected}} map[string][2][2]int{ },
	true,
	{{else}} map[string][2][2]int{ {{ range $key, $value := .Expected }} "{{ $key }}": { { {{ $value.Start.Column.Minus1 }}, {{ $value.Start.Row.Minus1 }}, }, { {{ $value.End.Column.Minus1 }}, {{ $value.End.Row.Minus1 }} } }, {{ end }} },
	false,
	{{- end}}
},{{end}}
}
`
