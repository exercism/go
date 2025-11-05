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

type listCase struct {
	Description string `json:"description"`
	Input       struct {
		Score int `json:"score"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `package allergies

{{.Header}}

// allergicTo
type allergicToInput struct {
	allergen string
	score    uint
}

var allergicToTests = []struct {
	description string
	input       allergicToInput
	expected    bool
}{
{{range .J.allergicTo}}{
		description: {{printf "%q"  .Description}},
		input:       allergicToInput{
			allergen: {{printf "%q"  .Input.Item}},
			score:    {{printf "%d"  .Input.Score}},
		},
		expected:    {{printf "%v"  .Expected}},
	},
{{end}}
}

// list
var listTests = []struct {
	description string
	score       uint
	expected    []string
}{
	{{range .J.list}}{
		description: {{printf "%q"  .Description}},
	    score: 	     {{printf "%d"  .Input.Score}},
		expected:    {{printf "%#v"  .Expected}},
	},
	{{end}}
}
`
