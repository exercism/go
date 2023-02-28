package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

type colorCodePropertyCase struct {
	Description string `json:"description"`
	Input       struct {
		Color string `json:"color"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type colorsPropertyTestCase struct {
	Description string `json:"description"`
	Input       interface{}
	Expected    []string `json:"expected"`
}

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"colorCode": &[]colorCodePropertyCase{},
		"colors":    &[]colorsPropertyTestCase{},
	}
	if err := gen.Gen("resistor-color", j, t); err != nil {
		log.Fatal(err)
	}
}

// Template to generate test cases.
var tmpl = `package resistorcolor

{{.Header}}

type colorCodeTestCase struct {
	description	string
	input		string
	expected	int
}

type colorsTestCase struct {
	description	string
	input		string
	expected	[]string
}

var colorCodeTestCases = []colorCodeTestCase{
	{{range .J.colorCode}}
{
	description:	{{printf "%q"  .Description}},
	input:			{{printf "%q"  .Input.Color}},
	expected:		{{printf "%d"  .Expected}},
},{{end}}
}

var colorsTestCases = []colorsTestCase{
	{{range .J.colors}}
{
	description:	{{printf "%q"  .Description}},
	expected:		{{printf "%#v" .Expected}},
},{{end}}
}


`
