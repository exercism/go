// +build ignore

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
	if err := gen.Gen("clock", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to umarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []OneCase
	}
}

type OneCase struct {
	Description string
	Property    string
	Hour        int // "create"/"add" cases
	Minute      int // "create"/"add" cases
	Add         int // "add" cases only

	Clock1   struct{ Hour, Minute int } // "equal" cases only
	Clock2   struct{ Hour, Minute int } // "equal" cases only
	Expected interface{}                // string or bool
}

func (c OneCase) IsTimeCase() bool  { return c.Property == "create" }
func (c OneCase) IsAddCase() bool   { return c.Property == "add" }
func (c OneCase) IsEqualCase() bool { return c.Property == "equal" }

var tmpl = `package clock

{{.Header}}

// Test creating a new clock with an initial time.
var timeTests = []struct {
	h, m int
	want string
}{ {{range .J.Cases}} {{range .Cases}}
{{if .IsTimeCase}}{ {{.Hour}}, {{.Minute}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
{{- end}}{{end}}{{end}} }

// Test adding and subtracting minutes.
var addTests = []struct {
	h, m, a int
	want string
}{ {{range .J.Cases}} {{range .Cases}}
{{if .IsAddCase}}{ {{.Hour}}, {{.Minute}}, {{.Add}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
{{- end}}{{end}}{{end}} }

// Construct two separate clocks, set times, test if they are equal.
type hm struct{ h, m int }

var eqTests = []struct {
	c1, c2 hm
	want   bool
}{ {{range .J.Cases}} {{range .Cases}}
{{if .IsEqualCase}} // {{.Description}}
{
	hm{ {{.Clock1.Hour}}, {{.Clock1.Minute}}},
	hm{ {{.Clock2.Hour}}, {{.Clock2.Minute}}},
	{{.Expected}},
}, {{- end}}{{end}}{{end}}
}
`
