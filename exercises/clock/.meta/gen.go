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
	Groups TestGroups `json:"Cases"`
}

type TestGroups []struct {
	Description string
	Cases       []OneCase
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

func (groups TestGroups) GroupComment(property string) string {
	for _, group := range groups {
		propertyGroupMatch := true
		for _, testcase := range group.Cases {
			if testcase.Property != property {
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

var tmpl = `package clock

{{.Header}}

{{with .J.Groups}}
    // {{ .GroupComment "create"}}
{{end}} var timeTests = []struct {
	h, m int
	want string
}{ {{range .J.Groups}} {{range .Cases}}
{{if .IsTimeCase}}{ {{.Hour}}, {{.Minute}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
{{- end}}{{end}}{{end}} }

{{with .J.Groups}}
    // {{ .GroupComment "add"}}
{{end}} var addTests = []struct {
	h, m, a int
	want string
}{ {{range .J.Groups}} {{range .Cases}}
{{if .IsAddCase}}{ {{.Hour}}, {{.Minute}}, {{.Add}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
{{- end}}{{end}}{{end}} }

{{with .J.Groups}}
    // {{ .GroupComment "equal"}}
{{end}} type hm struct{ h, m int }

var eqTests = []struct {
	c1, c2 hm
	want   bool
}{ {{range .J.Groups}} {{range .Cases}}
{{if .IsEqualCase}} // {{.Description}}
{
	hm{ {{.Clock1.Hour}}, {{.Clock1.Minute}}},
	hm{ {{.Clock2.Hour}}, {{.Clock2.Minute}}},
	{{.Expected}},
}, {{- end}}{{end}}{{end}}
}
`
