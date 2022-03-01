package expenses

import (
	"testing"
	"time"
)

var testExpensesRecords = []Record{
	{
		Date:     time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
		Amount:   5.15,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
		Amount:   3.45,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 13, 0, 0, 0, 0, time.UTC),
		Amount:   55.67,
		Category: "utility-bills",
	},
	{
		Date:     time.Date(2021, time.December, 15, 0, 0, 0, 0, time.UTC),
		Amount:   11,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 18, 0, 0, 0, 0, time.UTC),
		Amount:   244.33,
		Category: "utility-bills",
	},
	{
		Date:     time.Date(2021, time.December, 20, 0, 0, 0, 0, time.UTC),
		Amount:   300,
		Category: "university",
	},
	{
		Date:     time.Date(2021, time.December, 23, 0, 0, 0, 0, time.UTC),
		Amount:   20.0,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC),
		Amount:   24.65,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 30, 0, 0, 0, 0, time.UTC),
		Amount:   1300,
		Category: "rent",
	},
}

func stringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestTotal(t *testing.T) {
	testCases := []struct {
		name  string
		p     DatePeriod
		total float64
	}{
		{
			name: "total expenses is 0 when no records found in the provided date period",
			p: DatePeriod{
				From: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC),
			},
			total: 0,
		},
		{
			name: "total expenses for partial date period",
			p: DatePeriod{
				From: time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2021, time.December, 26, 0, 0, 0, 0, time.UTC),
			},
			total: 24.65,
		},
		{
			name: "total expenses for the full date period",
			p: DatePeriod{
				From: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC),
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
		p        DatePeriod
		total    float64
		err      string
	}{
		{
			name:     "returns error when no records with category found in any date period",
			category: "food",
			p: DatePeriod{
				From: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC),
			},
			total: 0,
			err:   "unknown category food",
		},
		{
			name:     "returns total category expenses in the provided date period",
			category: "grocieries",
			p: DatePeriod{
				From: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2021, time.December, 15, 0, 0, 0, 0, time.UTC),
			},
			total: 19.6,
			err:   "",
		},
		{
			name:     "returns 0 when no category expenses found in the provided date period",
			category: "grocieries",
			p: DatePeriod{
				From: time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC),
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
