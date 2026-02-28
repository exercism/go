package transmission

import (
	"slices"
	"testing"
)

func TestTransmit(t *testing.T) {
	for _, tc := range transmitCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Transmit(tc.input)
			if !slices.Equal(tc.expected, actual) {
				t.Errorf("Transmit([% x])\nexpected: [% x]\ngot:      [% x]", tc.input, tc.expected, actual)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range decodeCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Decode(tc.input)
			if tc.expectedError != "" {
				if err == nil {
					t.Errorf("Decode([% x])\nexpected error: %q\ngot value: [% x]", tc.input, tc.expectedError, actual)
				} else if tc.expectedError != err.Error() {
					t.Errorf("Decode([% x])\nexpected error: %q\ngot:            %q", tc.input, tc.expectedError, err.Error())
				}
			} else if !slices.Equal(tc.expectedValue, actual) {
				t.Errorf("Decode([% x])\nexpected: [% x]\ngot:      [% x]", tc.input, tc.expectedValue, actual)
			}
		})
	}
}
