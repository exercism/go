package expenses

import (
	"fmt"
	"time"
)

// Record represents an expense record.
type Record struct {
	Date     time.Time
	Amount   float64
	Category string
}

// Filter creates a new records collection by applying predicate function to
// collection items and keeping the items when the function returns true.
func Filter(in []Record, f func(Record) bool) []Record {
	var out []Record
	for _, r := range in {
		if f(r) {
			out = append(out, r)
		}
	}
	return out
}

// ByDatePeriod returns predicate function. The predicate returns true when
// Record.Date is within the date period.
func ByDatePeriod(p DatePeriod) func(Record) bool {
	return func(r Record) bool {
		return p.Includes(r.Date)
	}
}

// ByCategory returns predicate function. The predicate returns true when
// Record.Category is equal to the provided category.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// DatePeriod represents a period of time for expenses.
type DatePeriod struct {
	From time.Time
	To   time.Time
}

func (p DatePeriod) Includes(d time.Time) bool {
	return (p.From.Before(d) || p.From.Equal(d)) &&
		(p.To.After(d) || p.To.Equal(d))
}

// Total returns total amount of expenses in collection a, within time period p.
func Total(in []Record, p DatePeriod) float64 {
	periodExpenses := Filter(in, ByDatePeriod(p))
	var total float64
	for _, r := range periodExpenses {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses collection a.
func CategoryExpenses(in []Record, p DatePeriod, c string) (float64, error) {
	categoryExpenses := Filter(in, ByCategory(c))
	if len(categoryExpenses) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return Total(categoryExpenses, p), nil
}
