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
		"sharedBirthday":                       &[]sharedTestCase{},
		"estimatedProbabilityOfSharedBirthday": &[]probabilityTestCase{},
		// These tests are hand written.
		"randomBirthdates": &[]interface{}{},
	}
	if err := gen.Gen("baffling-birthdays", j, t); err != nil {
		log.Fatal(err)
	}
}

type sharedTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Birthdates []string `json:"birthdates"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

type probabilityTestCase struct {
	Description string `json:"description"`
	Input       struct {
		GroupSize int `json:"groupSize"`
	} `json:"input"`
	Expected float64 `json:"expected"`
}

var tmpl = `{{.Header}}

var sharedTestCase = []struct {
	description string
	input       []string
	expected    bool
}{ {{range .J.sharedBirthday}}
	{
		description: {{printf "%q" .Description}},
		input:       {{printf "%#v" .Input.Birthdates}},
		expected:    {{printf "%t" .Expected}},
	},{{end}}
}

var probabilityTestCase = []struct {
	description string
	input       int
	expected    float64
}{ {{range .J.estimatedProbabilityOfSharedBirthday}}
	{
		description: {{printf "%q" .Description}},
		input:       {{.Input.GroupSize}},
		expected:    {{.Expected}},
	},{{end}}
}`
