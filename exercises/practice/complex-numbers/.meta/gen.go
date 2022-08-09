package main

import (
	"fmt"
	"log"
	"math"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j = map[string]interface{}{
		"real":      &[]Case{},
		"imaginary": &[]Case{},
		"mul":       &[]Case{},
		"add":       &[]Case{},
		"sub":       &[]Case{},
		"div":       &[]Case{},
		"abs":       &[]Case{},
		"conjugate": &[]Case{},
		"exp":       &[]Case{},
	}
	if err := gen.Gen("complex-numbers", j, t); err != nil {
		log.Fatal(err)
	}
}

type Case struct {
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Z  interface{} `json:"z"`
		Z1 interface{} `json:"z1"`
		Z2 interface{} `json:"z2"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

type complex struct {
	A float64
	B float64
}

func (c Case) GetZ() complex {
	return getComplex(c.Input.Z)
}

func (c Case) GetZ1() complex {
	return getComplex(c.Input.Z1)
}

func (c Case) GetZ2() complex {
	return getComplex(c.Input.Z2)
}

func (c Case) GetNeededZ() complex {
	_, ok := c.Input.Z1.(float64)
	if ok {
		return c.GetZ2()
	}
	return c.GetZ1()
}

func (c Case) IsSimpleFactor() bool {
	_, ok := c.Input.Z1.(float64)
	if ok {
		return ok
	}
	_, ok = c.Input.Z2.(float64)
	return ok
}

func (c Case) GetFactor() float64 {
	v, ok := c.Input.Z1.(float64)
	if ok {
		return v
	}
	v, ok = c.Input.Z2.(float64)
	return v
}

func (c Case) GetExpected() complex {
	return getComplex(c.Expected)
}

func getComplex(in interface{}) complex {
	v, ok := in.(float64)
	if ok {
		return complex{A: v}
	}
	s, ok := in.([]interface{})
	if !ok {
		panic(fmt.Sprintf("unknown value for field: %v", in))
	}
	var result []float64
	for _, element := range s {
		switch e := element.(type) {
		case float64:
			result = append(result, e)
		case string:
			v, ok := stringToNumber[e]
			if !ok {
				panic(fmt.Sprintf("value for %q not found", e))
			}
			result = append(result, v)
		default:
			panic("uknown type")
		}
	}
	return complex{
		A: result[0],
		B: result[1],
	}
}

var stringToNumber = map[string]float64{
	"e":       math.E,
	"pi":      math.Pi,
	"ln(2)":   math.Ln2,
	"ln(2)/2": math.Ln2 / 2,
	"pi/4":    math.Pi / 4,
}

// Template to generate two sets of test cases, one for Score tests and one for Roll tests.
var tmpl = `package complex

{{.Header}}

type complexNumber struct {
	a float64
	b float64
}

var realTestCases = []struct {
	description string
	in          complexNumber
	want    float64
}{
	{{range .J.real}}{
		description: {{printf "%q"  	.Description}},
	    in: complexNumber{
			a: {{printf "%f"  	.GetZ.A}},
			b: {{printf "%f"  	.GetZ.B}},
		},
		want:    {{printf "%f"  	.GetExpected.A}},
	},
	{{end}}
}

var imaginaryTestCases = []struct {
	description string
	in          complexNumber
	want    float64
}{
	{{range .J.imaginary}}{
		description: {{printf "%q"  	.Description}},
	    in: complexNumber{
			a: {{printf "%f"  	.GetZ.A}},
			b: {{printf "%f"  	.GetZ.B}},
		},
		want:  {{printf "%f"  	.GetExpected.A}},
	},
	{{end}}
}

var addTestCases = []struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want    complexNumber
}{
	{{range .J.add}}{
		description: {{printf "%q"  	.Description}},
	    n1: complexNumber{
			a: {{printf "%f"  	.GetZ1.A}},
			b: {{printf "%f"  	.GetZ1.B}},
		},
		n2: complexNumber{
			a: {{printf "%f"  	.GetZ2.A}},
			b: {{printf "%f"  	.GetZ2.B}},
		},
		want:    complexNumber{
			a: {{printf "%f"  	.GetExpected.A}},
			b: {{printf "%f"  	.GetExpected.B}},
		},
	},
	{{end}}
}

