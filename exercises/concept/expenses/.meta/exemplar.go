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

// Records describes a collection of expenses events.
type Records []Record

// Filter creates a new records collection by applying predicate function to
// collection items and keeping the items when the function returns true.
func (rr Records) Filter(f func(Record) bool) Records {
	var out Records
	for _, r := range rr {
		if f(r) {
			out = append(out, r)
		}
	}
	return out
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

type categoryExpenses struct {
	Category string
	Total    float64
}

type byTotal []categoryExpenses

func (x byTotal) Len() int           { return len(x) }
func (x byTotal) Less(i, j int) bool { return x[i].Total < x[j].Total }
func (x byTotal) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byCategoryName []categoryExpenses

func (x byCategoryName) Len() int           { return len(x) }
func (x byCategoryName) Less(i, j int) bool { return x[i].Category < x[j].Category }
func (x byCategoryName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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

// Total returns total amount of expenses in collection a, within time period p.
func Total(rr Records, p DatePeriod) float64 {
	periodExpenses := rr.Filter(byDatePeriod(p))
	var total float64
	for _, r := range periodExpenses {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses collection a.
func CategoryExpenses(rr Records, p DatePeriod, c string) (float64, error) {
	categoryExpenses := rr.Filter(byCategory(c))
	if len(categoryExpenses) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	return Total(categoryExpenses, p), nil
}