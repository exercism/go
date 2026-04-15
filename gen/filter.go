package gen

import (
	"encoding/json"
	"fmt"
	"slices"
)

// getAllTestCasesFiltered recursively collects all test cases in a flattened list except the ones excluded by excludedTests.
func getAllTestCasesFiltered(jSrc []byte, excludedTests map[string]struct{}) (*[]testCase, error) {
	result := &[]testCase{}

	// put the json object in an array to match the recursive structure
	jSrc = append([]byte{'['}, append(jSrc, ']')...)

	// recursively get all test cases except the excluded ones
	err := recursiveFilterCases(jSrc, result, nil, excludedTests)
	if err != nil {
		return nil, fmt.Errorf("failed to get/filter all test cases: %w", err)
	}
	return result, nil
}

func recursiveFilterCases(cases json.RawMessage, result *[]testCase, parentDescriptions []string, excludedTests map[string]struct{}) error {
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
				return fmt.Errorf("failed to unmarshal uuid %v: %w", uuid, err)
			}

			// ignore test-cases with include=false in tests.toml
			if _, isExcluded := excludedTests[uuidStr]; isExcluded {
				continue
			}

			jTestCase, err := json.Marshal(j)
			if err != nil {
				return err
			}

			tc := testCase{Parents: parentDescriptions}
			err = json.Unmarshal(jTestCase, &tc)
			if err != nil {
				return err
			}
			*result = append(*result, tc)

			continue
		}

		// If uuid key is not present, this item is an object containing nested cases.
		// Update the parent description slice.
		updatedParentDescriptions := slices.Clone(parentDescriptions)
		if _, ok := j["description"]; ok {
			var description string
			err = json.Unmarshal(j["description"], &description)
			if err != nil {
				return err
			}
			updatedParentDescriptions = append(updatedParentDescriptions, description)
		}
		err := recursiveFilterCases(j["cases"], result, updatedParentDescriptions, excludedTests)
		if err != nil {
			return err
		}
	}

	return nil
}
