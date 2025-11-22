package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

type labelPropertyCase struct {
	Description string `json:"description"`
	Input       struct {
		Colors []string `json:"colors"`
	} `json:"input"`
	Expected struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"expected"`
}

func (v labelPropertyCase) InputColorsString() string {
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
		"label": &[]labelPropertyCase{},
	}
	if err := gen.Gen("resistor-color-trio", j, t); err != nil {
		log.Fatal(err)
	}
}

// Template to generate test cases.
var tmpl = `package resistorcolortrio

{{.Header}}

type labelTestCase struct {
	description	string
	input		[]string
	expected	string
}

var labelTestCases = []labelTestCase{
	{{range .J.label}} {
	description:	{{printf "%q" .Description}},
	input:			{{printf "%v" .InputColorsString}},
	expected:		"{{printf "%d" .Expected.Value}} {{printf "%s" .Expected.Unit}}",
},
{{end}}
}



`
