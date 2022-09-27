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
	j := map[string]interface{}{
		"encode": &[]testCase{},
		"decode": &[]testCase{},
	}
	if err := gen.Gen("variable-length-quantity", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Integers []uint32 `json:"integers"` // supports both []byte and []uint32 in JSON
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ValueSliceByte() []byte {
	v, ok := t.Expected.([]interface{})
	var vals []byte
	if ok {
		for _, n := range v {
			number, ok := n.(float64)
			if !ok {
				log.Fatal("[ERROR] expected values in array in `expected` to be a number")
			}
			vals = append(vals, byte(number))
		}
	}
	return vals
}

func (t testCase) ValueSliceUint32() []uint32 {
	v, ok := t.Expected.([]interface{})
	var vals []uint32
	if ok {
		for _, n := range v {
			number, ok := n.(float64)
			if !ok {
				log.Fatal("[ERROR] expected values in array in `expected` to be a number")
			}
			vals = append(vals, uint32(number))
		}
	}
	return vals
}

func (t testCase) ErrorExpected() bool {
	v, ok := t.Expected.(map[string]interface{})
	if ok {
		_, ok := v["error"].(string)
		return ok
	}
	return false
}

// template applied to above data structure generates the Go test cases
var tmpl = `package variablelengthquantity

{{.Header}}

var encodeTestCases = []struct {
	description string
	input		[]uint32
	expected	[]byte
}{ 
{{range .J.encode}}{
		description: 	{{printf  "%q"  .Description    }},
		input: 			{{printf  "%#v" .Input.Integers }},
		expected: 		{{printf  "%#v" .ValueSliceByte }},
	},
{{end}}
}

var decodeTestCases = []struct {
	description     string
	input			[]byte
	expected		[]uint32
	errorExpected	bool
}{ {{range .J.decode}}{
		description: 	{{printf "%q" 	.Description}},
		input: 			[]byte{ {{range .Input.Integers}} {{printf "%#x" .}},{{end}}},
		expected: 		{{printf "%#v" 	.ValueSliceUint32 }},
		errorExpected: 	{{printf "%t" 	.ErrorExpected }},
	},
{{end}}
}
`
