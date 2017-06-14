package main

import (
	"log"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("secret-handshake", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Cases       []struct {
			Description string
			Property    string
			Input       uint
			Expected    []string
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package secret

{{.Header}}

type secretHandshakeTest struct {
	code uint
	h []string
}

var tests = []secretHandshakeTest {
{{range .J.Cases}} {{range .Cases}}{ {{ .Input }}, {{printf "%#v" .Expected }},
},
{{end}}{{end}}
}
`
