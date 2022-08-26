package ledger

import (
	"reflect"
	"testing"
)

var successTestCases = []struct {
	name     string
	currency string
	locale   string
	entries  []Entry
	expected string
}{
	{
		name:     "empty ledger",
		currency: "USD",
		locale:   "en-US",
		entries:  nil,
		expected: `
Date       | Description               | Change
`,
	},
	{
		name:     "one entry",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      ($10.00)
`,
	},
	{
		name:     "credit and debit",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-02",
				Description: "Get present",
				Change:      1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      ($10.00)
01/02/2015 | Get present               |       $10.00 
`,
	},
	{
		name:     "multiple entries on same date ordered by description",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Get present",
				Change:      1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      ($10.00)
01/01/2015 | Get present               |       $10.00 
`,
	},
	{
		name:     "final order tie breaker is change",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      0,
			},
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      -1,
			},
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      1,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Something                 |       ($0.01)
01/01/2015 | Something                 |        $0.00 
01/01/2015 | Something                 |        $0.01 
`,
	},
	{
		name:     "overlong descriptions",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Freude schoner Gotterfunken",
				Change:      -123456,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Freude schoner Gotterf... |   ($1,234.56)
`,
	},
	{
		name:     "euros",
		currency: "EUR",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: `
Date       | Description               | Change
01/01/2015 | Buy present               |      (€10.00)
`,
	},
	{
		name:     "Dutch locale",
		currency: "USD",
		locale:   "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      123456,
			},
		},
		expected: `
Datum      | Omschrijving              | Verandering
12-03-2015 | Buy present               |   $ 1.234,56 
`,
	},
	{
		name:     "Dutch negative number with 3 digits before decimal point",
		currency: "USD",
		locale:   "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: `
Datum      | Omschrijving              | Verandering
12-03-2015 | Buy present               |     $ 123,45-
`,
	},
	{
		name:     "American negative number with 3 digits before decimal point",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: `
Date       | Description               | Change
03/12/2015 | Buy present               |     ($123.45)
`,
	},
}

var failureTestCases = []struct {
	name     string
	currency string
	locale   string
	entries  []Entry
}{
	{
		name:     "empty currency",
		currency: "",
		locale:   "en-US",
		entries:  nil,
	},
	{
		name:     "invalid currency",
		currency: "ABC",
		locale:   "en-US",
		entries:  nil,
	},
	{
		name:     "empty locale",
		currency: "USD",
		locale:   "",
		entries:  nil,
	},
	{
		name:     "invalid locale",
		currency: "USD",
		locale:   "nl-US",
		entries:  nil,
	},
	{
		name:     "invalid date (way too high month)",
		currency: "USD",
		locale:   "en-US",
		entries: []Entry{
			{
				Date:        "2015-131-11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
	{
		name:     "invalid date (wrong separator)",
		currency: "USD",
		locale:   "en-US",
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
	for _, tc := range successTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := FormatLedger(tc.currency, tc.locale, tc.entries)
			if err != nil {
				t.Fatalf("FormatLedger for input named %q returned unexpected error %v", tc.name, err)
			}
			expected := tc.expected[1:] // Strip initial newline
			if actual != expected {
				t.Fatalf("FormatLedger for input named %q failed\ngot:\n%s\nwant:\n%s", tc.name, actual, expected)
			}
		})
	}
}

func TestFormatLedgerFailure(t *testing.T) {
	for _, tt := range failureTestCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FormatLedger(tt.currency, tt.locale, tt.entries)
			if err == nil {
				t.Fatalf("FormatLedger for input %q expected error, got nil", tt.name)
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
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range successTestCases {
			FormatLedger(tt.currency, tt.locale, tt.entries)
		}
	}
}
