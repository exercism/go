// +build ignore

package main

import (
	"log"
	"text/template"

	"../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("change", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Exercise string
	Version  string
	Comments []string
	Cases    []OneCase
}

// template applied to above data structure generates the Go test cases

type OneCase struct {
	Description string
	Property    string
	Coins       []int
	Target      int
	Expected    interface{}
}

func (c OneCase) Valid() bool {
	valid, _ := determineExpected(c.Expected)
	return valid
}

func (c OneCase) IntSlice() []int {
	_, list := determineExpected(c.Expected)
	return list
}

// determineExpected examines an .Expected interface{} object and determines
// whether an error is indicated by an int value (actually -1) in the JSON,
// or whether .Expected is a slice of integer coins.
func determineExpected(expected interface{}) (valid bool, list []int) {
	_, ok := expected.(int)
	if ok {
		return false, nil
	}
	ilist, ok := expected.([]interface{})
	if !ok {
		return false, nil
	}
	list = make([]int, 0)
	for _, iv := range ilist {
		// The literals from the JSON are unmarshalled to float64 values,
		// which are converted to int for the template output.
		v, isFloat64 := iv.(float64)
		if isFloat64 {
			list = append(list, int(v))
		}
	}
	return true, list
}

var tmpl = `package change

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

var testCases = []struct {
	description    string
	coins          []int
	target         int
	valid          bool     // true => no error, false => error expected
	expectedChange []int    // when .valid == true, the expected change coins
}{
{{range .J.Cases}}{
	{{printf "%q"  .Description}},
	{{printf "%#v" .Coins}},
	{{printf "%d"  .Target}},
	{{printf "%v"  .Valid}},
	{{printf "%#v" .IntSlice}},
},
{{end}}}
`
