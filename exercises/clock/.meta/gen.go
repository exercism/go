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

type OneCase map[string]interface{}

func (c OneCase) IsProperty(property string) bool { return c["property"].(string) == property }

func (groups TestGroups) GroupComment(property string) string {
	for _, group := range groups {
		propertyGroupMatch := true
		for _, testcase := range group.Cases {
			if !testcase.IsProperty(property) {
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
{{if .IsProperty "create"}}{ {{.hour}}, {{.minute}}, {{.expected | printf "%#v"}}}, // {{.description}}
{{- end}}{{end}}{{end}} }

{{with .J.Groups}}
    // {{ .GroupComment "add"}}
{{end}} var addTests = []struct {
	h, m, a int
	want string
}{ {{range .J.Groups}} {{range .Cases}}
{{if .IsProperty "add"}}{ {{.hour}}, {{.minute}}, {{.add}}, {{.expected | printf "%#v"}}}, // {{.description}}
{{- end}}{{end}}{{end}} }

{{with .J.Groups}}
    // {{ .GroupComment "equal"}}
{{end}} type hm struct{ h, m int }

var eqTests = []struct {
	c1, c2 hm
	want   bool
}{ {{range .J.Groups}} {{range .Cases}}
{{if .IsProperty "equal"}} // {{.description}}
{
	hm{ {{.clock1.hour}}, {{.clock1.minute}}},
	hm{ {{.clock2.hour}}, {{.clock2.minute}}},
	{{.expected}},
}, {{- end}}{{end}}{{end}}
}
`
