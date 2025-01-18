package expenses

import "testing"

var testExpensesRecords = []Record{
	{
		Day:      1,
		Amount:   5.15,
		Category: "groceries",
	},
	{
		Day:      1,
		Amount:   3.45,
		Category: "groceries",
	},
	{
		Day:      13,
		Amount:   55.67,
		Category: "utility-bills",
	},
	{
		Day:      15,
		Amount:   11,
		Category: "groceries",
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
		Category: "groceries",
	},
	{
		Day:      25,
		Amount:   24.65,
		Category: "groceries",
	},
	{
		Day:      30,
		Amount:   1300,
		Category: "rent",
	},
}

// testRunnerTaskID=2
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
					Category: "groceries",
				},
				{
					Day:      1,
					Amount:   3.45,
					Category: "groceries",
				},
				{
					Day:      13,
					Amount:   55.67,
					Category: "utility-bills",
				},
				{
					Day:      15,
					Amount:   11,
					Category: "groceries",
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

// testRunnerTaskID=3
func TestFilterByCategory(t *testing.T) {
	testCases := []struct {
		name     string
		category string
		expected []Record
	}{
		{
			name:     "returns expenses in groceries category",
			category: "groceries",
			expected: []Record{
				{
					Day:      1,
					Amount:   5.15,
					Category: "groceries",
				},
				{
					Day:      1,
					Amount:   3.45,
					Category: "groceries",
				},
				{
					Day:      15,
					Amount:   11,
					Category: "groceries",
				},
				{
					Day:      23,
					Amount:   20.0,
					Category: "groceries",
				},
				{
					Day:      25,
					Amount:   24.65,
					Category: "groceries",
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

// testRunnerTaskID=4
func TestTotalByPeriod(t *testing.T) {
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
			got := TotalByPeriod(testExpensesRecords, tC.p)
			if got != tC.total {
				t.Errorf("TotalByPeriod(%v, %v) = %.2f, want %.2f", testExpensesRecords, tC.p, got, tC.total)
			}
		})
	}
}

// testRunnerTaskID=5
func TestCategoryExpenses(t *testing.T) {
	testCases := []struct {
		name     string
		category string
		p        DaysPeriod
		total    float64
		wantErr  bool
	}{
		{
			name:     "returns error when no records with category found in any days period",
			category: "food",
			p: DaysPeriod{
				From: 1,
				To:   30,
			},
			total:   0,
			wantErr: true,
		},
		{
			name:     "returns total category expenses in the provided days period",
			category: "groceries",
			p: DaysPeriod{
				From: 1,
				To:   15,
			},
			total:   19.6,
			wantErr: false,
		},
		{
			name:     "returns 0 when no category expenses found in the provided days period",
			category: "groceries",
			p: DaysPeriod{
				From: 40,
				To:   50,
			},
			total:   0,
			wantErr: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got, err := CategoryExpenses(testExpensesRecords, tC.p, tC.category)
			if tC.wantErr && err == nil {
				t.Fatalf("CategoryExpenses(%v, %s, %v)=%.2f,%v but want a non-nil error",
					testExpensesRecords, tC.category, tC.p, got, err)
			}

			if !tC.wantErr && err != nil {
				t.Fatalf("CategoryExpenses(%v, %s, %v)=%.2f,%v but a non-nil error was not expected",
					testExpensesRecords, tC.category, tC.p, got, err)
			}

			if got != tC.total {
				var errStr string

				if tC.wantErr {
					errStr = "unknown category"
				} else {
					errStr = "nil"
				}
				t.Fatalf("CategoryExpenses(%v, %s, %v) = %.2f,%v but want %.2f,%s",
					testExpensesRecords, tC.category, tC.p, got, err, tC.total, errStr)
			}
		})
	}
}
