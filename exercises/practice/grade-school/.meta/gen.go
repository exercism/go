package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"roster": &[]testCaseRoster{},
		"add":    &[]testCaseAdd{},
		"grade":  &[]testCaseGrade{},
	}
	if err := gen.Gen("grade-school", j, t); err != nil {
		log.Fatal(err)
	}
}

type Student [2]any

func (s Student) String() string {
	name, okN := s[0].(string)
	grade, okG := s[1].(float64)
	if !okN || !okG {
		log.Fatal("Could not parse student")
	}
	return fmt.Sprintf("student{%q, %d}", name, int(grade))
}

type testCaseRoster struct {
	Description string `json:"description"`
	Input       struct {
		Students []Student `json:"students"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

type testCaseAdd struct {
	Description string `json:"description"`
	Input       struct {
		Students []Student `json:"students"`
	} `json:"input"`
	Expected []bool `json:"expected"`
}

type testCaseGrade struct {
	Description string `json:"description"`
	Input       struct {
		Students     []Student `json:"students"`
		DesiredGrade int       `json:"desiredGrade"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

var tmpl = `{{.Header}}

type student struct {
	name string
	grade int
}

type testCaseAdd struct {
	description string
	students    []student
	expected    []bool
}

var testCasesAdd = []testCaseAdd { {{range .J.add}}
	{
		description: 	{{printf "%q" .Description}},
		students:       []student { {{range .Input.Students}}
			{{.}}, {{end}}
		},
		expected:    	{{printf "%#v" .Expected}},
	},{{end}}
}

type testCaseEnrollment struct {
	description string
	students    []student
	expected    []string
}

var testCasesEnrollment = []testCaseEnrollment { {{range .J.roster}}
	{
		description: 	{{printf "%q" .Description}},
		students:       []student { {{range .Input.Students}}
			{{.}}, {{end}}
		},
		expected:    	{{printf "%#v" .Expected}},
	},{{end}}
}

type testCaseGrade struct {
	description  string
	students     []student
	grade        int
	expected     []string
}

var testCasesGrade = []testCaseGrade { {{range .J.grade}}
	{
		description: 	{{printf "%q" .Description}},
		students:       []student { {{range .Input.Students}}
			{{.}}, {{end}}
		},
		grade:          {{.Input.DesiredGrade}},
		expected:    	{{printf "%#v" .Expected}},
	},{{end}}
}`
