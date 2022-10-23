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
	var j = map[string]interface{}{
		"recite": &[]testCase{},
	}
	if err := gen.Gen("bottle-song", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		StartBottles int `json:"startBottles"`
		TakeDown     int `json:"takeDown"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `package bottlesong

{{.Header}}
type bottleSongInput struct {
	startBottles	int
	takeDown		int
}

var testCases = []struct {
	description    string
	input          bottleSongInput
	expected       []string       
}{
	{{range .J.recite}}
	{
		description: {{printf "%q"  .Description}},
		input: bottleSongInput{
			startBottles: 	{{printf "%d"  .Input.StartBottles}},
			takeDown: 		{{printf "%d"  .Input.TakeDown}},
		},
		expected: {{printf "%#v"  .Expected}},
	},{{end}}
}
`
