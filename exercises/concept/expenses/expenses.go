package expenses

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter creates a new records collection by applying the predicate function to
// records and keeping the records for which the function returns true.
func Filter(in []Record, f func(Record) bool) []Record {
	panic("Please implement the Filter function.")
}

// ByDaysPeriod returns predicate function. The predicate returns true when
// Record.Day is within the days period.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	panic("Please implement the ByDaysPeriod function.")
}

// ByCategory returns predicate function. The predicate returns true when
// Record.Category is equal to the provided category.
func ByCategory(c string) func(Record) bool {
	panic("Please implement the ByCategory function.")
}

// Total returns total amount of expenses in collection of records within
// days period p.
func Total(in []Record, p DaysPeriod) float64 {
	panic("Please implement the Total function.")
}

// CategoryExpenses returns total amount of expenses in category c. It returns
// error when a category is not present in expenses records.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	panic("Please implement the CategoryExpenses function.")
}
