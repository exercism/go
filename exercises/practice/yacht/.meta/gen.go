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
		"score": &[]testCase{},
	}
	if err := gen.Gen("yacht", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Dice     []int  `json:"dice"`
		Category string `json:"category"`
	} `json:"input"`
	Expected int `json:"expected"`
}

var tmpl = `package yacht

{{.Header}}

var testCases = []struct {
	description    string
	dice           []int
	category       string
	expected          int 
}{ {{range .J.score}}
{
	description: 	{{printf "%q"  .Description}},
	dice: 			{{printf "%#v" .Input.Dice }} ,
	category: 		{{printf "%q" .Input.Category}},
	expected: 		{{printf "%d" .Expected}},
},{{end}}
}
`
