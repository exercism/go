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
		"saddlePoints": &[]Case{},
	}
	if err := gen.Gen("saddle-points", j, t); err != nil {
		log.Fatal(err)
	}
}

type Case struct {
	Description string `json:"description"`
	Input       struct {
		Matrix [][]int `json:"matrix"`
	} `json:"input"`
	Expected []struct {
		Row    int `json:"row"`
		Column int `json:"column"`
	} `json:"expected"`
}

// Template to generate two sets of test cases, one for Score tests and one for Roll tests.
var tmpl = `{{.Header}}

var testCases = []struct {
	description    string
	input          [][]int
	expectedOutput []Pair
}{
	{{range .J.saddlePoints}}
		{
			description: {{printf "%q"  .Description}},
			input: [][]int{
				{{range .Input.Matrix}} { {{range .}} {{printf "%v" .}}, {{end}} }, {{end}}
			},
			expectedOutput: []Pair{
				{{range .Expected}}{
						{{printf "%d" .Row}},
						{{printf "%d" .Column}},
					},
				{{end}}
			},
		},
	{{end}}
}
`
