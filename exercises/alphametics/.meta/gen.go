package main

import (
	"fmt"
	"log"
	"sort"
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
	letters := make([]string, 0, len(c.Expected))
	for s, _ := range c.Expected {
		letters = append(letters, s)
	}
	result := ""
	sort.Strings(letters)
	for _, s := range letters {
		result += fmt.Sprintf(`"%s": %d, `, s, c.Expected[s])
	}
	return result
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
