package gen

import (
	"encoding/json"
	"fmt"
)

func getAllTestCasesFiltered(jsrc []byte, excludedTests map[string]struct{}) (*[]TestCase, error) {
	var result = &[]TestCase{}

	var s map[string]json.RawMessage

	err := json.Unmarshal(jsrc, &s)
	if err != nil {
		return nil, err
	}

	// recursively get all test cases except the excluded ones
	err = recursiveFilterCases(s["cases"], result, excludedTests)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func recursiveFilterCases(cases json.RawMessage, result *[]TestCase, excludedTests map[string]struct{}) error {
	var js []map[string]json.RawMessage

	// 'cases' is always an array, where every item is either a single test or an object containing nested cases
	err := json.Unmarshal(cases, &js)
	if err != nil {
		return err
	}

	// iterating over every item in 'cases'
	for _, j := range js {

		// If uuid key is present, this item is a single test-case
		if uuid, ok := j["uuid"]; ok {
			var uuidStr string

			err := json.Unmarshal(uuid, &uuidStr)
			if err != nil {
				return fmt.Errorf("failed to unmarshal uuid (%v) (%v)", uuid, err)
			}

			//ignore test-cases with include=false in tests.toml
			if _, isExcluded := excludedTests[uuidStr]; isExcluded {
				continue
			}

			testCase, err := json.Marshal(j)
			if err != nil {
				return err
			}

			var tc = TestCase{}
			err = json.Unmarshal(testCase, &tc)
			if err != nil {
				return err
			}
			*result = append(*result, tc)

			continue
		}

		// If uuid key is not present, this item is an object containing nested cases.
		err := recursiveFilterCases(j["cases"], result, excludedTests)
		if err != nil {
			return err
		}
	}

	return nil
}