var subtractTestCases = []struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want    complexNumber
}{
	{{range .J.sub}}{
		description: {{printf "%q"  	.Description}},
	    n1: complexNumber{
			a: {{printf "%f"  	.GetZ1.A}},
			b: {{printf "%f"  	.GetZ1.B}},
		},
		n2: complexNumber{
			a: {{printf "%f"  	.GetZ2.A}},
			b: {{printf "%f"  	.GetZ2.B}},
		},
		want:    complexNumber{
			a: {{printf "%f"  	.GetExpected.A}},
			b: {{printf "%f"  	.GetExpected.B}},
		},
	},
	{{end}}
}

var divideTestCases = []struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want    complexNumber
}{
	{{range .J.div}}{
		description: {{printf "%q"  	.Description}},
	    n1: complexNumber{
			a: {{printf "%f"  	.GetZ1.A}},
			b: {{printf "%f"  	.GetZ1.B}},
		},
		n2: complexNumber{
			a: {{printf "%f"  	.GetZ2.A}},
			b: {{printf "%f"  	.GetZ2.B}},
		},
		want:    complexNumber{
			a: {{printf "%f"  	.GetExpected.A}},
			b: {{printf "%f"  	.GetExpected.B}},
		},
	},
	{{end}}
}

var multiplyTestCases = []struct {
	description string
	n1          complexNumber
	n2          *complexNumber // if n2 is nil it is a multiplication with the factor
	factor 		float64
	want    	complexNumber
}{
	{{range .J.mul}}{
		description: {{printf "%q"  	.Description}},
		{{- if .IsSimpleFactor }}
			n1: complexNumber{
				a: {{printf "%f"  	.GetNeededZ.A}},
				b: {{printf "%f"  	.GetNeededZ.B}},
			},
		{{else}}
			n1: complexNumber{
				a: {{printf "%f"  	.GetZ1.A}},
				b: {{printf "%f"  	.GetZ1.B}},
			},
		{{- end}}
		{{- if .IsSimpleFactor }}
			n2: nil,
		{{else}}
			n2: &complexNumber{
				a: {{printf "%f"  	.GetZ2.A}},
				b: {{printf "%f"  	.GetZ2.B}},
			},
		{{- end}}
		factor: {{printf "%f"  	.GetFactor}},
		want:    complexNumber{
			a: {{printf "%f"  	.GetExpected.A}},
			b: {{printf "%f"  	.GetExpected.B}},
		},
	},
	{{end}}
}

var conjugateTestCases = []struct {
	description string
	in          complexNumber
	want    	complexNumber
}{
	{{range .J.conjugate}}{
		description: {{printf "%q"  	.Description}},
	    in: complexNumber{
			a: {{printf "%f"  	.GetZ.A}},
			b: {{printf "%f"  	.GetZ.B}},
		},
		want:    complexNumber{
			a: {{printf "%f"  	.GetExpected.A}},
			b: {{printf "%f"  	.GetExpected.B}},
		},
	},
	{{end}}
}

var absTestCases = []struct {
	description string
	in          complexNumber
	want    float64
}{
	{{range .J.abs}}{
		description: {{printf "%q"  	.Description}},
	    in: complexNumber{
			a: {{printf "%f"  	.GetZ.A}},
			b: {{printf "%f"  	.GetZ.B}},
		},
		want:    {{printf "%f"  	.GetExpected.A}},
	},
	{{end}}
}

var expTestCases = []struct {
	description string
	in          complexNumber
	want    complexNumber
}{
	{{range .J.exp}}{
		description: {{printf "%q"  	.Description}},
	    in: complexNumber{
			a: {{printf "%f"  	.GetZ.A}},
			b: {{printf "%f"  	.GetZ.B}},
		},
		want:    complexNumber{
			a: {{printf "%f"  	.GetExpected.A}},
			b: {{printf "%f"  	.GetExpected.B}},
		},
	},
	{{end}}
}

`
