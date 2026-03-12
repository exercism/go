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
		"simulateGame": &[]testCase{},
	}
	if err := gen.Gen("camicia", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		PlayerA []string `json:"playerA"`
		PlayerB []string `json:"playerB"`
	} `json:"input"`
	Expected struct {
		Status string `json:"status"`
		Cards  int    `json:"cards"`
		Tricks int    `json:"tricks"`
	} `json:"expected"`
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	playerA     []string
	playerB     []string
	expected    Outcome
}{ {{range .J.simulateGame}}
	{
		description: {{printf "%q" .Description}},
		playerA:     {{printf "%#v" .Input.PlayerA}},
		playerB:     {{printf "%#v" .Input.PlayerB}},
		expected:    Outcome {
			finishes: {{eq .Expected.Status "finished" | printf "%t"}},
			cards:  {{.Expected.Cards}},
			tricks: {{.Expected.Tricks}},
		},
	},{{end}}
}`
