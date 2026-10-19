package main

import (
	"../../../../gen"
	"log"
	"strconv"
	"text/template"
)

type labelDecoderCase struct {
	Description string `json:"description"`
	Input       struct {
		Colors []string `json:"colors"`
	} `json:"input"`
	Expected struct {
		Value     float64 `json:"value"`
		Unit      string  `json:"unit"`
		Tolerance string  `json:"Tolerance"`
	} `json:"expected"`
}

func (v labelDecoderCase) InputColorsString() string {
	s := "[]string{"

	for _, c := range v.Input.Colors {
		s += `"` + c + `", `
	}

	s = s[:len(s)-2] + "}"
	return s
}

func (v labelDecoderCase) ExpectedString() string {
	s := strconv.FormatFloat(v.Expected.Value, 'g', -1, 64)
	s += " " + v.Expected.Unit
	if v.Expected.Tolerance != "" {
		s += " " + v.Expected.Tolerance
	}
	return s
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"labelDecoder": &[]labelDecoderCase{},
	}
	if err := gen.Gen("resistor-color-tolerance", j, t); err != nil {
		log.Fatal(err)
	}
}

// Template to generate test cases.
var tmpl = `{{.Header}}

type labelTestCase struct {
	description	string
	input		[]string
	expected	string
}

var labelTestCases = []labelTestCase{
	{{range .J.labelDecoder}} {
	description: {{printf "%q" .Description}},
	input:       {{printf "%v" .InputColorsString}},
	expected:    "{{.ExpectedString}}",
},
{{end}}
}



`
