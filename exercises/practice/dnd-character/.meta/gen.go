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

// The problem specifications expect properties "Ability" and "Character" to be present
// but, we don't want to generate test cases for these as they don't have any input. To
// circumvent the test generator failing on missing properties, we've introduced
// the emptyTestCase struct.
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
