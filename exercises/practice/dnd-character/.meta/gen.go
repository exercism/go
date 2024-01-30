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
	var j = map[string]interface{}{
		"modifier":  &[]modifierTestCase{},
		"ability":   &[]emptyTestCase{},
		"character": &[]emptyTestCase{},
	}
	if err := gen.Gen("dnd-character", j, t); err != nil {
		log.Fatal(err)
	}
}

// Problem specifications have the 'ability' and 'character' properties,
// and the test generator expects them to be present.
// However, for these properties, generating test cases automatically
// from the 'input' and the 'expected' objects is not trivial.
// To satisfy the test generator, we create an emptyTestCase for these properties
// and implement the tests manually.
// In the future we might adapt the test generator to also cover this exercise.
type emptyTestCase struct{}

type modifierTestInput struct {
	Score int `json:"score"`
}

type modifierTestCase struct {
	Description string            `json:"description"`
	Input       modifierTestInput `json:"input"`
	Expected    int               `json:"expected"`
}

var tmpl = `
package dndcharacter

{{.Header}}

type modifierTestInput struct {
	Score int
}

var modifierTests = []struct {
	description    string
	input          modifierTestInput
	expected       int
}{
	{{range .J.modifier}} 
		{
			description: {{printf "%q"  .Description}},
			input: modifierTestInput { 
				Score: {{printf "%d" .Input.Score}},
			},
			expected: {{printf "%d"  .Expected}},
		},
	{{end}}
}
`
