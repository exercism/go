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
		"findFewestCoins": &[]testCase{},
	}
	if err := gen.Gen("change", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Coins  []int `json:"coins"`
		Target int   `json:"target"`
	} `json:"input"`
	Expected any `json:"expected"`
}

func (t testCase) ExpectedValues() []int {
	return gen.FloatSliceToInts(t.Expected)
}

func (t testCase) Valid() bool {
	return t.ExpectedValues() != nil
}

var tmpl = `{{.Header}}

type testCase struct {
	description    string
	coins          []int
	target         int
	valid          bool     // true => no error, false => error expected
	expectedChange []int    // when .valid == true, the expected change coins
}

var testCases = []testCase { {{range .J.findFewestCoins}}
	{
		description: 		{{printf "%q"  .Description}},
		coins: 				{{printf "%#v" .Input.Coins}},
		target: 			{{printf "%d"  .Input.Target}},
		valid: 				{{printf "%v"  .Valid}},
		expectedChange: 	{{printf "%#v" .ExpectedValues}},
	},{{end}}
}
`
