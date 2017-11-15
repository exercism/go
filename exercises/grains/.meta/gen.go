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
	SquareCases []SquareCase      `property:"square"`
	// Note: canonical-data.json has a element of "cases"
	// which includes a test for property 'total', but it is ignored here,
	// since "expected" is a single known value.
}

type SquareCase struct {
	Description string
	Input       int
	Expected    interface{}
}

func (c SquareCase) Valid() bool {
	valid, _ := determineExpected(c.Expected)
	return valid
}

func (c SquareCase) Answer() uint64 {
	_, answer := determineExpected(c.Expected)
	return answer
}

func (c SquareCase) EditedDescription() string {
	// Go doesn't raise exceptions, so replace the wording in .Description.
	return strings.Replace(c.Description, "raises an exception", "returns an error", 1)
}

// determineExpected examines an .Expected interface{} object and determines
// whether a test case is valid(bool) and has an answer or expects an error.
// returning valid and answer.
func determineExpected(expected interface{}) (bool, uint64) {
	ans, ok := expected.(float64)
	if ok {
		if ans == float64(-1) {
			return false, 0
		}
		return ok, uint64(ans)
	}
	return false, 0
}

func (groups testGroup) GroupShortName() string {
	return strings.ToLower(strings.Fields(groups.Description)[0])
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
				input: {{.Input}},
				{{- if .Valid}}
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
