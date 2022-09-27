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
		"create":   &[]defaultTestCase{},
		"add":      &[]defaultTestCase{},
		"subtract": &[]defaultTestCase{},
		"equal":    &[]equalTestCase{},
	}
	if err := gen.Gen("clock", j, t); err != nil {
		log.Fatal(err)
	}
}

type defaultTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Hour   int `json:"hour"`
		Minute int `json:"minute"`
		Value  int `json:"value"`
	} `json:"input"`
	Expected string `json:"expected"`
}

type equalTestCase struct {
	Description string `json:"description"`
	Input       struct {
		Clock1 struct {
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
		} `json:"clock1"`
		Clock2 struct {
			Hour   int `json:"hour"`
			Minute int `json:"minute"`
		} `json:"clock2"`
	} `json:"input"`
	Expected bool `json:"expected"`
}

var tmpl = `package clock

{{.Header}}

var timeTestCases = []struct {
	description string
	h, m        int
	expected    string
}{
{{range .J.create}}{
			description: {{printf "%q"  .Description}},
			h:           {{printf "%d"  .Input.Hour}},
			m:           {{printf "%d"  .Input.Minute}},
			expected:    {{printf "%q"  .Expected}},
	},
{{end}}
}

var addTestCases = []struct {
	description      string
	h, m, addedValue int
	expected         string
}{
{{range .J.add}}{
		description: {{printf "%q"  .Description}},
		h:           {{printf "%d"  .Input.Hour}},
		m:           {{printf "%d"  .Input.Minute}},
		addedValue:  {{printf "%d"  .Input.Value}},
		expected:    {{printf "%q"  .Expected}},
	},
{{end}}
}

var subtractTestCases = []struct {
	description      		string
	h, m, subtractedValue 	int
	expected         		string
}{
{{range .J.subtract}}{
		description: 		{{printf "%q"  .Description}},
		h:           		{{printf "%d"  .Input.Hour}},
		m:           		{{printf "%d"  .Input.Minute}},
		subtractedValue:  	{{printf "%d"  .Input.Value}},
		expected:    		{{printf "%q"  .Expected}},
	},
{{end}}
}

// Compare two clocks for equality
type hm struct{ h, m int }

var equalTestCases = []struct {
	description string
	c1, c2      hm
	expected    bool
}{ {{range .J.equal}}{
		description: {{printf "%q"  .Description}},
		c1:          hm{ {{printf "%d"  .Input.Clock1.Hour}} , {{printf "%d"  .Input.Clock1.Minute}} },
		c2:          hm{ {{printf "%d"  .Input.Clock2.Hour}} , {{printf "%d"  .Input.Clock2.Minute}} },
		expected:    {{printf "%v"  .Expected}},
	},
{{end}}
}
`
