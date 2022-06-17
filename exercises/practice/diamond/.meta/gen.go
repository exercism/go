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
	var j js
	if err := gen.Gen("diamond", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string `json:"description"`
		Input       struct {
			Letter string `json:"letter"`
		} `json:"input"`
		Expected []string `json:"expected"`
	} `json:"cases"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package diamond

{{.Header}}

var testCases = []struct {
	description string
	input       string
	expected    []string
}{
{{range .J.Cases}}{
		description: {{printf "%q"   .Description}},
	input: {{printf "%q"   .Input.Letter}},
	expected: []string { {{range $line := .Expected}}{{printf "\n%q," $line}}{{end}}
},
},
{{end}}
}`
