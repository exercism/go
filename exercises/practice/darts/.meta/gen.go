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
	j := map[string]interface{}{
		"score": &[]testCase{},
	}
	if err := gen.Gen("darts", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"input"`
	Expected int `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package darts

{{.Header}}

var testCases = []struct {
	description	string
	x		float64
	y		float64
	expected	int
}{ {{range .J.score}}
{
	description:	{{printf "%q"  .Description}},
	x:				{{printf "%.1f"  .Input.X}},
	y:				{{printf "%.1f"  .Input.Y}},
	expected:		{{printf "%d"  .Expected}},
},{{end}}
}
`
