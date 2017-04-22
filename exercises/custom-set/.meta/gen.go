// +build ignore

package main

import (
	"errors"
	"log"
	"strconv"
	"text/template"

	"../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"str":   str,
		"strs":  strSlice,
		"istrs": istrSlice,
		"dict":  dict,
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

// str converts an integer 1..26 to a letter 'a'..'z'.
func str(n int) string {
	if n >= 1 && n <= 26 {
		return string('a' - 1 + n)
	}
	return strconv.Itoa(n)
}

// strSlice converts a slice of int to a slice of string.
func strSlice(ns []int) []string {
	s := make([]string, len(ns))
	for i, n := range ns {
		s[i] = str(n)
	}
	return s
}

// istrSlice converts a slice of interface{} values whose
// underlying members should be float64(JSON decoded integers)
// to a slice of string.
func istrSlice(ns []interface{}) []string {
	s := make([]string, len(ns))
	for i, n := range ns {
		if fn, ok := n.(float64); ok {
			s[i] = str(int(fn))
		}
	}
	return s
}

// dict sets up and returns a map for pairs of items given;
// it is used in order to pass multiple parameters in a template call.
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call: odd number of values given")
	}
	valueMap := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("Expected string for dict key")
		}
		valueMap[key] = values[i+1]
	}
	return valueMap, nil
}

// The JSON structure we expect to be able to unmarshal into
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
	Set         []int       // "empty"/"contains"/"add" cases
	Set1        []int       // "empty"/"contains"/"add" cases
	Set2        []int       // "subset"/"disjoint"/"equal"/"difference"/"intersection"/"union" cases
	Element     int         // "contains"/"add" cases
	Expected    interface{} // bool or []int
}

func (c OneCase) PropertyMatch(property string) bool { return c.Property == property }

// GroupComment looks in each of the test case groups to find the
// group for which every test case has the .Property matching given property;
// it returns the .Description field for the matching property group,
// or a 'Note: ...' if no test group consistently matches given property.
func (groups TestGroups) GroupComment(property string) string {
	for _, group := range groups {
		propertyGroupMatch := true
		for _, testcase := range group.Cases {
			if !testcase.PropertyMatch(property) {
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

// template applied to above data structure generates the Go test cases
//
// There's some extra complexity because json test cases use ints but it's
// all converted to strings here.  It just seemed like strings would make
// a better and more practical example.
var tmpl = `
{{/* nested templates for repeated stuff */}}

{{define "unaryBool"}}{{$property := .PropertyType}}{{with .Groups}}
// {{ .GroupComment $property}}
var {{$property}}Cases = []unaryBoolCase{ {{range .}} {{range .Cases}}
{{if .PropertyMatch $property}} { // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{.Expected}},
},
{{- end}}{{end}}{{end}}
}
{{end}}{{end}}

{{define "eleBool"}}{{$property := .PropertyType}}{{with .Groups}}
// {{ .GroupComment $property}}
var {{$property}}Cases = []eleBoolCase{ {{range .}} {{range .Cases}}
{{if .PropertyMatch $property}} { // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{.Expected}},
},
{{- end}}{{end}}{{end}}
}
{{end}}{{end}}

{{define "eleOp"}}{{$property := .PropertyType}}{{with .Groups}}
// {{ .GroupComment $property}}
var {{$property}}Cases = []eleOpCase{ {{range .}} {{range .Cases}}
{{if .PropertyMatch $property}} { // {{.Description}}
	{{strs .Set | printf "%#v"}},
	{{str .Element | printf "%q"}},
	{{istrs .Expected | printf "%#v"}},
},
{{- end}}{{end}}{{end}}
}
{{end}}{{end}}

{{define "binaryBool"}}{{$property := .PropertyType}}{{with .Groups}}
// {{ .GroupComment $property}}
var {{$property}}Cases = []binBoolCase{ {{range .}} {{range .Cases}}
{{if .PropertyMatch $property}} { // {{.Description}}
	{{strs .Set1 | printf "%#v"}},
	{{strs .Set2 | printf "%#v"}},
	{{.Expected}},
},
{{- end}}{{end}}{{end}}
}
{{end}}{{end}}

{{define "binaryOp"}}{{$property := .PropertyType}}{{with .Groups}}
// {{ .GroupComment $property}}
var {{$property}}Cases = []binOpCase{ {{range .}} {{range .Cases}}
{{if .PropertyMatch $property}} { // {{.Description}}
	{{strs .Set1 | printf "%#v"}},
	{{strs .Set2 | printf "%#v"}},
	{{istrs .Expected | printf "%#v"}},
},
{{- end}}{{end}}{{end}}
}
{{end}}{{end}}

{{/* begin template body */}}
package stringset

{{.Header}}

{{/* These template calls utilize a dict helper function in order to pass multiple parameters. */}}
{{template "unaryBool"  dict "PropertyType" "empty"        "Groups" .J.Groups}}
{{template "eleBool"    dict "PropertyType" "contains"     "Groups" .J.Groups}}
{{template "binaryBool" dict "PropertyType" "subset"       "Groups" .J.Groups}}
{{template "binaryBool" dict "PropertyType" "disjoint"     "Groups" .J.Groups}}
{{template "binaryBool" dict "PropertyType" "equal"        "Groups" .J.Groups}}
{{template "eleOp"      dict "PropertyType" "add"          "Groups" .J.Groups}}
{{template "binaryOp"   dict "PropertyType" "intersection" "Groups" .J.Groups}}
{{template "binaryOp"   dict "PropertyType" "difference"   "Groups" .J.Groups}}
{{template "binaryOp"   dict "PropertyType" "union"        "Groups" .J.Groups}}
`
