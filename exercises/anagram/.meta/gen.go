package main

import (
	"text/template"

	"github.com/exercism/xgo/gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		panic(err)
	}
	if err = gen.Gen("anagram", &js{}, t); err != nil {
		panic(err)
	}
}

type js struct {
	Cases []struct {
		Description string
		Property    string
		Subject     string
		Candidates  []string
		Expected    []string
	}
}

const tmpl = `package anagram

{{.Header}}

type testcase struct{
  subject string
  candidates []string
  expected []string
  description string
}

var testCases = []testcase{
  {{range .J.Cases}}
	{
	  subject: {{printf "%q" .Subject}},
	  candidates: {{printf "%#v" .Candidates}},
      expected: {{printf "%#v" .Expected}},
	  description: {{printf "%q" .Description}},
	},
  {{end}}
}
`
