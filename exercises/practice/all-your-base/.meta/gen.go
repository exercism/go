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
		"rebase": &[]testCase{},
	}
	if err := gen.Gen("all-your-base", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		InputBase  int   `json:"inputBase"`
		Digits     []int `json:"digits"`
		OutputBase int   `json:"outputBase"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (o testCase) Result() []int {
	s, ok := o.Expected.([]any)
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

func (o testCase) Err() string {
	m, ok := o.Expected.(map[string]any)
	if !ok {
		return ""
	}
	return m["error"].(string)
}

// Template to generate test cases.
var tmpl = `{{.Header}}

var testCases = []struct {
	description   string
	inputBase     int
	inputDigits   []int
	outputBase    int
	expected      []int
	expectedError string
}{ {{range .J.rebase}}
{
	description:	{{printf "%q"  .Description}},
	inputBase:		{{printf "%d"  .Input.InputBase}},
	inputDigits:		{{printf "%#v"  .Input.Digits}},
	outputBase:	{{printf "%d"  .Input.OutputBase}},
	expected:	{{printf "%#v"  .Result}},
	expectedError:	{{printf "%q"  .Err}},
},{{end}}
}
`
