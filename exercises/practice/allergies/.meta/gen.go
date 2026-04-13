package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"allergicTo": &[]allergicToCase{},
		"list":       &[]listCase{},
	}
	if err := gen.Gen("allergies", j, t); err != nil {
		log.Fatal(err)
	}
}

type allergicToCase struct {
	Description string `json:"description"`
	Input       struct {
		Item  string `json:"item"`
		Score uint   `json:"score"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

func (tc allergicToCase) ExpandedDescription() string {
	return fmt.Sprintf("%q", tc.Description + ": check " + tc.Input.Item)
}

type listCase struct {
	Description string `json:"description"`
	Input       struct {
		Score int `json:"score"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `{{.Header}}

type allergicToInput struct {
	allergen string
	score    uint
}

type allergicToTestCase struct {
	description string
	input       allergicToInput
	expected    bool
}

var allergicToTests = []allergicToTestCase { {{range .J.allergicTo}}
	{
		description: {{.ExpandedDescription}},
		input:       allergicToInput{
			allergen: {{printf "%q"  .Input.Item}},
			score:    {{printf "%d"  .Input.Score}},
		},
		expected:    {{printf "%v"  .Expected}},
	},{{end}}
}

type listTestCase struct {
	description string
	score       uint
	expected    []string
}
var listTests = []listTestCase { {{range .J.list}}
	{
		description: {{printf "%q"  .Description}},
		score: 	     {{printf "%d"  .Input.Score}},
		expected:    {{printf "%#v"  .Expected}},
	},{{end}}
}
`
