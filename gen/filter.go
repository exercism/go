package gen

import "encoding/json"

func filterTestsJson(jsrc []byte, excludeList map[string]struct{}) ([]byte, error) {
	var s map[string](json.RawMessage)
	var filteredData []byte
	err := json.Unmarshal(jsrc, &s)
	if err != nil {
		return filteredData, err
	}

	// recursively remove excluded cases from json source, starting from top level
	filteredCases, err := recursiveFilterCases(s["cases"], excludeList)
	if err != nil {
		return filteredData, err
	}

	// convert filteredCases map to json string
	filteredCasesJson, err := json.Marshal(filteredCases)
	if err != nil {
		return filteredData, err
	}

	// typecast it as json.RawMessage to match types
	s["cases"] = json.RawMessage(filteredCasesJson)

	resp, err := json.Marshal(s)
	if err != nil {
		return filteredData, err
	}
	return resp, nil
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
