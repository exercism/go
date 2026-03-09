package ledger

import (
	"reflect"
	"testing"
)

var failureTestCases = []testCase{
	{
		description: "empty currency",
		currency:    "",
		locale:      "en-US",
		entries:     nil,
	},
	{
		description: "invalid currency",
		currency:    "ABC",
		locale:      "en-US",
		entries:     nil,
	},
	{
		description: "empty locale",
		currency:    "USD",
		locale:      "",
		entries:     nil,
	},
	{
		description: "invalid locale",
		currency:    "USD",
		locale:      "nl-US",
		entries:     nil,
	},
	{
		description: "invalid date (way too high month)",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-131-11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
	{
		description: "invalid date (wrong separator)",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-12/11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
}

func TestFormatLedgerSuccess(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := FormatLedger(tc.currency, tc.locale, tc.entries)
			if err != nil {
				t.Fatalf("FormatLedger for test %q returned unexpected error %q", tc.description, err)
			}
			if actual != tc.expected {
				t.Fatalf("FormatLedger for test %q failed\ngot:\n%#v\nwant:\n%#v", tc.description, actual, tc.expected)
			}
		})
	}
}

func TestFormatLedgerFailure(t *testing.T) {
	for _, tt := range failureTestCases {
		t.Run(tt.description, func(t *testing.T) {
			_, err := FormatLedger(tt.currency, tt.locale, tt.entries)
			if err == nil {
				t.Fatalf("FormatLedger for test %q expected error, got nil", tt.description)
			}
		})
	}
}

func TestFormatLedgerNotChangeInput(t *testing.T) {
	entries := []Entry{
		{
			Date:        "2015-01-02",
			Description: "Freude schöner Götterfunken",
			Change:      1000,
		},
		{
			Date:        "2015-01-01",
			Description: "Buy present",
			Change:      -1000,
		},
	}
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	FormatLedger("USD", "en-US", entries)
	if !reflect.DeepEqual(entries, entriesCopy) {
		t.Fatalf("FormatLedger modifies the input entries array")
	}
}

func BenchmarkFormatLedger(b *testing.B) {
	for range b.N {
		for _, tt := range testCases {
			FormatLedger(tt.currency, tt.locale, tt.entries)
		}
	}
}
