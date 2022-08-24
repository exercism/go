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
		"canChain": &[]testCase{},
	}
	if err := gen.Gen("dominoes", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string   `json:"description"`
	Comments    []string `json:"comments"`
	Input       struct {
		Dominoes [][]int `json:"dominoes"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

var tmpl = `package dominoes

{{.Header}}

var testCases = []struct {
	description string
	dominoes    []Domino
	valid       bool // true => can chain, false => cannot chain
}{
	{{range .J.canChain}}{
			description: {{printf "%q" .Description}},
			{{range .Comments}}// {{.}}
			{{end}}dominoes: []Domino{ {{range .Input.Dominoes}}{ {{range .}} {{printf "%v" .}}, {{end}}}, {{end}} },
			valid: {{printf "%v" .Expected}},
		},
	{{end}}
}
`
