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
	j := map[string]any{
		"degreeOfSeparation": &[]testCase{},
	}
	if err := gen.Gen("relative-distance", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		FamilyTree map[string][]string `json:"familyTree"`
		PersonA    string              `json:"personA"`
		PersonB    string              `json:"personB"`
	} `json:"input"`
	Expected *int `json:"expected"`
}

func (tc testCase) ExpectedDegree() int {
	if tc.Expected == nil {
		return 0
	}
	return *tc.Expected
}

var tmpl = `{{.Header}}

var testCases = []struct {
	description    string
	tree           map[string][]string
	personA        string
	personB        string
	expectedDegree int
	expectedOk     bool
}{
{{range .J.degreeOfSeparation}}{
	description:    {{printf "%q" .Description}},
	tree:           {{printf "%#v" .Input.FamilyTree}},
	personA:        {{printf "%q" .Input.PersonA}},
	personB:        {{printf "%q" .Input.PersonB}},
	expectedDegree: {{.ExpectedDegree}},
	expectedOk:     {{ne nil .Expected}},
},
{{end}}
}`
