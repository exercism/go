package ledger

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: 2801085 Update ledger canonical-data.json (#2355)

import "strings"

type testCase = struct {
	description string
	currency    string
	locale      string
	entries     []Entry
	expected    string
}

var testCases = []testCase{
	{
		description: "empty ledger",
		currency:    "USD",
		locale:      "en-US",
		entries:     []Entry{},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
		}, "\n") + "\n",
	},
	{
		description: "one entry",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      ($10.00)",
		}, "\n") + "\n",
	},
	{
		description: "credit and debit",
		currency:    "USD",
		locale:      "en-US",
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
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      ($10.00)",
			"01/02/2015 | Get present               |       $10.00 ",
		}, "\n") + "\n",
	},
	{
		description: "final order tie breaker is change",
		currency:    "USD",
		locale:      "en-US",
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
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Something                 |       ($0.01)",
			"01/01/2015 | Something                 |        $0.00 ",
			"01/01/2015 | Something                 |        $0.01 ",
		}, "\n") + "\n",
	},
	{
		description: "overlong description is truncated",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Freude schoner Gotterfunken",
				Change:      -123456,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Freude schoner Gotterf... |   ($1,234.56)",
		}, "\n") + "\n",
	},
	{
		description: "euros",
		currency:    "EUR",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      (€10.00)",
		}, "\n") + "\n",
	},
	{
		description: "Dutch locale",
		currency:    "USD",
		locale:      "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      123456,
			},
		},
		expected: strings.Join([]string{
			"Datum      | Omschrijving              | Verandering  ",
			"12-03-2015 | Buy present               |   $ 1.234,56 ",
		}, "\n") + "\n",
	},
	{
		description: "Dutch locale and euros",
		currency:    "EUR",
		locale:      "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      123456,
			},
		},
		expected: strings.Join([]string{
			"Datum      | Omschrijving              | Verandering  ",
			"12-03-2015 | Buy present               |   € 1.234,56 ",
		}, "\n") + "\n",
	},
	{
		description: "Dutch negative number with 3 digits before decimal point",
		currency:    "USD",
		locale:      "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: strings.Join([]string{
			"Datum      | Omschrijving              | Verandering  ",
			"12-03-2015 | Buy present               |    $ -123,45 ",
		}, "\n") + "\n",
	},
	{
		description: "American negative number with 3 digits before decimal point",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"03/12/2015 | Buy present               |     ($123.45)",
		}, "\n") + "\n",
	},
	{
		description: "multiple entries on same date ordered by description",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Get present",
				Change:      1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      ($10.00)",
			"01/01/2015 | Get present               |       $10.00 ",
		}, "\n") + "\n",
	},
}
