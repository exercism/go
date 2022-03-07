package expenses

import (
	"fmt"
	"time"
)

// Record describes a record of expense event.
type Record struct {
	Date     time.Time
	Amount   float64
	Category string
}

// Filter creates a new records collection by applying predicate function to
// collection items and keeping the items when the function returns true.
func Filter(in []Record, f func(Record) bool) (out []Record) {
	for _, r := range in {
		if f(r) {
			out = append(out, r)
		}
	}
	return out
}

func byDatePeriod(p DatePeriod) func(Record) bool {
	return func(r Record) bool {
		return p.Includes(r.Date)
	}
}

func byCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// DatePeriod describes time period.
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
	periodExpenses := Filter(in, byDatePeriod(p))
	var total float64
	for _, r := range periodExpenses {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses collection a.
func CategoryExpenses(in []Record, p DatePeriod, c string) (float64, error) {
	categoryExpenses := Filter(in, byCategory(c))
	if len(categoryExpenses) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return Total(categoryExpenses, p), nil
}
