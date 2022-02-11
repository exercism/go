package expenses

import (
	"testing"
	"time"
)

var testExpensesRecords = []Record{
	{
		Date:     time.Date(2021, time.December, 1, 10, 11, 12, 13, time.UTC),
		Amount:   5.15,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 1, 10, 11, 12, 13, time.UTC),
		Amount:   3.45,
		Category: "grocieries",
	},
	{
		Date:     time.Date(2021, time.December, 13, 11, 12, 13, 14, time.UTC),
		Amount:   55.67,
		Category: "utility-bills",
	},
	{
		Date:     time.Date(2021, time.December, 15, 8, 0, 0, 0, time.UTC),
		Amount:   11,
		Category: "entertainment",
	},
	{
		Date:     time.Date(2021, time.December, 23, 9, 0, 0, 0, time.UTC),
		Amount:   20.0,
		Category: "fitness",
	},
	{
		Date:     time.Date(2021, time.December, 25, 20, 0, 0, 0, time.UTC),
		Amount:   24.65,
		Category: "entertainment",
	},
	{
		Date:     time.Date(2021, time.December, 30, 10, 0, 0, 0, time.UTC),
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
		p     Period
		total float64
	}{
		{
			p: Period{
				DateFrom: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateTo:   time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC),
			},
			total: 0,
		},
		{
			p: Period{
				DateFrom: time.Date(2021, time.December, 25, 0, 0, 0, 0, time.UTC),
				DateTo:   time.Date(2021, time.December, 30, 0, 0, 0, 0, time.UTC),
			},
			total: 24.65,
		},
		{
			p: Period{
				DateFrom: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
				DateTo:   time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			total: 1419.92,
		},
	}
	for _, tC := range testCases {
		got := Total(testExpensesRecords, tC.p)
		if got != tC.total {
			t.Errorf("Total(%v, %v) = %.2f, want %.2f", testExpensesRecords, tC.p, got, tC.total)
		}
	}
}

func TestTopCategoriesN(t *testing.T) {
	testCases := []struct {
		p        Period
		n        int
		expected []string
	}{
		{
			p: Period{
				DateFrom: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
				DateTo:   time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			n:        2,
			expected: []string{"rent", "utility-bills"},
		},
	}
	for _, tC := range testCases {
		got := TopCategoriesN(testExpensesRecords, tC.p, tC.n)
		if !stringsEqual(got, tC.expected) {
			t.Errorf("TopCategoriesN(%v, %d, %v) = %v, want: %v",
				testExpensesRecords, tC.n, tC.p, got, tC.expected)
		}
	}
}

func TestCategoryExpenses(t *testing.T) {
	testCases := []struct {
		category string
		p        Period
		total    float64
		err      string
	}{
		{
			category: "food",
			p: Period{
				DateFrom: time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				DateTo:   time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC),
			},
			total: 0,
			err:   "unknown category: food",
		},
		{
			category: "entertainment",
			p: Period{
				DateFrom: time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
				DateTo:   time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC),
			},
			total: 35.65,
			err:   "",
		},
	}
	for _, tC := range testCases {
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
	}
}
