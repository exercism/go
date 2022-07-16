package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
	"time"
)

const (
	// canonicalDataURL is the URL for the raw canonical-data.json data and requires the exercise name.
	canonicalDataURL = "https://raw.githubusercontent.com/exercism/problem-specifications/master/exercises/%s/canonical-data.json"
	// commitsURL is the GitHub api endpoint for the commit history of canonical-data.json and requires the exercise name.
	commitsURL = "https://api.github.com/repos/exercism/problem-specifications/commits?path=exercises/%s/canonical-data.json"
	// defaultOrigin is the origin used in the header of cases_test.go
	defaultOrigin = "exercism/problem-specifications"
)

// problemSpecificationsDir is the location of the problem-specifications repository on the filesystem.
// We're making the assumption that the problem-specifications repository has been cloned to
// the same parent directory as the Exercism Go track repository.
// E.g.
//
//     $ tree -L 1 .
//     .
//     ├── problem-specifications
//     └── go
//
var problemSpecificationsDir string

// exerciseDir is the location of the exercise and also the cases_test.go file.
// This assumes that the generator script lives in the .meta directory within
// the exercise directory. Falls back to the present working directory.
var exerciseDir string

// httpClient creates a http client with a timeout of 10 seconds in order not to get stuck waiting for a response.
var httpClient = &http.Client{Timeout: 10 * time.Second}

// Header tells how the test data was generated, for display in the header of cases_test.go
type Header struct {
	Commit string
	Origin string
}

// String generates the header for cases_test.go file.
func (h Header) String() string {
	return fmt.Sprintf("// Source: %s\n// Commit: %s", h.Origin, h.Commit)
}

type testCase struct {
	UUID        string      `json:"uuid"`
	Description string      `json:"description"`
	Comments    []string 	`json:"comments"`
	Property    string      `json:"property"`
	Scenario    string      `json:"scenario"`
	Input       interface{} `json:"input"`
	Expected    interface{} `json:"expected"`
}

// Gen generates the exercise cases_test.go file from the relevant canonical-data.json
func Gen(exercise string, tests map[string]interface{}, t *template.Template) error {
	// Determine the exercise directory.
	// Use runtime.Caller to determine path location of the generator main package.
	// Call frames: 0 is this frame, Gen().
	//              1 should be a func in <exercise-dir>/.meta/gen.go (generator main package)
	//                which is calling Gen().
	//                                  |
	//                                  V
	if _, path, _, ok := runtime.Caller(1); ok {
		// Construct a path 2 directories higher than the path for .meta/gen.go (should be exercise directory).
		exerciseDir = filepath.Join(path, "..", "..")
	} else {
		return fmt.Errorf("unable to determine directory of exercise %s", exercise)
	}

	if exerciseDir == "" {
		exerciseDir = "."
	}

	problemSpecificationsDir = filepath.Join(exerciseDir, "..", "..", "..", "..", "problem-specifications")

	jFile := filepath.Join("exercises", exercise, "canonical-data.json")
	fmt.Printf("[LOCAL] fetching %s test data from canonical-data.json\n", exercise)

	var header Header
	jTestData, err := getLocalTestData(jFile)
	if err == nil {
		header, err = getLatestLocalCommitMessage(jFile)
	}

	if err != nil {
		// fetch json data remotely if there's no local file
		fmt.Printf("[LOCAL] No test data found: %v\n", err)
		fmt.Printf("[REMOTE] fetching %s test data\n", exercise)
		jTestData, err = getRemoteTestData(exercise)
		if err != nil {
			return err
		}
		header, err = getRemoteCommit(exercise)
		if err != nil {
			return err
		}
	}

	if !json.Valid(jTestData) {
		return fmt.Errorf("[ERROR] canonical-data.json seems not to be valid json")
	}

	// read tests.toml file to find which test cases should be excluded
	tomlFile := filepath.Join(exerciseDir, ".meta", "tests.toml")
	fmt.Printf("[LOCAL] reading tests.toml file from exercise directory %s\n", tomlFile)
	excludedTests, err := getExcludedTestCases(tomlFile)
	if err != nil {
		return fmt.Errorf("[LOCAL] unable to read tests.toml file (%v)", err)
	}

	fmt.Println("collecting and filtering all test cases from the fetched test data")

	allTestCases, err := getAllTestCasesFiltered(jTestData, excludedTests)
	if err != nil {
		return fmt.Errorf("failed to get filtered test-cases: %w", err)
	}

	var casesPerProperty = map[string][]testCase{}

	for _, testCase := range *allTestCases {
		casesPerProperty[testCase.Property] = append(casesPerProperty[testCase.Property], testCase)
	}

	for property, testCases := range casesPerProperty {
		fmt.Printf(" > parsing cases for property %s\n", property)
		marshal, err := json.Marshal(testCases)
		if err != nil {
			return fmt.Errorf("[ERROR] failed to marshal test cases with property %s: %w", property, err)
		}

		// valueForProperty is an instance of the struct (defined in gen.go) for the current property
		// and is used to unmarshal the test cases
		valueForProperty, ok := tests[property]
		if !ok {
			return fmt.Errorf("[ERROR] failed to get struct for tests with property %s", property)
		}
		err = json.Unmarshal(marshal, valueForProperty)
		if err != nil {
			return err
		}
	}

	// package up a little meta data
	d := struct {
		Header
		J map[string]interface{}
	}{Header: header, J: tests}

	casesFile := filepath.Join(exerciseDir, "cases_test.go")

	// render the Go test cases
	var out bytes.Buffer
	if err := t.Execute(&out, &d); err != nil {
		return fmt.Errorf("[ERROR] template.Execute failed. The template has a semantic error: %w", err)
	}

	casesFileContent := out.Bytes()
	formattedFileContent, err := format.Source(casesFileContent)
	if err != nil {
		fmt.Print("[ERROR] failed to format the output with gofmt (the generated source has a syntax error)")
		debugFileContent := append(casesFileContent, []byte("// !NOTE: Error during source formatting: Line:Column "+fmt.Sprint(err)+"\n")...)
		// Save the raw unformatted, error-containing source for purposes of debugging the generator.
		_ = outputSource("ERROR", casesFile, debugFileContent)
		return err
	}
	// write output file for the Go test cases.
	return outputSource("SUCCESS", casesFile, formattedFileContent)
}

// outputSource writes the src text to the given fileName and outputs a log message with given [status].
func outputSource(status, fileName string, src []byte) error {
	err := os.WriteFile(fileName, src, 0666)
	if err != nil {
		return fmt.Errorf("[FAILED] %q\n", err)
	}
	fmt.Printf("[%s] output: %s\n", status, fileName)
	return nil
}
