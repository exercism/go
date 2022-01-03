package main

import (
	"encoding/json"
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("binary-search", &j, t); err != nil {
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
	Input       struct {
		Array []int
		Value int
	}
	Expected ExpectedType
}

type ExpectedType struct {
	ValueInt    int
	ValueString string
}

func (e *ExpectedType) UnmarshalJSON(b []byte) error {
	if b[0] != '{' {
		var i int
		if err := json.Unmarshal(b, &i); err != nil {
			return err
		}
		e.ValueInt = i
		return nil
	}
	var s map[string]string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	e.ValueInt = -1
	e.ValueString = s["error"]
	return nil
}

// Template to generate test cases.
var tmpl = `package binarysearch

{{.Header}}

var testCases =	[]struct {
	description	string
	slice		[]int
	key			int
	x			int
	err			string
}{ {{range .J.Cases}}
{
	description:	{{printf "%q"  .Description}},
	slice:		{{printf "%#v" .Input.Array}},
	key:			{{printf "%d"  .Input.Value}},
	x:			{{printf "%d"  .Expected.ValueInt}},
	err:			{{printf "%q"  .Expected.ValueString}},
},{{end}}
}
`
