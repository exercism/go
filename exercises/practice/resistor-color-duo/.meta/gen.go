package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

type valuePropertyCase struct {
	Description string `json:"description"`
	Input       struct {
		Colors []string `json:"colors"`
	} `json:"input"`
	Expected int `json:"expected"`
}

func (v valuePropertyCase) InputColorsString() string {
	s := "[]string{"

	for _, c := range v.Input.Colors {
		s += `"` + c + `", `
	}

	s = s[:len(s)-2] + "}"
	return s
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"value": &[]valuePropertyCase{},
	}
	if err := gen.Gen("resistor-color-duo", j, t); err != nil {
		log.Fatal(err)
	}
}

// Template to generate test cases.
var tmpl = `package resistorcolorduo

{{.Header}}

type valueTestCase struct {
	description	string
	input		[]string
	expected	int
}

var valueTestCases = []valueTestCase{
	{{range .J.value}}
{
	description:	{{printf "%q"  .Description}},
	input:			{{printf "%v" .InputColorsString}},
	expected:		{{printf "%d"  .Expected}},
},{{end}}
}



`
