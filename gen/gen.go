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
	"text/template"
	"time"
)

// dirMetadata is the location of the x-common repository on the filesystem.
// We're making the assumption that the x-common repository has been cloned to
// the same parent directory as the xgo repository.
// E.g.
//
//     $ tree -L 1 .
//     .
//     ├── x-common
//     └── xgo
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
	canonicalDataURL = "https://raw.githubusercontent.com/exercism/x-common/master/exercises/%s/canonical-data.json"
	// commitsURL is the GitHub api endpoint for the canonical-data.json
	// file commit history, requires exercise name.
	commitsURL = "https://api.github.com/repos/exercism/x-common/commits?path=exercises/%s/canonical-data.json"
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
		s += fmt.Sprintf("// x-common version: %s\n", h.Version)
	}
	return s
}

func init() {
	if _, path, _, ok := runtime.Caller(0); ok {
		dirMetadata = filepath.Join(path, "..", "..", "..", "x-common")
	}
	if _, path, _, ok := runtime.Caller(2); ok {
		dirExercise = filepath.Join(path, "..", "..")
	}
	if dirExercise == "" {
		dirExercise = "."
	}
}

// Gen generates the exercise cases_test.go file from the relevant canonical-data.json
func Gen(exercise string, j interface{}, t *template.Template) error {
	if dirMetadata == "" {
		return errors.New("unable to determine current path")
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

	// package up a little meta data
	d := struct {
		Header
		J interface{}
	}{Header{
		Origin:  jOrigin,
		Commit:  jCommit,
		Version: commonMetadata.Version,
	}, j}

	// render the Go test cases
	var b bytes.Buffer
	if err := t.Execute(&b, &d); err != nil {
		return err
	}
	// clean it up
	src, err := format.Source(b.Bytes())
	if err != nil {
		return err
	}
	// write output file for the Go test cases.
	return ioutil.WriteFile(filepath.Join(dirExercise, "cases_test.go"), src, 0666)
}

func getLocal(jFile string) (jPath, jOrigin, jCommit string) {
	// Ideally draw from a .json which is pulled from the official x-common
	// repository.  For development however, accept a file in current directory
	// if there is no .json in source control.  Also allow an override in any
	// case by environment variable.
	if jPath := os.Getenv("EXTEST"); jPath > "" {
		return jPath, "local file", "" // override
	}
	c := exec.Command("git", "log", "-1", "--oneline", jFile)
	c.Dir = dirMetadata
	origin, err := c.Output()
	if err != nil {
		return "", "local file", "" // no source control
	}
	if _, err = os.Stat(filepath.Join(c.Dir, jFile)); err != nil {
		return "", "local file", "" // not in source control
	}
	// good.  return source control dir and commit.
	return c.Dir, "exercism/x-common", string(bytes.TrimSpace(origin))
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
	return body, "exercism/x-common", c, nil
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
	return fmt.Sprintf("%s %s", c[0].Sha[0:7], c[0].Commit.Message), nil
}
