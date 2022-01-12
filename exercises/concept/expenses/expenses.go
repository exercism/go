package expenses

import "time"

// Record describes a record of expense event.
type Record struct {
	Date        time.Time
	Amount      float64
	Description string
	Category    string
}

// Period describes time period.
type Period struct {
	DateFrom time.Time
	DateTo   time.Time
}

// Total returns total amount of expenses in collection a, within time period p.
func Total(a []Record, p Period) float64 {
	panic("Total not implemented")
}

// TopCategoriesN returns top n categories of expenses in collection a, within
// time period p.
func TopCategoriesN(a []Record, n int, p Period) []string {
	panic("TopCategoriesN not implemented")
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses collection a.
func CategoryExpenses(a []Record, c string, p Period) (float64, error) {
	panic("CategoryExpenses not implemented")
}
