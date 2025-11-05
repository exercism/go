package main

import (
	"../../../../gen"
	"log"
	"strings"
	"text/template"
)

func main() {
	t := template.New("").Funcs(template.FuncMap{"fileContentData": FileContentData})
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"grep": &[]testCase{},
	}
	if err := gen.Gen("grep", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Pattern string   `json:"pattern"`
		Flags   []string `json:"flags"`
		Files   []string `json:"files"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

// FileContentData returns the lines containing the file contents.
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

var fileContentData = []string{ {{range $line := fileContentData .Comments}}{{printf "\n%q," $line}}{{end}}
}

var testCases = []struct {
	description string
	pattern     string
	flags       []string
	files       []string
	expected    []string
}{
{{range .J.grep}}{
		description: {{printf "%q" .Description}},
		pattern: {{printf "%q" .Input.Pattern}},
		flags: {{printf "%#v" .Input.Flags}},
		files: {{printf "%#v" .Input.Files}},
		expected: {{printf "%#v" .Expected}},
},
{{end}}
}
`
