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
	Expected interface{} `json:"expected"`
}

func (t testCase) Value() int {
	val, ok := t.Expected.(float64)
	if ok {
		return int(val)
	}
	return -1
}

func (t testCase) Error() string {
	if _, ok := t.Expected.(float64); !ok {
		m, ok := t.Expected.(map[string]interface{})
		if !ok {
			return ""
		}
		b, ok := m["error"].(string)
		if !ok {
			return ""
		}
		return b
	}
	return ""
}

// Template to generate test cases.
var tmpl = `package binarysearch

{{.Header}}

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
