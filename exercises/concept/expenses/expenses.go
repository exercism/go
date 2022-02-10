package expenses

import "time"

// Record describes a record of expense event.
type Record struct {
	Date        time.Time
	Amount      float64
	Description string
	Category    string
}

// Records describes a collection of expenses events.
type Records []Record

// Filter creates a new records collection by applying the predicate function to
// records and keeping the records for which the function returns true.
func (rr Records) Filter(f func(Record) bool) Records {
	var out Records
	for _, r := range rr {
		if f(r) {
			out = append(out, r)
		}
	}
	return out
}

// Period describes time period.
type Period struct {
	DateFrom time.Time
	DateTo   time.Time
}

// Total returns total amount of expenses in collection of records rr
// within time period p.
func Total(rr Records, p Period) float64 {
	panic("Total not implemented")
}

// TopCategoriesN returns top n categories of expenses in collection of records rr
// within time period p.
func TopCategoriesN(rr Records, n int, p Period) []string {
	panic("TopCategoriesN not implemented")
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses records.
func CategoryExpenses(rr Records, c string, p Period) (float64, error) {
	panic("CategoryExpenses not implemented")
}
