package main

import (
	"encoding/json"
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
	if err := gen.Gen("largest-series-product", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Digits string
			Span   int
		}
		Expected ExpectedType
	}
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

// template applied to above data structure generates the Go test cases
var tmpl = `package lsproduct

{{.Header}}

var tests =	[]struct {
	digits	string
	span		int
	product	int64
	ok		bool
	error	string
}{
{{range .J.Cases}}{ "{{.Input.Digits}}", {{.Input.Span}}, {{.Expected.ValueInt}}, {{ge .Expected.ValueInt 0}}, {{printf "%q" .Expected.ValueString}}},
{{end}}}
`
