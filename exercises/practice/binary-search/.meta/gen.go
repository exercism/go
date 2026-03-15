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
		"find": &[]testCase{},
	}
	if err := gen.Gen("binary-search", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Uuid        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Array []int `json:"array"`
		Value int   `json:"value"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) Value() int {
	val, ok := t.Expected.(float64)
	if ok {
		return int(val)
	}
	return -1
}

func (t testCase) Error() string {
	return gen.ErrorMessage(t.Expected)
}

// Template to generate test cases.
var tmpl = `{{.Header}}

var testCases =	[]struct {
	description string
	inputList   []int
	inputKey    int
	expectedKey int
	err         string
}{ {{range .J.find}}
{
	description:	    {{printf "%q"  .Description}},
	inputList:		    {{printf "%#v" .Input.Array}},
	inputKey:			{{printf "%d"  .Input.Value}},
	expectedKey:	    {{printf "%d"  .Value}},
	err:			    {{printf "%q"  .Error}},
},{{end}}
}
`
