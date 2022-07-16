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
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectedValues() []int {
	values, ok := t.Expected.([]interface{})
	if !ok {
		return nil
	}
	var intValues = make([]int, 0)
	for _, v := range values {
		i, ok := v.(float64)
		if !ok {
			return nil
		}
		intValues = append(intValues, int(i))
	}
	return intValues
}

func (t testCase) Valid() bool {
	return t.ExpectedValues() != nil
}

var tmpl = `package change

{{.Header}}

var testCases = []struct {
	description    string
	coins          []int
	target         int
	valid          bool     // true => no error, false => error expected
	expectedChange []int    // when .valid == true, the expected change coins
}{
{{range .J.findFewestCoins}}{
	description: 		{{printf "%q"  .Description}},
	coins: 				{{printf "%#v" .Input.Coins}},
	target: 			{{printf "%d"  .Input.Target}},
	valid: 				{{printf "%v"  .Valid}},
	expectedChange: 	{{printf "%#v" .ExpectedValues}},
},
{{end}}}
`
