package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("all-your-base", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Cases    []oneCase
}

// Test cases
type oneCase struct {
	Description string
	Property    string
	InputBase   int   `json:"input_base"`
	InputDigits []int `json:"input_digits"`
	OutputBase  int   `json:"output_base"`
	Expected    interface{}
}

func (o oneCase) Result() []int {
	s, ok := o.Expected.([]interface{})
	if !ok {
		return nil
	}
	var res []int
	for _, v := range s {
		f, _ := v.(float64)
		res = append(res, int(f))
	}
	return res
}
func (o oneCase) Err() string {
	m, ok := o.Expected.(map[string]interface{})
	if !ok {
		return ""
	}
	return m["error"].(string)
}

// Template to generate test cases.
var tmpl = `package allyourbase

{{.Header}}

var testCases = []struct {
	description	string
	inputBase	int
	inputDigits	[]int
	outputBase	int
	expected    []int
	err         string
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	inputBase:		{{printf "%d"  .InputBase}},
	inputDigits:		{{printf "%#v"  .InputDigits}},
	outputBase:	{{printf "%d"  .OutputBase}},
	expected:	{{printf "%#v"  .Result}},
	err:	{{printf "%q"  .Err}},
},{{end}}
}
`
