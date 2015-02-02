// +build ignore

package main

import (
	"log"
	"text/template"

	"../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("clock.json", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to umarshal into
type js struct {
	Create struct {
		Description []string
		Cases       []timeCase
	}
	Add struct {
		Description []string
		Cases       []addCase
	}
	Equal struct {
		Description []string
		Cases       []eqCase
	}
}

// Handle the three tests similarly

type timeCase struct {
	Description  string
	Hour, Minute int
	Expected     string
}
type addCase struct {
	Description       string
	Hour, Minute, Add int
	Expected          string
}
type eqCase struct {
	Description    string
	Clock1, Clock2 struct{ Hour, Minute int }
	Expected       bool
}

var tmpl = `package clock

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

{{range .J.Create.Description}}// {{.}}
{{end}}var timeTests = []struct {
	h, m int
	want string
}{
{{range .J.Create.Cases}}{ {{.Hour}}, {{.Minute}}, {{.Expected | printf "%#v"}}},	// {{.Description}}
{{end}}}

{{range .J.Add.Description}}// {{.}}
{{end}}var addTests = []struct {
	h, m, a int
	want  string
}{
{{range .J.Add.Cases}}{ {{.Hour}}, {{.Minute}}, {{.Add}}, {{.Expected | printf "%#v"}}}, // {{.Description}}
{{end}}}

{{range .J.Equal.Description}}// {{.}}
{{end}}type hm struct{ h, m int }

var eqTests = []struct {
	c1, c2 hm
	want   bool
}{
{{range .J.Equal.Cases}}// {{.Description}}
{
	hm{ {{.Clock1.Hour}}, {{.Clock1.Minute}}},
	hm{ {{.Clock2.Hour}}, {{.Clock2.Minute}}},
	{{.Expected}},
},
{{end}}}`
