package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"abs":         &[]testCaseUnaryRationalToRational{},
		"add":         &[]testCaseBinaryRationalToRational{},
		"div":         &[]testCaseBinaryRationalToRational{},
		"mul":         &[]testCaseBinaryRationalToRational{},
		"sub":         &[]testCaseBinaryRationalToRational{},
		"exprational": &[]testCaseExprational{},
		"expreal":     &[]testCaseExpreal{},
		"reduce":      &[]testCaseUnaryRationalToRational{},
	}
	if err := gen.Gen("rational-numbers", j, t); err != nil {
		log.Fatal(err)
	}
}

type Rational [2]int

func (r Rational) String() string {
	return fmt.Sprintf("Rational{%d, %d}", r[0], r[1])
}

type RationalPair struct {
	R1 Rational `json:"r1"`
	R2 Rational `json:"r2"`
}

func (r RationalPair) String() string {
	return fmt.Sprintf("[2]Rational{{%d, %d}, {%d, %d}}", r.R1[0], r.R1[1], r.R2[0], r.R2[1])
}

type testCaseUnaryRationalToRational struct {
	Description string `json:"description"`
	Input       struct {
		R Rational `json:"r"`
	} `json:"input"`
	Expected Rational `json:"expected"`
}

type testCaseBinaryRationalToRational struct {
	Description string       `json:"description"`
	Input       RationalPair `json:"input"`
	Expected    Rational     `json:"expected"`
}

type testCaseExprational struct {
	Description string `json:"description"`
	Input       struct {
		R Rational `json:"r"`
		N int      `json:"n"`
	} `json:"input"`
	Expected Rational `json:"expected"`
}

type testCaseExpreal struct {
	Description string `json:"description"`
	Input       struct {
		X int      `json:"x"`
		R Rational `json:"r"`
	} `json:"input"`
	Expected float64 `json:"expected"`
}

var tmpl = `{{.Header}}

type testCaseAbs struct {
	description string
	input       Rational
	expected    Rational
}

var testCasesAbs = []testCaseAbs { {{range .J.abs}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{.Input.R}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseAdd struct {
	description string
	input       [2]Rational
	expected    Rational
}

var testCasesAdd = []testCaseAdd { {{range .J.add}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{.Input}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseSub struct {
	description string
	input       [2]Rational
	expected    Rational
}

var testCasesSub = []testCaseSub { {{range .J.sub}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{.Input}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseMul struct {
	description string
	input       [2]Rational
	expected    Rational
}

var testCasesMul = []testCaseMul { {{range .J.mul}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{.Input}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseDiv struct {
	description string
	input       [2]Rational
	expected    Rational
}

var testCasesDiv = []testCaseDiv { {{range .J.div}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{.Input}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseExprational struct {
	description string
	inputR      Rational
	inputInt    int
	expected    Rational
}

var testCasesExprational = []testCaseExprational { {{range .J.exprational}}{
	description: 	{{printf "%q" .Description}},
	inputR:         {{.Input.R}},
	inputInt:       {{.Input.N}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseExpreal struct {
	description string
	inputInt    int
	inputR      Rational
	expected    float64
}

var testCasesExpreal = []testCaseExpreal { {{range .J.expreal}}{
	description: 	{{printf "%q" .Description}},
	inputInt:       {{.Input.X}},
	inputR:         {{.Input.R}},
	expected:    	{{.Expected}},
},
{{end}}
}

type testCaseReduce struct {
	description string
	input       Rational
	expected    Rational
}

var testCasesReduce = []testCaseReduce { {{range .J.reduce}}{
	description: 	{{printf "%q" .Description}},
	input:	        {{.Input.R}},
	expected:    	{{.Expected}},
},
{{end}}
}
`
