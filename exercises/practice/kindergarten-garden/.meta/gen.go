package main

import (
	"../../../../gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	j := map[string]any{
		"plants": &[]testCase{},
	}
	if err := gen.Gen("kindergarten-garden", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		Diagram string `json:"diagram"`
		Student string `json:"student"`
	} `json:"input"`
	Expected []string `json:"expected"`
}

func (tc testCase) Children() string {
	children := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"}
	count := len(strings.Split(tc.Input.Diagram, "\n")[1]) / 2
	if count == 12 {
		return "nil"
	}
	return fmt.Sprintf("%#v", children[:count])
}

var tmpl = `{{.Header}}

type lookup struct {
	child  string
	plants []string
	ok     bool
}

type gardenTest struct {
	description string
	diagram     string
	children    []string
	expectError bool
	lookups     []lookup
}

var testCases = []gardenTest { {{range .J.plants}}
	{
		description: 	{{printf "%q" .Description}},
		diagram:        "\n" + {{printf "%q" .Input.Diagram}},
		children:       {{.Children}},
		expectError:    false,
		lookups:        []lookup{
			{
				child: {{printf "%q" .Input.Student}},
				plants: {{printf "%#v" .Expected}},
				ok: true,
			},
		},
	},{{end}}
}`
