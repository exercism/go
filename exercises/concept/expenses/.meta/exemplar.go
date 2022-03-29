package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
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

// ByDaysPeriod returns predicate function. The predicate returns true when
// Record.Day is within the days period.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return p.Includes(r.Day)
	}
}

// ByCategory returns predicate function. The predicate returns true when
// Record.Category is equal to the provided category.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

func (p DaysPeriod) Includes(d int) bool {
	return p.From <= d && d <= p.To
}

// Total returns total amount of expenses in collection of records within
// days period p.
func Total(in []Record, p DaysPeriod) float64 {
	periodExpenses := Filter(in, ByDaysPeriod(p))
	var total float64
	for _, r := range periodExpenses {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses collection a.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	categoryExpenses := Filter(in, ByCategory(c))
	if len(categoryExpenses) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return Total(categoryExpenses, p), nil
}
