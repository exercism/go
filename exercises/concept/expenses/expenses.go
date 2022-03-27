package expenses

import "time"

// Record describes a record of expense event.
type Record struct {
	Date     time.Time
	Amount   float64
	Category string
}

// DatePeriod describes date period.
type DatePeriod struct {
	From time.Time
	To   time.Time
}

// Filter creates a new records collection by applying the predicate function to
// records and keeping the records for which the function returns true.
func Filter(in []Record, f func(Record) bool) (out []Record) {
	panic("Please implement the Filter function.")
}

// ByDatePeriod returns predicate function. The predicate returns true when
// Record.Date is within the date period.
func ByDatePeriod(p DatePeriod) func(Record) bool {
	panic("Please implement the ByDatePeriod function.")
}

// ByCategory returns predicate function. The predicate returns true when
// Record.Category is equal to the provided category.
func ByCategory(c string) func(Record) bool {
	panic("Please implement the ByCategory function.")
}

// Total returns total amount of expenses in collection of records rr
// within time period p.
func Total(in []Record, p DatePeriod) float64 {
	panic("Please implement the Total function.")
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses records.
func CategoryExpenses(in []Record, p DatePeriod, c string) (float64, error) {
	panic("Please implement the CategoryExpenses function.")
}
