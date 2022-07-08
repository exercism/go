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
	var j = map[string]interface{}{
		"saddlePoints": &[]Case{},
	}
	if err := gen.Gen("saddle-points", j, t); err != nil {
		log.Fatal(err)
	}
}

type Case struct {
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Matrix [][]int `json:"matrix"`
	} `json:"input"`
	Expected []struct {
		Row    int `json:"row"`
		Column int `json:"column"`
	} `json:"expected"`
}

// Template to generate two sets of test cases, one for Score tests and one for Roll tests.
var tmpl = `package matrix

{{.Header}}

var testCases = []struct {
	description    string
	input          [][]int
	expectedOutput []Pair
}{
	{{range .J.saddlePoints}}
		{
			{{printf "%q"  .Description}},
			{{printf "%#v" .Input.Matrix}},
			[]Pair{
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
