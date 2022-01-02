package main

import (
	"log"
	"text/template"

	"../../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("matching-brackets", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			Value string
		}
		Expected bool
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package brackets

{{.Header}}

type bracketTest struct {
	input       string
	expected    bool
}

var testCases = []bracketTest {
{{range .J.Cases}}{
	// {{.Description}}
	{{printf "%q" .Input.Value}},
	{{.Expected}},
},
{{end}}}
`
