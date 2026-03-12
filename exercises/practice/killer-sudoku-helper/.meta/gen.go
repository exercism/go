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
		"combinations": &[]testCase{},
	}
	if err := gen.Gen("killer-sudoku-helper", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Cage struct {
			Sum     int   `json:"sum"`
			Size    int   `json:"size"`
			Exclude []int `json:"exclude"`
		} `json:"cage"`
	} `json:"input"`
	Expected [][]int `json:"expected"`
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description string
	sum         int
	size        int
	exclude     []int
	expected    [][]int
}{ {{range .J.combinations}}
	{
		description: 	{{printf "%q" .Description}},
		sum:    	{{.Input.Cage.Sum}},
		size:    	{{.Input.Cage.Size}},
		exclude:    	{{printf "%#v" .Input.Cage.Exclude}},
		expected:    	{{printf "%#v" .Expected}},
	},{{end}}
}`
