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
		"latest":              &[]intTestCase{},
		"latestAfterBest":     &[]intTestCase{},
		"latestAfterTopThree": &[]intTestCase{},
		"personalBest":        &[]intTestCase{},
		"personalTopThree":    &[]sliceTestCase{},
		"scores":              &[]sliceTestCase{},
		"scoresAfterBest":     &[]sliceTestCase{},
		"scoresAfterTopThree": &[]sliceTestCase{},
	}
	if err := gen.Gen("high-scores", j, t); err != nil {
		log.Fatal(err)
	}
}

type intTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Scores []int `json:"scores"`
	} `json:"input"`
	Expected int `json:"expected"`
}

type sliceTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Scores []int `json:"scores"`
	} `json:"input"`
	Expected []int `json:"expected"`
}

var tmpl = `{{.Header}}

var intTestCases = []struct {
	description  string
	input        []int
	call         func(h *HighScores) int
	callMsg      string
	expected     int
}{ {{range .J.latest}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) int { return h.Latest(); },
		callMsg:        "h.Latest()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}} {{range .J.personalBest}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) int { return h.PersonalBest(); },
		callMsg:        "h.PersonalBest()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}} {{range .J.latestAfterBest}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) int { h.PersonalBest(); return h.Latest(); },
		callMsg:        "h.PersonalBest(); h.Latest()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}} {{range .J.latestAfterTopThree}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) int { h.TopThree(); return h.Latest(); },
		callMsg:        "h.TopThree(); h.Latest()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}}
}

var sliceTestCases = []struct {
	description  string
	input        []int
	call         func(h *HighScores) []int
	callMsg      string
	expected     []int
}{ {{range .J.scores}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) []int { return h.Scores(); },
		callMsg:        "h.Scores()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}} {{range .J.personalTopThree}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) []int { return h.TopThree(); },
		callMsg:        "h.TopThree()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}} {{range .J.scoresAfterBest}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) []int { h.PersonalBest(); return h.Scores() },
		callMsg:        "h.PersonalBest(); h.Scores()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}} {{range .J.scoresAfterTopThree}}
	{
		description: 	{{printf "%q" .Description}},
		input:	        {{printf "%#v" .Input.Scores}},
		call:           func(h *HighScores) []int { h.TopThree(); return h.Scores() },
		callMsg:        "h.TopThree(); h.Scores()",
		expected:    	{{printf "%#v" .Expected}},
	},{{end}}
}`
