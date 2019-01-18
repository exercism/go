package gen

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
	"time"
)

// dirMetadata is the location of the problem-specifications repository on the filesystem.
// We're making the assumption that the problem-specifications repository has been cloned to
// the same parent directory as the Exercism Go track repository.
// E.g.
//
//     $ tree -L 1 .
//     .
//     ├── problem-specifications
//     └── go
//
var dirMetadata string

// dirExercise is the location that the test cases should be generated to.
// This assumes that the generator script lives in the .meta directory within
// the exercise directory. Falls back to the present working directory.
var dirExercise string

// genClient creates an http client with a 10 second timeout so we don't get
// stuck waiting for a response.
var genClient = &http.Client{Timeout: 10 * time.Second}

const (
	// canonicalDataURL is the URL for the raw canonical-data.json data,
	// requires exercise name.
	canonicalDataURL = "https://raw.githubusercontent.com/exercism/problem-specifications/master/exercises/%s/canonical-data.json"
	// commitsURL is the GitHub api endpoint for the canonical-data.json
	// file commit history, requires exercise name.
	commitsURL = "https://api.github.com/repos/exercism/problem-specifications/commits?path=exercises/%s/canonical-data.json"
)

// Header tells how the test data was generated, for display in the header of
// cases_test.go
type Header struct {
	Origin  string
	Commit  string
	Version string
}

func (h Header) String() string {
	s := fmt.Sprintf("// Source: %s\n", h.Origin)
	if h.Commit != "" {
		s += fmt.Sprintf("// Commit: %s\n", h.Commit)
	}
	if h.Version != "" {
		s += fmt.Sprintf("// Problem Specifications Version: %s\n", h.Version)
	}
	return s
}

func init() {
	if _, path, _, ok := runtime.Caller(0); ok {
		dirMetadata = filepath.Join(path, "..", "..", "..", "problem-specifications")
	}
}

// outputSource puts the src text into given fileName
// and outputs a log message with given [status].
func outputSource(status string, fileName string, src []byte) error {
	ioerr := ioutil.WriteFile(fileName, src, 0666)
	if ioerr != nil {
		log.Printf("[FAILED] %s\n", ioerr)
		return ioerr
	}
	log.Printf("[%s] output: %s\n", status, fileName)
	return nil
}

