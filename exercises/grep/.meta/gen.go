package main

import (
	"log"
	"strings"
	"text/template"

	"../../../gen"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{
		"fileContentData": FileContentData,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("grep", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Comments []string
	Cases    []struct {
		Description string
		Cases       []struct {
			Description string
			Input       struct {
				Pattern string
				Flags   []string
				Files   []string
			}
			Expected []string
		}
	}
}

func FileContentData(comments []string) []string {
	var start int
	for i, c := range comments {
		if strings.Contains(c, "contents are listed below:") {
			start = i + 1
			break
		}
	}
	return comments[start:]
}

// template applied to above data structure generates the Go test cases
var tmpl = `package grep

{{.Header}}

var fileContentData = []string{ {{range $line := fileContentData .J.Comments}}{{printf "\n%q," $line}}{{end}}
}

var testCases = []struct {
	description string
	pattern     string
	flags       []string
	files       []string
	expected    []string
}{
{{range .J.Cases}}
	{{range .Cases}}{
		description: {{printf "%q" .Description}},
		pattern: {{printf "%q" .Input.Pattern}},
		flags: {{printf "%#v" .Input.Flags}},
		files: {{printf "%#v" .Input.Files}},
		expected: {{printf "%#v" .Expected}},
},
{{end}}{{end}}
}
`
