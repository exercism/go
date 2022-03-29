package expenses

import "testing"

var testExpensesRecords = []Record{
	{
		Day:      1,
		Amount:   5.15,
		Category: "grocieries",
	},
	{
		Day:      1,
		Amount:   3.45,
		Category: "grocieries",
	},
	{
		Day:      13,
		Amount:   55.67,
		Category: "utility-bills",
	},
	{
		Day:      15,
		Amount:   11,
		Category: "grocieries",
	},
	{
		Day:      18,
		Amount:   244.33,
		Category: "utility-bills",
	},
	{
		Day:      20,
		Amount:   300,
		Category: "university",
	},
	{
		Day:      23,
		Amount:   20.0,
		Category: "grocieries",
	},
	{
		Day:      25,
		Amount:   24.65,
		Category: "grocieries",
	},
	{
		Day:      30,
		Amount:   1300,
		Category: "rent",
	},
}

func TestFilterByDaysPeriod(t *testing.T) {
	testCases := []struct {
		name     string
		p        DaysPeriod
		expected []Record
	}{
		{
			name: "returns expenses records from 1st to 15th day",
			p: DaysPeriod{
				From: 1,
				To:   15,
			},
			expected: []Record{
				{
					Day:      1,
					Amount:   5.15,
					Category: "grocieries",
				},
				{
					Day:      1,
					Amount:   3.45,
					Category: "grocieries",
				},
				{
					Day:      13,
					Amount:   55.67,
					Category: "utility-bills",
				},
				{
					Day:      15,
					Amount:   11,
					Category: "grocieries",
				},
			},
		},
		{
			name: "returns empty list when no expenses found in the days period",
			p: DaysPeriod{
				From: 40,
				To:   50,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := Filter(testExpensesRecords, ByDaysPeriod(tC.p))
			if len(got) != len(tC.expected) {
				t.Fatalf("Filter by period got %d records, want %d", len(got), len(tC.expected))
			}

			for i, expected := range tC.expected {
				if got[i] != expected {
					t.Fatalf("Filter by period got %v, want %v", got, tC.expected)
				}
			}
		})
	}
}

func TestFilterByCategory(t *testing.T) {
	testCases := []struct {
		name     string
		category string
		expected []Record
	}{
		{
			name:     "returns expenses in grocieries category",
			category: "grocieries",
			expected: []Record{
				{
					Day:      1,
					Amount:   5.15,
					Category: "grocieries",
				},
				{
					Day:      1,
					Amount:   3.45,
					Category: "grocieries",
				},
				{
					Day:      15,
					Amount:   11,
					Category: "grocieries",
				},
				{
					Day:      23,
					Amount:   20.0,
					Category: "grocieries",
				},
				{
					Day:      25,
					Amount:   24.65,
					Category: "grocieries",
				},
			},
		},
		{
			name:     "returns empty list for unknown category",
			category: "ABC",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := Filter(testExpensesRecords, ByCategory(tC.category))
			if len(got) != len(tC.expected) {
				t.Fatalf("Filter by category got %d records, want %d", len(got), len(tC.expected))
			}

			for i, expected := range tC.expected {
				if got[i] != expected {
					t.Fatalf("Filter by category got %v, want %v", got, tC.expected)
				}
			}
		})
	}
}

func TestTotal(t *testing.T) {
	testCases := []struct {
		name  string
		p     DaysPeriod
		total float64
	}{
		{
			name: "total expenses is 0 when no records found in the provided days period",
			p: DaysPeriod{
				From: 40,
				To:   50,
			},
			total: 0,
		},
		{
			name: "total expenses for days period from 25th to 26th day",
			p: DaysPeriod{
				From: 25,
				To:   26,
			},
			total: 24.65,
		},
		{
			name: "total expenses for the full days period",
			p: DaysPeriod{
				From: 1,
				To:   100,
			},
			total: 1964.25,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := Total(testExpensesRecords, tC.p)
			if got != tC.total {
				t.Errorf("Total(%v, %v) = %.2f, want %.2f", testExpensesRecords, tC.p, got, tC.total)
			}
		})
	}
}

func TestCategoryExpenses(t *testing.T) {
	testCases := []struct {
		name     string
		category string
		p        DaysPeriod
		total    float64
		err      string
	}{
		{
			name:     "returns error when no records with category found in any days period",
			category: "food",
			p: DaysPeriod{
				From: 1,
				To:   30,
			},
			total: 0,
			err:   "unknown category food",
		},
		{
			name:     "returns total category expenses in the provided days period",
			category: "grocieries",
			p: DaysPeriod{
				From: 1,
				To:   15,
			},
			total: 19.6,
			err:   "",
		},
		{
			name:     "returns 0 when no category expenses found in the provided days period",
			category: "grocieries",
			p: DaysPeriod{
				From: 40,
				To:   50,
			},
			total: 0,
			err:   "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := CategoryExpenses(testExpensesRecords, tC.p, tC.category)
			if tC.err != "" && (err == nil || tC.err != err.Error()) {
				t.Errorf("CategoryExpenses(%v, %s, %v) failed: %v, want: %s",
					testExpensesRecords, tC.category, tC.p, err, tC.err)
			}

			if tC.err == "" && err != nil {
				t.Errorf("CategoryExpenses(%v, %s, %v) failed: %v, want %.2f, %s",
					testExpensesRecords, tC.category, tC.p, err, tC.total, tC.err)
			}

			if tC.err == "" && err == nil && got != tC.total {
				t.Errorf("CategoryExpenses(%v, %s, %v) = %.2f, want %.2f",
					testExpensesRecords, tC.category, tC.p, got, tC.total)
			}
		})
	}
}
