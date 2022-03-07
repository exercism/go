package expenses

import "time"

// Record describes a record of expense event.
type Record struct {
	Date     time.Time
	Amount   float64
	Category string
}

// Filter creates a new records collection by applying the predicate function to
// records and keeping the records for which the function returns true.
func Filter(in []Record, f func(Record) bool) (out []Record) {
	panic("Total not implemented")
}

// DatePeriod describes date period.
type DatePeriod struct {
	From time.Time
	To   time.Time
}

// Total returns total amount of expenses in collection of records rr
// within time period p.
func Total(in []Record, p DatePeriod) float64 {
	panic("Total not implemented")
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses records.
func CategoryExpenses(in []Record, p DatePeriod, c string) (float64, error) {
	panic("CategoryExpenses not implemented")
}
