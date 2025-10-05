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
		"encode": &[]testCase{},
		"decode": &[]testCase{},
	}
	if err := gen.Gen("rail-fence-cipher", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Msg   string `json:"msg"`
		Rails int    `json:"rails"`
	} `json:"input"`
	Expected string `json:"expected"`
}

// Template to generate test cases.
var tmpl = `package railfence

{{.Header}}

type testCase struct {
	description string
	message     string
	rails       int
	expected    string
}

var encodeTests = []testCase{
{{range .J.encode}}{
		description: {{printf "%q" .Description}},
		message:     {{printf "%q" .Input.Msg}},
		rails:       {{printf "%d" .Input.Rails}},
		expected:    {{printf "%q" .Expected}},
	},
{{end}}
}

var decodeTests = []testCase{
{{range .J.decode}}{
		description: {{printf "%q" .Description}},
		message:     {{printf "%q" .Input.Msg}},
		rails:       {{printf "%d" .Input.Rails}},
		expected:    {{printf "%q" .Expected}},
	},
{{end}}
}
`
