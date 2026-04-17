package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"math"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"real":      &[]TestCase{},
		"imaginary": &[]TestCase{},
		"mul":       &[]TestCase{},
		"add":       &[]TestCase{},
		"sub":       &[]TestCase{},
		"div":       &[]TestCase{},
		"abs":       &[]TestCase{},
		"conjugate": &[]TestCase{},
		"exp":       &[]TestCase{},
	}
	if err := gen.Gen("complex-numbers", j, t); err != nil {
		log.Fatal(err)
	}
}

type TestCase struct {
	UUID        string `json:"uuid"`
	Description string `json:"description"`
	Property    string `json:"property"`
	Input       struct {
		Z  any `json:"z"`
		Z1 any `json:"z1"`
		Z2 any `json:"z2"`
	} `json:"input"`
	Expected any `json:"expected"`
}

type complexNumber struct {
	A float64
	B float64
}

func (tc TestCase) GetDescription() string {
	return fmt.Sprintf("%q", strings.ReplaceAll(tc.Description, "/", " "))
}

func (tc TestCase) GetZ() complexNumber {
	return getComplex(tc.Input.Z)
}

func (tc TestCase) GetZ1() complexNumber {
	return getComplex(tc.Input.Z1)
}

func (tc TestCase) GetZ2() complexNumber {
	return getComplex(tc.Input.Z2)
}

func (tc TestCase) GetNeededZ() complexNumber {
	_, ok := tc.Input.Z1.(float64)
	if ok {
		return tc.GetZ2()
	}
	return tc.GetZ1()
}

func (tc TestCase) IsSimpleFactor() bool {
	_, ok := tc.Input.Z1.(float64)
	if ok {
		return ok
	}
	_, ok = tc.Input.Z2.(float64)
	return ok
}

func (tc TestCase) GetFactor() float64 {
	v, ok := tc.Input.Z1.(float64)
	if ok {
		return v
	}
	v, _ = tc.Input.Z2.(float64)
	return v
}

func (tc TestCase) GetExpected() complexNumber {
	return getComplex(tc.Expected)
}

func getComplex(in any) complexNumber {
	v, ok := in.(float64)
	if ok {
		return complexNumber{A: v}
	}
	s, ok := in.([]any)
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
	return complexNumber{
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
var tmpl = `{{.Header}}

type complexNumber struct {
	a float64
	b float64
}

type testCaseReal struct {
	description string
	in          complexNumber
	want    float64
}

var realTestCases = []testCaseReal {
	{{range .J.real}}{
		description: {{.GetDescription}},
		in: complexNumber{
			a: {{printf "%f" .GetZ.A}},
			b: {{printf "%f" .GetZ.B}},
		},
		want:    {{printf "%f" .GetExpected.A}},
	},
	{{end}}
}

type testCaseImag struct {
	description string
	in          complexNumber
	want    float64
}

var imaginaryTestCases = []testCaseImag {
	{{range .J.imaginary}}{
		description: {{.GetDescription}},
		in: complexNumber{
			a: {{printf "%f" .GetZ.A}},
			b: {{printf "%f" .GetZ.B}},
		},
		want:  {{printf "%f" .GetExpected.A}},
	},
	{{end}}
}

type testCaseAdd struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want    complexNumber
}

var addTestCases = []testCaseAdd {
	{{range .J.add}}{
		description: {{.GetDescription}},
		n1: complexNumber{
			a: {{printf "%f" .GetZ1.A}},
			b: {{printf "%f" .GetZ1.B}},
		},
		n2: complexNumber{
			a: {{printf "%f" .GetZ2.A}},
			b: {{printf "%f" .GetZ2.B}},
		},
		want:    complexNumber{
			a: {{printf "%f" .GetExpected.A}},
			b: {{printf "%f" .GetExpected.B}},
		},
	},
	{{end}}
}

type testCaseSubtract struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want    complexNumber
}

var subtractTestCases = []testCaseSubtract {
	{{range .J.sub}}{
		description: {{.GetDescription}},
		n1: complexNumber{
			a: {{printf "%f" .GetZ1.A}},
			b: {{printf "%f" .GetZ1.B}},
		},
		n2: complexNumber{
			a: {{printf "%f" .GetZ2.A}},
			b: {{printf "%f" .GetZ2.B}},
		},
		want:    complexNumber{
			a: {{printf "%f" .GetExpected.A}},
			b: {{printf "%f" .GetExpected.B}},
		},
	},
	{{end}}
}

type testCaseDivide struct {
	description string
	n1          complexNumber
	n2          complexNumber
	want    complexNumber
}

var divideTestCases = []testCaseDivide {
	{{range .J.div}}{
		description: {{.GetDescription}},
		n1: complexNumber{
			a: {{printf "%f" .GetZ1.A}},
			b: {{printf "%f" .GetZ1.B}},
		},
		n2: complexNumber{
			a: {{printf "%f" .GetZ2.A}},
			b: {{printf "%f" .GetZ2.B}},
		},
		want:    complexNumber{
			a: {{printf "%f" .GetExpected.A}},
			b: {{printf "%f" .GetExpected.B}},
		},
	},
	{{end}}
}

type testCaseMultiply struct {
	description string
	n1     complexNumber
	n2     *complexNumber // if n2 is nil it is a multiplication with the factor
	factor float64
	want   complexNumber
}

var multiplyTestCases = []testCaseMultiply {
	{{range .J.mul}}{
		description: {{.GetDescription}},
		{{- if .IsSimpleFactor }}
			n1: complexNumber{
				a: {{printf "%f" .GetNeededZ.A}},
				b: {{printf "%f" .GetNeededZ.B}},
			},
		{{else}}
			n1: complexNumber{
				a: {{printf "%f" .GetZ1.A}},
				b: {{printf "%f" .GetZ1.B}},
			},
		{{- end}}
		{{- if .IsSimpleFactor }}
			n2: nil,
		{{else}}
			n2: &complexNumber{
				a: {{printf "%f" .GetZ2.A}},
				b: {{printf "%f" .GetZ2.B}},
			},
		{{- end}}
		factor: {{printf "%f" .GetFactor}},
		want:   complexNumber{
			a: {{printf "%f" .GetExpected.A}},
			b: {{printf "%f" .GetExpected.B}},
		},
	},
	{{end}}
}

type testCaseConjugate struct {
	description string
	in          complexNumber
	want    	complexNumber
}

var conjugateTestCases = []testCaseConjugate {
	{{range .J.conjugate}}{
		description: {{.GetDescription}},
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

type testCaseAbs struct {
	description string
	in          complexNumber
	want    float64
}

var absTestCases = []testCaseAbs {
	{{range .J.abs}}{
		description: {{.GetDescription}},
	    in: complexNumber{
			a: {{printf "%f"  	.GetZ.A}},
			b: {{printf "%f"  	.GetZ.B}},
		},
		want:    {{printf "%f"  	.GetExpected.A}},
	},
	{{end}}
}

type testCaseExp struct {
	description string
	in          complexNumber
	want    complexNumber
}

var expTestCases = []testCaseExp {
	{{range .J.exp}}{
		description: {{.GetDescription}},
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
