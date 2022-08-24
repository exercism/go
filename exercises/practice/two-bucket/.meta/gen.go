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
		"measure": &[]testCase{},
	}
	if err := gen.Gen("two-bucket", j, t); err != nil {
		log.Fatal(err)
	}
}

type testCase struct {
	Description string `json:"description"`
	Input       struct {
		BucketOne   int    `json:"bucketOne"`
		BucketTwo   int    `json:"bucketTwo"`
		Goal        int    `json:"goal"`
		StartBucket string `json:"startBucket"`
	} `json:"input"`
	Expected struct {
		Moves       int    `json:"moves"`
		GoalBucket  string `json:"goalBucket"`
		OtherBucket int    `json:"otherBucket"`
		Error       string `json:"error"`
	} `json:"expected"`
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
	expectedError   string
}

var testCases = []bucketTestCase {
{{range .J.measure}}{
		description:   	{{ printf "%q" .Description}},
		bucketOne:     	{{ printf "%d" .Input.BucketOne}},
		bucketTwo:      {{ printf "%d" .Input.BucketTwo}},
		goal:           {{ printf "%d" .Input.Goal}},
		startBucket:   	{{ printf "%q" .Input.StartBucket}},
		goalBucket:    	{{ printf "%q" .Expected.GoalBucket}},
		moves:        	{{ printf "%d" .Expected.Moves}},
		otherBucket:   	{{ printf "%d" .Expected.OtherBucket}},
		expectedError: 	{{ printf "%q" .Expected.Error}},
},
{{end}}}
`
