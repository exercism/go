package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

type item struct {
	Weight int `json:"weight"`
	Value  int `json:"value"`
}
type maximumValueCaseInput struct {
	MaximumWeight int    `json:"maximumWeight"`
	Items         []item `json:"items"`
}

type maximumValueCase struct {
	Description string                `json:"description"`
	Input       maximumValueCaseInput `json:"input"`
	Expected    int                   `json:"expected"`
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
		"maximumValue": &[]maximumValueCase{},
	}
	if err := gen.Gen("knapsack", j, t); err != nil {
		log.Fatal(err)
	}
}

var tmpl = `package knapsack

{{.Header}}

type maximumValueCaseInput struct {
	MaximumWeight int
	Items         []Item
}

type maximumValueTest struct {
	description string       
	input       maximumValueCaseInput
	expected    int
}	

var maximumValueTests = []maximumValueTest { 
	{{range .J.maximumValue}}
		{
			description: {{printf "%q" .Description}},
			input: maximumValueCaseInput {
				MaximumWeight: {{printf "%d" .Input.MaximumWeight}},
				Items: []Item {
					{{range .Input.Items}}
						{
							Weight: {{printf "%d" .Weight}},
							Value: {{printf "%d" .Value}},
						},
					{{end}}
				},
			},
			expected: {{printf "%d" .Expected}},
		},
	{{end}}
}
`
