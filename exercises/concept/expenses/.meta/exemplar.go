package expenses

import (
	"fmt"
	"sort"
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

// Period describes time period.
type Period struct {
	DateFrom time.Time
	DateTo   time.Time
}

func (p Period) Includes(d time.Time) bool {
	return (p.DateFrom.Before(d) || p.DateFrom.Equal(d)) &&
		(p.DateTo.After(d) || p.DateTo.Equal(d))
}

type categoryExpenses struct {
	Category string
	Total    float64
}

type byTotal []categoryExpenses

func (x byTotal) Len() int           { return len(x) }
func (x byTotal) Less(i, j int) bool { return x[i].Total < x[j].Total }
func (x byTotal) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func byPeriod(p Period) func(Record) bool {
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
func Total(rr Records, p Period) float64 {
	periodExpenses := rr.Filter(byPeriod(p))
	var total float64
	for _, r := range periodExpenses {
		total += r.Amount
	}
	return total
}

// TopCategoriesN returns top n categories of expenses in collection a, within
// time period p.
func TopCategoriesN(rr Records, p Period, n int) []string {
	if n < 1 {
		return nil
	}

	periodExpenses := rr.Filter(byPeriod(p))

	// reduce records to categoryExpenses collection
	catExp := map[string]float64{}
	for _, r := range periodExpenses {
		catExp[r.Category] += r.Amount
	}
	catsExps := make([]categoryExpenses, 0, len(catExp))
	for category, total := range catExp {
		catsExps = append(catsExps, categoryExpenses{category, total})
	}

	// sort categoryExpenses
	sort.Sort(sort.Reverse(byTotal(catsExps)))

	// map only the category names
	categories := make([]string, 0, n)
	for _, v := range catsExps[:n] {
		categories = append(categories, v.Category)
	}

	return categories
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses collection a.
func CategoryExpenses(rr Records, p Period, c string) (float64, error) {
	categoryExpenses := rr.Filter(byCategory(c))
	if len(categoryExpenses) == 0 {
		return 0, fmt.Errorf("unknown category: %s", c)
	}
	return Total(categoryExpenses, p), nil
}
