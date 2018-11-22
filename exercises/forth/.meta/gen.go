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
	if err := gen.Gen("forth", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Groups []Group `json:"cases"`
}

// A group of test cases.
type Group struct {
	Name  string `json:"description"`
	Cases []OneCase
}

// One test case.
type OneCase struct {
	Description string
	Input       struct {
		Instructions []string
	}
	Expected interface{}
}

// IntSlice converts an .Expected interface{} object
// to either nil when an error is indicated by a map[string]interface{} in the JSON,
// or a slice of integers.
func (c OneCase) IntSlice() (list []int) {
	_, ok := c.Expected.(map[string]interface{})
	if ok {
		return nil
	}
	ilist, ok := c.Expected.([]interface{})
	if !ok {
		return nil
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
	return list
}

// template applied to above data structure generates the Go test cases
var tmpl = `package forth

{{.Header}}

type testGroup struct {
	group	string
	tests	[]testCase
}

type testCase struct {
	description	string
	input		[]string
	expected	[]int // nil slice indicates error expected.
}

var testGroups = []testGroup{
{{range .J.Groups}}{
group: {{printf "%q"  .Name}},
tests: []testCase{
{{range .Cases}}{
{{printf "%q"  .Description}},
{{printf "%#v" .Input.Instructions}},
{{printf "%#v" .IntSlice}},
},
{{end}}
},
},
{{end}}
}
`
