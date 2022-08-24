package main

import (
	"log"
	"sort"
	"strings"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"solve": &[]testCase{},
	}
	if err := gen.Gen("alphametics", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Puzzle string `json:"puzzle"`
	} `json:"input"`
	Expected map[string]int `json:"expected"`
}

func (c testCase) ErrorExpected() bool {
	return len(c.Expected) == 0
}

func (c testCase) SortedMapString() string {
	strs := make([]string, 0, len(c.Expected))
	for s, v := range c.Expected {
		strs = append(strs, `"`+s+`": `+string(rune(v+'0')))
	}
	sort.Strings(strs)
	return strings.Join(strs, ",")
}

// template applied to above data structure generates the Go test cases
var tmpl = `package alphametics

{{.Header}}

var testCases = []struct {
	description   string
	input         string
	expected      map[string]int
	errorExpected bool
}{
{{range .J.solve}}{
	description:	{{printf "%q" .Description}},
	input:		{{printf "%q" .Input.Puzzle}},
	{{if .ErrorExpected}}errorExpected:	true,
	{{else}}expected:	map[string]int{ {{.SortedMapString}} },
	{{- end}}
},
{{end}}
}
`
