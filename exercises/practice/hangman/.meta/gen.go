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
		"guess": &[]testCase{},
	}
	if err := gen.Gen("hangman", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Word    string   `json:"word"`
		Guesses []string `json:"guesses"`
	} `json:"input"`
	Expected struct {
		State             string `json:"state"`
		MaskedWord        string `json:"maskedWord"`
		RemainingFailures int    `json:"remainingFailures"`
		Error             string `json:"error"`
	} `json:"expected"`
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description       string
	word              string
	guesses           []rune
	state             string
	maskedWord        string
	remainingFailures int
	error             string
}{ {{range .J.guess}}
	{
		description:       {{printf "%q" .Description}},
		word:              {{printf "%q" .Input.Word}},
		guesses:           []rune{ {{range .Input.Guesses}}'{{.}}', {{end}} },
		state:             {{printf "%q" .Expected.State}},
		maskedWord:        {{printf "%q" .Expected.MaskedWord}},
		remainingFailures: {{.Expected.RemainingFailures}},
		error:             {{printf "%q" .Expected.Error}},
	},{{end}}
}`
