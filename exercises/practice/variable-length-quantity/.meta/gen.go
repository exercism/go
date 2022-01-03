package main

import (
	"encoding/json"
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"GroupComment": GroupComment,
		"byteSlice":    byteSlice,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("variable-length-quantity", &j, t); err != nil {
		log.Fatal(err)
	}
}

// byteSlice converts a slice of uint32 to a byte slice.
func byteSlice(ns []uint32) []byte {
	b := make([]byte, len(ns))
	for i, n := range ns {
		b[i] = byte(n)
	}
	return b
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Groups []TestGroup `json:"Cases"`
}

type TestGroup struct {
	Description string
	Cases       []OneCase
}

type OneCase struct {
	Description string
	Property    string // "encode" or "decode"
	Input       struct {
		Integers []uint32 // supports both []byte and []uint32 in JSON.
	}
	Expected ExpectedType // can be {"error": "message"} or [x,y,z]
}

type ExpectedType struct {
	ValueSlice     []uint32
	ValueErrorFlag bool
}

func (e *ExpectedType) UnmarshalJSON(b []byte) error {
	if b[0] == '[' {
		var values []uint32
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		e.ValueSlice = values
		e.ValueErrorFlag = false
	} else {
		var s map[string]string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		e.ValueSlice = []uint32{}
		e.ValueErrorFlag = true
	}
	return nil
}

// PropertyMatch returns true when given test case c has .Property field matching property;
// this serves as a filter to put test cases with "like" property into the same group.
func (c OneCase) PropertyMatch(property string) bool { return c.Property == property }

// GroupComment looks in each of the test case groups to find the
// group for which every test case has the .Property matching given property;
// it returns the .Description field for the matching property group,
// or a 'Note: ...' if no test group consistently matches given property.
func GroupComment(groups []TestGroup, property string) string {
	for _, group := range groups {
		propertyGroupMatch := true
		for _, testcase := range group.Cases {
			if !testcase.PropertyMatch(property) {
				propertyGroupMatch = false
				break
			}
		}
		if propertyGroupMatch {
			return group.Description
		}
	}
	return "Note: Apparent inconsistent use of \"property\": \"" + property + "\" within test case group!"
}

// template applied to above data structure generates the Go test cases
var tmpl = `package variablelengthquantity

{{.Header}}

// {{GroupComment .J.Groups "encode"}}
var encodeTestCases = []struct {
	description string
	input	[]uint32
	output	[]byte
}{ {{range .J.Groups}} {{range .Cases}}
	{{if .PropertyMatch "encode"}} {
		{{printf "%q" .Description}},
		{{printf "%#v" .Input.Integers }},
		{{byteSlice .Expected.ValueSlice | printf "%#v" }},
	},{{- end}}{{end}}{{end}}
}

// {{GroupComment .J.Groups "decode"}}
var decodeTestCases = []struct {
	description string
	input			[]byte
	output			[]uint32
	errorExpected	bool
}{ {{range .J.Groups}} {{range .Cases}}
	{{if .PropertyMatch "decode"}} {
		{{printf "%q" .Description}},
		{{byteSlice .Input.Integers | printf "%#v" }},
		{{printf "%#v" .Expected.ValueSlice }},
		{{printf "%#v" .Expected.ValueErrorFlag }},
	},{{- end}}{{end}}{{end}}
}
`
