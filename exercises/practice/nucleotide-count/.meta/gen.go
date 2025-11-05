package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"sort"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"nucleotideCounts": &[]testCase{},
	}
	if err := gen.Gen("nucleotide-count", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Strand string `json:"strand"`
	} `json:"input"`
	Expected map[string]any `json:"expected"`
}

func (t testCase) ErrorExpected() bool {
	_, gotError := t.Expected["error"]
	return gotError
}

// SortedMapString collects key:values for a map in sorted order
func (t testCase) SortedMapString() string {
	if t.ErrorExpected() {
		return ""
	}
	strs := make([]string, 0, len(t.Expected))
	for k, v := range t.Expected {
		switch t := v.(type) {
		case float64:
			strs = append(strs, fmt.Sprintf("'%s':%d", k, int(t)))
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

var testCases = []struct {
	description   string
	strand        string
	expected      Histogram
	errorExpected bool
}{
{{range .J.nucleotideCounts}}{
	description:	{{printf "%q"  .Description}},
	strand:			{{printf "%#v"  .Input.Strand}},
	errorExpected:	{{printf "%t"  .ErrorExpected}},
	expected:       Histogram{ {{.SortedMapString}} },
},
{{end}}
}
`
