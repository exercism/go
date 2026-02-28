package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

type Position struct {
	X     float64
	Y     float64
	Angle float64
}

type Prism struct {
	ID    int
	X     float64
	Y     float64
	Angle float64
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"findSequence": &[]testCase{},
	}
	if err := gen.Gen("prism", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Start  Position `json:"start"`
		Prisms []Prism  `json:"prisms"`
	} `json:"input"`
	Expected struct {
		Sequence []int `json:"sequence"`
	} `json:"expected"`
}

var tmpl = `package prism

{{.Header}}

var testCases = []struct {
	description string
	start       Position
	prisms      []Prism
	expected    []int
}{
{{range .J.findSequence}}{
	description: 	{{printf "%q" .Description}},
	start:          Position{x: {{.Input.Start.X}}, y: {{.Input.Start.Y}}, angle: {{.Input.Start.Angle}}, },
	prisms:         []Prism{
		{{range .Input.Prisms}} { id: {{.ID}}, x: {{.X}}, y: {{.Y}}, angle: {{.Angle}}, }, {{end}}
	},
	expected:    	[]int{ {{range .Expected.Sequence}} {{.}}, {{end}} },
},
{{end}}
}`
