package main

import (
	"log"
	"strings"
	"text/template"
	"unicode"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"gamestate": &[]testCase{},
	}
	if err := gen.Gen("state-of-tic-tac-toe", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Board []string `json:"board"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (o testCase) ExpectedResult() string {
	s, ok := o.Expected.(string)
	if !ok {
		return "\"\""
	}
	return Title(s)
}

func (o testCase) ExpectError() bool {
	_, ok := o.Expected.(map[string]interface{})
	return ok
}

// template applied to above data structure generates the Go test cases
var tmpl = `package stateoftictactoe

{{.Header}}

var testCases = []struct{
	description string
	board []string
	expected State
	wantErr bool
} {
{{range .J.gamestate}} {
	description: 	{{printf "%q" 	.Description}},
	board: 			{{printf "%#v" 	.Input.Board }},
	expected: 		{{printf "%s" 	.ExpectedResult }},
	wantErr:  		{{printf "%t" 	.ExpectError }},
},
{{end}}
}
`

// Title is a copy of strings.Title function of the stdlib.
// The copy is here because strings.Title is deprecated but we still
// want to use this function as the alternative would require us to support
// external dependencies which we don't yet (tracking issue https://github.com/exercism/go/issues/2379).
// Students should still be able to use strings.Title if they want.
// Since this exercise is currently deprecated, this shouldn't matter too much.
func Title(s string) string {
	// Use a closure here to remember state.
	// Hackish but effective. Depends on Map scanning in order and calling
	// the closure once per rune.
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

// Copy of strings.isSeparator function of the stdlib.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
