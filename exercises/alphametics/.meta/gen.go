package main

import (
	"log"
	"sort"
	"strings"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("alphametics", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []OneCase
	}
}
type OneCase struct {
	Description string
	Input       struct {
		Puzzle string
	}
	Expected map[string]int
}

func (c OneCase) ErrorExpected() bool {
	return len(c.Expected) == 0
}

func (c OneCase) SortedMapString() string {
	strs := make([]string, 0, len(c.Expected))
	for s, v := range c.Expected {
		strs = append(strs, `"`+s+`": `+string(v+'0'))
	}
	sort.Strings(strs)
	return strings.Join(strs, ",")
}

// template applied to above data structure generates the Go test cases
var tmpl = `package alphametics

{{.Header}}

{{range .J.Cases}}// {{.Description}}
var testCases = []struct {
	description   string
	input         string
	expected      map[string]int
	errorExpected bool
}{
{{range .Cases}}{
	description:	{{printf "%q" .Description}},
	input:		{{printf "%q" .Input.Puzzle}},
	{{if .ErrorExpected}}errorExpected:	true,
	{{else}}expected:	map[string]int{ {{.SortedMapString}} },
	{{- end}}
},
{{end}}{{end}}
}
`
