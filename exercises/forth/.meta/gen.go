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
	Input       []string
	Expected    []int
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
{{printf "%#v" .Input}},
{{printf "%#v" .Expected}},
},
{{end}}
},
},
{{end}}
}
`
