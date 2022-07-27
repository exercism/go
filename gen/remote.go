package gen

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRemoteTestData(exercise string) ([]byte, error) {
	url := fmt.Sprintf(canonicalDataURL, exercise)
	fmt.Printf("[REMOTE] source: %s\n", url)
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch remote test data: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching remote data: (status-code: %s)", resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type Commit struct {
	Sha     string
	Details struct {
		Message string
	} `json:"commit"`
}

func (c Commit) Summary() string {
	lines := strings.SplitN(c.Details.Message, "\n", 2)
	return fmt.Sprintf("%s %s", c.Sha[0:7], lines[0])
}

func getRemoteCommit(exercise, token string) (Header, error) {
	url := fmt.Sprintf(commitsURL, exercise)
	fmt.Printf("[REMOTE] fetching latest commit (source: %s)\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Header{}, err
	}
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return Header{}, fmt.Errorf("failed to fetch latest commit: %w", err)
	}
	defer resp.Body.Close()
	var c []Commit
	err = json.NewDecoder(resp.Body).Decode(&c)
	if err != nil {
		return Header{}, err
	}

	if len(c) == 0 {
		return Header{}, errors.New("no commits found")
	}

	return Header{Commit: c[0].Summary(), Origin: defaultOrigin}, nil
}
