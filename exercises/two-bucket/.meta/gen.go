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
	if err := gen.Gen("two-bucket", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       struct {
			BucketOne   int
			BucketTwo   int
			Goal        int
			StartBucket string
		}
		Expected struct {
			GoalBucket  string
			Moves       int
			OtherBucket int
		}
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package twobucket

{{.Header}}

type bucketTestCase struct {
	description     string
	bucketOne	int
	bucketTwo	int
	goal		int
	startBucket	string
	goalBucket	string
	moves		int
	otherBucket	int
	errorExpected   bool // always false for generated test cases.
}

var testCases = []bucketTestCase {
{{range .J.Cases}}{
	"{{.Description}}",
	{{.Input.BucketOne}}, {{.Input.BucketTwo}}, {{.Input.Goal}}, "{{.Input.StartBucket}}", "{{.Expected.GoalBucket}}", {{.Expected.Moves}}, {{.Expected.OtherBucket}}, false,
},
{{end}}}
`
