package main

import (
	"log"
	"sort"
	"strconv"
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
	if err := gen.Gen("nucleotide-count", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	exercise string
	version  string
	Cases    []struct {
		Description string
		Cases       []OneCase
	}
}

// OneCase represents each test case
type OneCase struct {
	Description string
	Property    string
	Input       struct {
		Strand string
	}
	Expected map[string]interface{}
}

// ErrorExpected returns true if an error should be raised
func (c OneCase) ErrorExpected() bool {
	_, exists := c.Expected["error"]
	return exists
}

// SortedMapString collects key:values for a map in sorted order
func (c OneCase) SortedMapString() string {
	strs := make([]string, 0, len(c.Expected))
	for k, v := range c.Expected {
		switch t := v.(type) {
		case float64:
			strs = append(strs, `'`+k+`': `+strconv.FormatFloat(t, 'f', -1, 64))
		default:
			log.Fatalf("unexpected type %T for %v", t, v)
		}

	}
	sort.Strings(strs)
	return strings.Join(strs, ",")
}

// template applied to above data structure generates the Go test cases
var tmpl = `package dna

{{.Header}}

{{range .J.Cases}}// {{.Description}}
var testCases = []struct {
	description   string
	strand        string
	expected      Histogram
	errorExpected bool
}{
{{range .Cases}}{
	description:	{{printf "%q"  .Description}},
	strand:		{{printf "%#v"  .Input.Strand}},
	{{if .ErrorExpected}}errorExpected:	true,
	{{else}}expected:       Histogram{ {{.SortedMapString}} },
	{{- end}}
},
{{end}}{{end}}
}
`
