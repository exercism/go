package main

import (
	"../../../gen"
	"fmt"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("complex-numbers", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Groups []Group `json:"cases"`
}

// A group of test cases.
type Group struct {
	Name     string `json:"description"`
	Cases    []OneCase
	Input    InputType
	Expected interface{}
}

// One test case.
type OneCase struct {
	Description string
	Input       InputType
	Expected    interface{}
	Cases       []OneCase
}

type InputType struct {
	Z  []interface{}
	Z1 []interface{}
	Z2 []interface{}
}

func (c OneCase) InputContainsZ() bool {
	return len(c.Input.Z) != 0
}

func (grp Group) IsImaginaryPart() bool {
	fmt.Println(grp.Cases)
	return len(grp.Cases) == 0
}

// template applied to above data structure generates the Go test cases
var tmpl = `package complexnumbers

{{.Header}}

type testGroup struct {
	group	  string
	tests	  []testCase

	input	  inputType
    expected  interface{}
}

type testCase struct {
	description	string
	input	    inputType
	expected	interface{}

	group	    string
	tests	    []testCase
	
}

type inputType struct {
	z  []interface{}
	z1 []interface{}
	z2 []interface{}
}

var testGroups = []testGroup{
{{range .J.Groups}}{
group: {{printf "%q"  .Name}},
{{if .IsImaginaryPart}}input : inputType{
z1:{{printf "%#v"  .Input.Z1}},
z2:{{printf "%#v"  .Input.Z2}},
},
expected:{{printf "%#v" .Expected}},
{{else}}tests: []testCase{
{{range .Cases}}{
{{if .InputContainsZ}}description:{{printf "%q"  .Description}},
input:inputType{
z:{{printf "%#v"  .Input.Z}},
},
expected:{{printf "%#v" .Expected}},
{{else}}group:{{printf "%q"  .Description}},
tests: []testCase{
{{range .Cases}}{
description:{{printf "%q"  .Description}},
input : inputType{
z1:{{printf "%#v"  .Input.Z1}},
z2:{{printf "%#v"  .Input.Z2}},
},
expected:{{printf "%#v" .Expected}},
},
{{end}}
}
{{end}}
},
{{end}}
},
{{end}}
},
{{end}}
}
`
