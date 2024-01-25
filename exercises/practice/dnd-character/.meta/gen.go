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
		"ability":   &[]abilityTestCase{},
		"character": &[]characterTestCase{},
		"strength":  &[]strengthTestCase{},
	}
	if err := gen.Gen("dnd-character", j, t); err != nil {
		log.Fatal(err)
	}
}

type modifierTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Score int `json:"score"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type abilityTestCase struct {
	Description string   `json:"description"`
	Input       struct{} `json:"input"`
	Expected    string   `json:"expected"`
}

type characterTestCase struct {
	Description string   `json:"description"`
	Input       struct{} `json:"input"`
	Expected    struct {
		Strength     string `json:"strength"`
		Dexterity    string `json:"dexterity"`
		Constitution string `json:"constitution"`
		Wisdom       string `json:"wisdom"`
		Intelligence string `json:"intelligence"`
		Charisma     string `json:"charisma"`
		Hitpoints    string `json:"hitpoints"`
	} `json:"expected"`
}

type strengthTestCase struct {
	Description string   `json:"description"`
	Input       struct{} `json:"input"`
	Expected    string   `json:"expected"`
}

var tmpl = `package dndcharacter

{{.Header}}

var modifierTests = []struct {
	description    string
	score          int 
	expected       int
}{
	{{range .J.modifier}} 
		{
			description: {{printf "%q"  .Description}},
			score: {{printf "%d" .Input.Score}},
			expected: {{printf "%d"  .Expected}},
		},
	{{end}}
}

var abilityTests = []struct {
	description string
	input struct{}
	expected string
}{
	{{range .J.ability}} 
		{
			description: {{printf "%q"  .Description}},
			input: {{printf "%#v"  .Input}},
			expected: {{printf "%q"  .Expected}},
		},
	{{end}}
}

var characterTests = []struct {
	description string
	input struct{}
	expected struct{
		Strength string
		Dexterity string
		Constitution string
		Intelligence string
		Wisdom string
		Charisma string
		Hitpoints string
	}
}{
	{{range .J.character}} 
		{
			description: {{printf "%q" .Description}},
			input: {{printf "%#v" .Input}},
			expected: struct{
				Strength string
				Dexterity string
				Constitution string
				Intelligence string
				Wisdom string
				Charisma string
				Hitpoints string
			}{
				Strength: {{ printf "%q" .Expected.Strength }},
				Dexterity: {{ printf "%q" .Expected.Dexterity }},
				Constitution: {{ printf "%q" .Expected.Constitution }},
				Intelligence: {{ printf "%q" .Expected.Intelligence }},
				Wisdom: {{ printf "%q" .Expected.Wisdom }},
				Charisma: {{ printf "%q" .Expected.Charisma }},
				Hitpoints: {{ printf "%q" .Expected.Hitpoints }},
			},
		},
	{{end}}
}

var strengthTests = []struct {
	description string
	input struct{}
	expected string
}{
	{{range .J.strength}} 
	{
		description: {{printf "%q"  .Description}},
		input: {{printf "%#v"  .Input}},
		expected: {{printf "%q"  .Expected}},
	},
	{{end}}
}
`
