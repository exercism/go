package main

import (
	"../../../../gen"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]interface{}{
		"distance": &[]testCase{},
	}
	if err := gen.Gen("hamming", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Strand1 string `json:"strand1"`
		Strand2 string `json:"strand2"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectedValue() int {
	v, ok := t.Expected.(float64)
	if !ok {
		return 0
	}
	return int(v)
}

func (t testCase) ExpectError() bool {
	_, ok := t.Expected.(float64)
	return !ok
}

// template applied to above data structure generates the Go test cases
var tmpl = `package hamming

{{.Header}}

var testCases = []struct {
	description string
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{{range .J.distance}}{
			description: {{printf "%q"  .Description}},
			s1:          {{printf "%q"  .Input.Strand1}},
			s2:          {{printf "%q"  .Input.Strand2}},
			want:        {{printf "%d"  .ExpectedValue}},
			expectError: {{printf "%t"  .ExpectError}},
		},
	{{end}}
}
`
