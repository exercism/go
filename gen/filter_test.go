package gen

import (
	"bytes"
	"strings"
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
		"input": {
			"some": "inp"
		},
		"uuid": "alskjdb-f781-4c52-b73b-d4b867f41540"
		},
		{
		"description": "test case 2",
		"expected": "rvedv",
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
		"8snv0f-f781-4c52-b73b-8e76427defd0":  struct{}{},
		"klnhng-f781-4c52-b73b-8e76427defd0":  struct{}{},
		"98axn89-29f9-46f2-8c95-6c5b7a595aee": struct{}{},
	}

	expectedJson = `
{
	"cases": [
		{
			"description": "test case 1",
			"expected": "abcde",
			"input": {
				"some": "inp"
			},
			"uuid": "alskjdb-f781-4c52-b73b-d4b867f41540"
		},
		{
			"cases": [
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
			"cases": [
				{
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
						}
					],
					"description": "nested cases"
				}
			],
			"description": "some nested cases"
		}
	],
	"comments": [
		"comment 123",
		"comment 456"
	],
	"exercise": "some-exercise-name"
}
`
)

func TestFilterTestsJson(t *testing.T) {
	tests := []struct {
		description    string
		inputJson      []byte
		excludeList    map[string]struct{}
		expectedOutput []byte
		wantErr        bool
	}{
		{
			description:    "Filter valid json successfully",
			inputJson:      []byte(validInputJson),
			excludeList:    excludeList,
			expectedOutput: []byte(strings.TrimSpace(expectedJson)),
		},
		{
			description: "Filtering invalid json should fail",
			inputJson:   []byte("{\"asd"),
			excludeList: excludeList,
			wantErr:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output, err := filterTestsJson(test.inputJson, test.excludeList)
			if test.wantErr && err == nil {
				t.Errorf("expected error but got none")
			}
			if !test.wantErr && err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if !bytes.Equal(test.expectedOutput, output) {
				t.Fatalf("wrong output. expected: %s, got %s", test.expectedOutput, output)
			}
		})
	}
}
