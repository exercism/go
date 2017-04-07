// +build ignore

package main

import (
	"log"
	"strconv"
	"text/template"

	"../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"str":  str,
		"strs": strSlice,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("custom-set", &j, t); err != nil {
		log.Fatal(err)
	}
}

func str(n int) string {
	if n >= 1 && n <= 26 {
		return string('a' - 1 + n)
	}
	return strconv.Itoa(n)
}

func strSlice(ns []int) []string {
	s := make([]string, len(ns))
	for i, n := range ns {
		s[i] = str(n)
	}
	return s
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	IsEmpty      unaryBool `json:"empty"`
	Contains     eleBool
	Subset       binaryBool
	Disjoint     binaryBool
	Equal        binaryBool
	Add          eleOp
	Intersection binaryOp
	Difference   binaryOp
	Union        binaryOp
}

type unaryBool struct {
	Description string
	Cases       []struct {
		Description string
		Set         []int
		Expected    bool
	}
}

type eleBool struct {
	Description string
	Cases       []struct {
		Description string
		Set         []int
		Element     int
		Expected    bool
	}
}

type binaryBool struct {
	Description string
	Cases       []struct {
		Description string
		Set1        []int
		Set2        []int
		Expected    bool
	}
}

type eleOp struct {
	Description string
	Cases       []struct {
		Description string
		Set         []int
		Element     int
		Expected    []int
	}
}

type binaryOp struct {
	Description string
	Cases       []struct {
		Description string
		Set1        []int
		Set2        []int
		Expected    []int
	}
}

// template applied to above data structure generates the Go test cases
//
// There's some extra complexity because json test cases use ints but it's
// all converted to strings here.  It just seemed like strings would make
// a better and more practical example.
var tmpl = `
{{/* nested templates for repeated stuff */}}

{{define "binaryBool"}} = []binBoolCase{
{{range .}}{ // {{.Description}}
	{{strs .Set1 | printf "%#v"}},
	{{strs .Set2 | printf "%#v"}},
	{{.Expected}},
},
{{end}}}{{end}}

{{define "binaryOp"}} = []binOpCase{
{{range .}}{ // {{.Description}}
	{{strs .Set1 | printf "%#v"}},
	{{strs .Set2 | printf "%#v"}},
	{{strs .Expected | printf "%#v"}},
},
{{end}}}{{end}}

{{define "eleOp"}} = []eleOpCase{
{{range .}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{strs .Expected | printf "%#v"}},
},
{{end}}}{{end}}

{{/* begin template body */}}
package stringset

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

// {{.J.IsEmpty.Description}}
var emptyCases = []unaryBoolCase{
{{range .J.IsEmpty.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{.Expected}},
},
{{end}}}

// {{.J.Contains.Description}}
var hasCases = []eleBoolCase{
{{range .J.Contains.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{.Expected}},
},
{{end}}}

// {{.J.Subset.Description}}
var subsetCases{{template "binaryBool" .J.Subset.Cases}}

// {{.J.Disjoint.Description}}
var disjointCases{{template "binaryBool" .J.Disjoint.Cases}}

// {{.J.Equal.Description}}
var eqCases{{template "binaryBool" .J.Equal.Cases}}

// {{.J.Add.Description}}
var addCases{{template "eleOp" .J.Add.Cases}}

// {{.J.Intersection.Description}}
var intersectionCases{{template "binaryOp" .J.Intersection.Cases}}

// {{.J.Difference.Description}}
var differenceCases{{template "binaryOp" .J.Difference.Cases}}

// {{.J.Union.Description}}
var unionCases{{template "binaryOp" .J.Union.Cases}}
`
