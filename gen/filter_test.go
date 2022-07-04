package gen

import (
	"reflect"
	"testing"
)

var (
	validInputJson = `
{
	"exercise": "some-exercise-name",
	"comments": [
		"comment 123",
		"comment 456"
	],
	"cases": [
		{
			"description": "test case 1",
			"expected": "abcde",
			"property" : "test",
			"input": {
				"some": "inp"
			},
			"uuid": "alskjdb-f781-4c52-b73b-d4b867f41540"
		},
		{
		"description": "test case 2",
		"expected": "rvedv",
		"property" : "test2",
		"input": {
			"some": "ukbt"
		},
		"uuid": "8snv0f-f781-4c52-b73b-8e76427defd0"
		},
		{
		"cases": [
			{
			"description": "test case 3",
			"expected": [
				0
			],
			"input": {
				"integers": [
				0
				]
			},
			"property": "some property",
			"uuid": "klnhng-f781-4c52-b73b-8e76427defd0"
			},
			{
			"description": "test case 4",
			"expected": [
				64
			],
			"input": {
				"integers": [
				64
				]
			},
			"property": "some property",
			"uuid": "dsvhsd-a151-4604-a10e-d4b867f41540"
			}
		]
		},
		{
		"description": "some nested cases",
		"cases": [
			{
			"description": "nested cases",
			"cases": [
				{
				"description": "test case 5",
				"expected": false,
				"input": {
					"bools": [
					true,
					false
					]
				},
				"property": "some other property",
				"uuid": "dvthrd4-4514-4915-bac0-f7f585e0e59a"
				},
				{
				"description": "test case 6",
				"expected": true,
				"input": {
					"bools": [
					false,
					false
					]
				},
				"property": "some other property",
				"uuid": "98axn89-29f9-46f2-8c95-6c5b7a595aee"
				}
			]
			}
		]
		}
	]
}
`

	excludeList = map[string]struct{}{
		"8snv0f-f781-4c52-b73b-8e76427defd0":  {}, //test case 2
		"klnhng-f781-4c52-b73b-8e76427defd0":  {}, //test case 3
		"98axn89-29f9-46f2-8c95-6c5b7a595aee": {}, //Test case 6
	}
)

func TestGetAllTestCasesFiltered(t *testing.T) {
	testCases := []struct {
		description    string
		inputJson      []byte
		excludeList    map[string]struct{}
		expectedOutput []TestCase
	}{
		{
			description: "Filter valid json successfully",
			inputJson:   []byte(validInputJson),
			excludeList: excludeList,
			expectedOutput: []TestCase{
				{
					UUID:        "alskjdb-f781-4c52-b73b-d4b867f41540",
					Description: "test case 1",
					Property:    "test",
					Input: map[string]interface{}{
						"some": "inp",
					},
					Expected: "abcde",
				},
				{
					UUID:        "dsvhsd-a151-4604-a10e-d4b867f41540",
					Description: "test case 4",
					Property:    "some property",
					Input: map[string]interface{}{
						"integers": []interface{}{64.0},
					},
					Expected: []interface{}{64.0},
				}, {
					UUID:        "dvthrd4-4514-4915-bac0-f7f585e0e59a",
					Description: "test case 5",
					Property:    "some other property",
					Input: map[string]interface{}{
						"bools": []interface{}{true, false},
					},
					Expected: false,
				},
			},
		},
		{
			description:    "Filtering invalid json should fail",
			inputJson:      []byte("{\"asd"),
			excludeList:    excludeList,
			expectedOutput: nil,
		},
		{
			description:    "invalid uuid (bool instead of string)",
			inputJson:      []byte("{\"cases\":[{\"uuid\":false}]}"),
			expectedOutput: nil,
		},
		{
			description:    "invalid description (number instead of string)",
			inputJson:      []byte("{\"cases\":[{\"uuid\":\"dvthrd4-4514-4915-bac0-f7f585e0e59a\", \"description\":510}]}"),
			expectedOutput: nil,
		},
		{
			description:    "invalid uuid in subcase",
			inputJson:      []byte("{\"cases\":[{\"cases\":{\"uuid\":false}}]}"),
			expectedOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			output, err := getAllTestCasesFiltered(tc.inputJson, tc.excludeList)
			if tc.expectedOutput == nil && err == nil {
				t.Errorf("expected error but got none")
			}

			if tc.expectedOutput != nil {
				if err != nil {
					t.Errorf("unexpected error %v", err)
				}
				if output == nil || !reflect.DeepEqual(tc.expectedOutput, *output) {
					t.Errorf("wrong output. expected: %v, got %v", tc.expectedOutput, output)
				}
			}
		})
	}
}
