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
	j := map[string]interface{}{
		"encode": &[]testCase{},
		"decode": &[]testCase{},
	}
	if err := gen.Gen("affine-cipher", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Phrase string `json:"phrase"`
		Key    struct {
			Num1 int `json:"a"`
			Num2 int `json:"b"`
		} `json:"key"`
	} `json:"input"`
	Expected interface{} `json:"expected"`
}

func (t testCase) ExpectedString() string {
	m, ok := t.Expected.(string)
	if !ok {
		return ""
	}
	return m
}

func (t testCase) Error() bool {
	m, ok := t.Expected.(map[string]interface{})
	if !ok {
		return ok
	}
	_, ok = m["error"].(string)
	return ok
}

// Template to generate encode and decode test cases.
var tmpl = `package affinecipher

{{.Header}}

type testCase struct {
	description    string
	inputPhrase    string
	inputA          int
	inputB          int
	expectError     bool
	expected       string
}

var encodeTests = []testCase{
{{range .J.encode}}{
		description:         {{printf "%q" .Description}},
		inputPhrase:         {{printf "%q" .Input.Phrase}},
		inputA:              {{printf "%d" .Input.Key.Num1}},
		inputB:			     {{printf "%d" .Input.Key.Num2}},
		expectError:         {{printf "%t" .Error}},
		expected:            {{printf "%q" .ExpectedString}},
	},
{{end}}
}

var decodeTests = []testCase{
{{range .J.decode}}{
		description:         {{printf "%q" .Description}},
		inputPhrase:         {{printf "%q" .Input.Phrase}},
		inputA:              {{printf "%d" .Input.Key.Num1}},
		inputB:			     {{printf "%d" .Input.Key.Num2}},
		expectError:         {{printf "%t" .Error}},
		expected:            {{printf "%q" .ExpectedString}},
	},
{{end}}
}
`