// Gen generates the exercise cases_test.go file from the relevant canonical-data.json
func Gen(exercise string, j interface{}, t *template.Template) error {
	if dirMetadata == "" {
		return errors.New("unable to determine current path")
	}
	// Determine the exercise directory.
	// Use runtime.Caller to determine path location of the generator main package.
	// Call frames: 0 is this frame, Gen().
	//              1 should be a func in <exercise-dir>/.meta/gen.go (generator main package)
	//                which is calling Gen().
	//                                  |
	//                                  V
	if _, path, _, ok := runtime.Caller(1); ok {
		// Construct a path 2 directories higher than the path for .meta/gen.go
		// and it should be exercise directory.
		dirExercise = filepath.Join(path, "..", "..")
	}
	if dirExercise == "" {
		dirExercise = "."
	}
	jFile := filepath.Join("exercises", exercise, "canonical-data.json")
	// try to find and read the local json source file
	log.Printf("[LOCAL] fetching %s test data\n", exercise)
	jPath, jOrigin, jCommit := getLocal(jFile)
	jFilePath := filepath.Join(jPath, jFile)
	if jPath != "" {
		log.Printf("[LOCAL] source: %s\n", jFilePath)
	}
	jSrc, err := ioutil.ReadFile(jFilePath)
	if err != nil {
		// fetch json data remotely if there's no local file
		log.Println("[LOCAL] No test data found")
		log.Printf("[REMOTE] fetching %s test data\n", exercise)
		jSrc, jOrigin, jCommit, err = getRemote(exercise)
		if err != nil {
			return err
		}
	}

	// unmarshal the json source to a Go structure
	if err = json.Unmarshal(jSrc, j); err != nil {
		// This error message is usually enough if the problem is a wrong
		// data structure defined here. Sadly it doesn't locate the error well
		// in the case of invalid JSON.  Use a real validator tool if you can't
		// spot the problem right away.
		return fmt.Errorf(`unexpected data structure: %v`, err)
	}

	// These fields are guaranteed to be in every problem
	var commonMetadata struct {
		Version string
	}
	if err := json.Unmarshal(jSrc, &commonMetadata); err != nil {
		return fmt.Errorf(`didn't contain version: %v`, err)
	}

	if err := classifyByProperty(j); err != nil {
		return fmt.Errorf("couldn't auto-classify based on property: %v", err)
	}

	// package up a little meta data
	d := struct {
		Header
		J interface{}
	}{Header{
		Origin:  jOrigin,
		Commit:  jCommit,
		Version: commonMetadata.Version,
	}, j}

	casesFileName := filepath.Join(dirExercise, "cases_test.go")

	// render the Go test cases
	var b bytes.Buffer
	if err := t.Execute(&b, &d); err != nil {
		log.Print("[ERROR] template.Execute failed. The template has a semantic error.")
		return err
	}
	// clean it up
	srcBuf := b.Bytes()
	src, err := format.Source(srcBuf)
	if err != nil {
		log.Print("[ERROR] format.Source failed. The generated source has a syntax error.")
		b.Reset()
		_, _ = b.Write(srcBuf)
		_, _ = b.Write([]byte(
			"// !NOTE: Error during source formatting: Line:Column " + fmt.Sprint(err) + "\n"))
		src = b.Bytes()
		// Save the raw unformatted, error-containing source for purposes of debugging the generator.
		_ = outputSource("ERROR", casesFileName, src)
		return err
	}
	// write output file for the Go test cases.
	return outputSource("SUCCESS", casesFileName, src)
}

func getLocal(jFile string) (jPath, jOrigin, jCommit string) {
	// Ideally draw from a .json which is pulled from the official problem-specifications
	// repository.  For development however, accept a file in current directory
	// if there is no .json in source control.  Also allow an override in any
	// case by environment variable.
	const local_file = "local file"
	if jPath := os.Getenv("EXTEST"); jPath > "" {
		return jPath, local_file, "" // override
	}
	c := exec.Command("git", "log", "-1", "--oneline", jFile)
	c.Dir = dirMetadata
	origin, err := c.Output()
	if err != nil {
		return "", local_file, "" // no source control
	}
	if _, err = os.Stat(filepath.Join(c.Dir, jFile)); err != nil {
		return "", local_file, "" // not in source control
	}
	// good.  return source control dir and commit.
	return c.Dir, "exercism/problem-specifications", string(bytes.TrimSpace(origin))
}

func getRemote(exercise string) (body []byte, jOrigin string, jCommit string, err error) {
	url := fmt.Sprintf(canonicalDataURL, exercise)
	resp, err := genClient.Get(url)
	if err != nil {
		return []byte{}, "", "", err
	}
	if resp.StatusCode != http.StatusOK {
		return []byte{}, "", "", fmt.Errorf("error fetching remote data: %s", resp.Status)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, "", "", err
	}
	c, err := getRemoteCommit(exercise)
	if err != nil {
		// we always expect to have the commit in the cases_test.go
		// file, so return the error if we can't fetch it
		return []byte{}, "", "", err
	}
	log.Printf("[REMOTE] source: %s\n", url)
	return body, "exercism/problem-specifications", c, nil
}

func getRemoteCommit(exercise string) (string, error) {
	type Commits struct {
		Sha    string
		Commit struct {
			Message string
		}
	}
	resp, err := genClient.Get(fmt.Sprintf(commitsURL, exercise))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var c []Commits
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		return "", err
	}
	// Use only 1st line of the commit message
	lines := strings.SplitN(c[0].Commit.Message, "\n", 2)
	return fmt.Sprintf("%s %s", c[0].Sha[0:7], lines[0]), nil
}
