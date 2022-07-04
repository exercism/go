package gen

import (
	"encoding/json"
)

func filterTestsJson(jsrc []byte, excludeList map[string]struct{}) ([]byte, error) {
	// put the json object in an array to match the recursive structure
	jsrcArr := make([]byte, 1, len(jsrc)+2)
	jsrcArr[0] = '['
	jsrcArr = append(jsrcArr, jsrc...)
	jsrcArr = append(jsrcArr, ']')

	// recursively remove excluded cases from json source, starting from top level
	filteredCases, err := recursiveFilterCases(json.RawMessage(jsrcArr), excludeList)
	if err != nil {
		return nil, err
	}

	// remove the json object back out of the array
	filteredData, err := json.MarshalIndent(filteredCases[0], "", "\t")
	if err != nil {
		return nil, err
	}
	return filteredData, nil
}

func recursiveFilterCases(cases json.RawMessage, excludeList map[string]struct{}) ([]map[string](json.RawMessage), error) {
	var js []map[string](json.RawMessage)
	var validCases []map[string](json.RawMessage)

	// 'cases' is always an array, where every item is either a single test or an object containing nested cases
	err := json.Unmarshal(cases, &js)
	if err != nil {
		return validCases, err
	}

	// iterating over every item in 'cases'
	for _, j := range js {
		var uuidStr string
		uuid, ok := j["uuid"]

		// If uuid key is not present, this item is an object containing nested cases.
		// So we recursively filter this nested array and set the filtered nested json.
		if !ok {
			filteredCases, err := recursiveFilterCases(j["cases"], excludeList)
			if err != nil {
				return validCases, err
			}

			// convert filteredCases map to json string
			filteredCasesJson, err := json.Marshal(filteredCases)
			if err != nil {
				return validCases, err
			}

			// typecast it as json.RawMessage to match types
			j["cases"] = json.RawMessage(filteredCasesJson)
			validCases = append(validCases, j)
			continue
		}

		// If uuid key is present, this item is a single test case.
		// So we check if this uuid is present in the exclude list.
		// If it is not present, we add it to the list of valid cases.
		err := json.Unmarshal(uuid, &uuidStr)
		if err != nil {
			return validCases, err
		}
		_, ok = excludeList[uuidStr]
		if !ok {
			validCases = append(validCases, j)
		}
	}

	return validCases, nil
}
