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
		"add": &[]testCase{},
	}
	if err := gen.Gen("gigasecond", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Moment string `json:"moment"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// template applied to above data structure generates the Go test cases
var tmpl = `package gigasecond

{{.Header}}

var addCases = []struct {
	description string
	in   string
	want string
}{
{{range .J.add}}{
	description: {{printf "%q" .Description}},
	in: {{printf "%q" .Input.Moment}},
	want: {{printf "%q" .Expected}},
},
{{end}}
}
`
