// +build ignore

package main

import (
	"log"
	"strconv"
	"text/template"

	"../gen"
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
	if err := gen.Gen("custom-set.json", &j, t); err != nil {
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
	Equal               binaryBool // binary function, boolean result
	Add                 eleOp      // set-element operator
	Delete              eleOp      // ...
	IsEmpty             unaryBool  `json:"is-empty"`
	Size                unaryInt   // Set.Len
	Element             eleBool    // Set.Has
	Subset              binaryBool
	Disjoint            binaryBool
	Union               binaryOp // set-set operator
	Intersection        binaryOp
	Difference          binaryOp
	SymmetricDifference binaryOp `json:"symmetric-difference"`
}

type binaryBool struct {
	Description []string
	Cases       []struct {
		Description string
		Set1        []int
		Set2        []int
		Expected    bool
	}
}

type eleOp struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Element     int
		Expected    []int
	}
}

type unaryBool struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Expected    bool
	}
}

type unaryInt struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Expected    int
	}
}

type eleBool struct {
	Description []string
	Cases       []struct {
		Description string
		Set         []int
		Element     int
		Expected    bool
	}
}

type binaryOp struct {
	Description []string
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

{{define "cmts"}}{{range .}}// {{.}}
{{end}}{{end}}

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

{{template "cmts" .J.Equal.Description}}var eqCases{{template "binaryBool" .J.Equal.Cases}}

{{template "cmts" .J.Add.Description}}var addCases{{template "eleOp" .J.Add.Cases}}

{{template "cmts" .J.Delete.Description}}var delCases{{template "eleOp" .J.Delete.Cases}}

{{range .J.IsEmpty.Description}}// {{.}}
{{end}}var emptyCases = []unaryBoolCase{
{{range .J.IsEmpty.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{.Expected}},
},
{{end}}}

{{range .J.Size.Description}}// {{.}}
{{end}}var lenCases = []unaryIntCase{
{{range .J.Size.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{.Expected}},
},
{{end}}}

{{range .J.Element.Description}}// {{.}}
{{end}}var hasCases = []eleBoolCase{
{{range .J.Element.Cases}}{ // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{.Expected}},
},
{{end}}}

{{template "cmts" .J.Subset.Description}}var subsetCases{{template "binaryBool" .J.Subset.Cases}}

{{template "cmts" .J.Disjoint.Description}}var disjointCases{{template "binaryBool" .J.Disjoint.Cases}}

{{template "cmts" .J.Union.Description}}var unionCases{{template "binaryOp" .J.Union.Cases}}

{{template "cmts" .J.Intersection.Description}}var intersectionCases{{template "binaryOp" .J.Intersection.Cases}}

{{template "cmts" .J.Difference.Description}}var differenceCases{{template "binaryOp" .J.Difference.Cases}}

{{template "cmts" .J.SymmetricDifference.Description}}var symmetricDifferenceCases{{template "binaryOp" .J.SymmetricDifference.Cases}}
`
