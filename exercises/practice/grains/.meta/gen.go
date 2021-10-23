package main

import (
	"encoding/json"
	"log"
	"strings"
	"text/template"

	"../../../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("grains", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Groups []testGroup `json:"Cases"`
}

type testGroup struct {
	Description string
	Cases       []json.RawMessage `property:"RAW"`
	SquareCases []squareCase      `property:"square"`
	// Note: canonical-data.json has a element of "cases"
	// which includes a test for property 'total', but it is ignored here,
	// since "expected" is a single known value.
}

type squareCase struct {
	Description string
	Input       struct {
		Square int
	}
	Expected interface{}
}

func (c squareCase) HasAnswer() bool {
	hasAnswer, _ := determineExpected(c.Expected)
	return hasAnswer
}

func (c squareCase) Answer() uint64 {
	_, answer := determineExpected(c.Expected)
	return answer
}

func (c squareCase) EditedDescription() string {
	// Go doesn't raise exceptions, so replace the wording in .Description.
	return strings.Replace(c.Description, "raises an exception", "returns an error", 1)
}

// determineExpected examines an .Expected interface{} object and determines
// whether a test case has an answer or expects an error.
// The return values are true and answer or false and zero.
func determineExpected(expected interface{}) (bool, uint64) {
	ans, ok := expected.(float64)
	if ok {
		if ans == float64(-1) {
			return false, 0
		}
		return true, uint64(ans)
	}
	return false, 0
}

var tmpl = `package grains

{{.Header}}

{{range .J.Groups}}
	{{- if .SquareCases }}
		// {{ .Description }}
		var squareTests = []struct {
			description string
			input int
			expectedVal uint64
			expectError bool
		}{
			{{- range .SquareCases}}
			{
				description: "{{.EditedDescription}}",
				input: {{.Input.Square}},
				{{- if .HasAnswer}}
					expectedVal: {{.Answer}},
				{{- else}}
					expectError: true,
				{{- end}}
			},
			{{- end }}
		}
	{{- end }}
{{end}}
`
