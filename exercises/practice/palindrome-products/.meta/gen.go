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
		"smallest": &[]testCase{},
		"largest":  &[]testCase{},
	}
	if err := gen.Gen("palindrome-products", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"input"`
	Expected struct {
		Value   int      `json:"value"`
		Factors [][2]int `json:"factors"`
		Error   string   `json:"error"`
	} `json:"expected"`
}

var tmpl = `{{.Header}}

type testCase struct {
	description string
	// input to Products(): range limits for factors of the palindrome
	fmin, fmax int
	// are we checking the min or max
	property string
	product int
	factors [][2]int
	errorMsg  string  // start of text if there is an error, "" otherwise
}

var testCases = []testCase{ {{range .J.smallest}}
	{
		description: {{printf "%q" .Description}},
		fmin:        {{.Input.Min}},
		fmax:        {{.Input.Max}},
		property:    {{printf "%q" .Property}},
		product:     {{.Expected.Value}},
		factors:     {{printf "%#v" .Expected.Factors}},
		errorMsg:       {{printf "%q" .Expected.Error}},
	},{{end}} {{range .J.largest}}
	{
		description: {{printf "%q" .Description}},
		fmin:        {{.Input.Min}},
		fmax:        {{.Input.Max}},
		property:    {{printf "%q" .Property}},
		product:     {{.Expected.Value}},
		factors:     {{printf "%#v" .Expected.Factors}},
		errorMsg:       {{printf "%q" .Expected.Error}},
	},{{end}}
}`
