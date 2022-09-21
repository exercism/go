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
		"sublist": &[]testCase{},
	}
	if err := gen.Gen("sublist", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		ListOne []int `json:"listOne"`
		ListTwo []int `json:"listTwo"`
	} `json:"input"`
	Expected string `json:"expected"`
}

var tmpl = `package sublist

{{.Header}}

var testCases = []struct {
	description	string
	listOne		[]int
	listTwo		[]int
	expected	Relation
}{ {{range .J.sublist}}
{
	description:	{{printf "%q"   .Description}},
	listOne:		{{printf "%#v"  .Input.ListOne}},
	listTwo:		{{printf "%#v"  .Input.ListTwo}},
	expected:		{{printf "%#v"  .Expected}},
},{{end}}
}
`
