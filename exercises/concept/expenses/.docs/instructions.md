# Instructions

Bob is a financial adviser and helps people to manage their expenses. His clients send expenses records and Bob analyses them. Bob has records for the previous periods so that he can see changes in spending.

Bob needs to build two reports:
* total expenses in a date period
* total expenses in a category in a date period

In this exercise, you're going to build a program to help Bob.

All functions receive expenses records and a date period:
```go
// Expenses record.
type Record struct {
	Date        time.Time
	Amount      float64
	Category    string
}

// Expenses period.
type DatePeriod struct {
	From time.Time
	To   time.Time
}
```

The period should include both dates, i.e:
```go
p := DarePeriod{From: "2000-10-01", To: "2000-10-31"}
// p includes "2000-10-01"
// ...
// p includes "2000-10-16"
// ...
// p includes "2000-10-31"
//
// p does not include "2000-09-30"
// p does not include "2000-11-01"
```

The program operates with variable date periods and items categories. But the logic of calculation of the items sum does not depend on these variables. Therefore, program's functionality can be split into two tasks:
* Items filtration - by date period or category
* Calculation of items sum

## 1. Implement records filter
Implement the generic `Filter` function. The filter accepts a collection of items and a predicate function. It iterates through the collection and applies a predicate to each item. When the predicate function returns true, the item is added to a new collection. Then the filter function returns the collection of filtered items.

## 2. Implement filter predicate "ByDatePeriod"
Implement the `ByDatePeriod` filter predicate. This function accepts date period and returns function of expenses record. Returned function returns true when record date is within provided date period.

```go
import "time"

from, _ := time.Parse("2006-01-02", "2000-10-01")
to, _ := time.Parse("2006-01-02", "2000-10-15")
period := DatePeriod{From: from, To: to}

records := []Record{
  // 2000-10-01
  {Date: time.Date(2000, time.October, 1, 0, 0, 0, 0, time.UTC), Amount: 15, Category: "grocieries"},
  // 2000-10-11
  {Date: time.Date(2000, time.October, 11, 0, 0, 0, 0, time.UTC), Amount: 300, Category: "utility-bills"},
  // 2000-10-12
  {Date: time.Date(2000, time.October, 12, 0, 0, 0, 0, time.UTC), Amount: 28, Category: "grocieries"},
  // 2000-10-26
  {Date: time.Date(2000, time.October, 26, 0, 0, 0, 0, time.UTC), Amount: 300, Category: "university"},
  // 2000-10-28
  {Date: time.Date(2000, time.October, 28, 0, 0, 0, 0, time.UTC), Amount: 1300, Category: "rent"},
}

Filter(records, ByDatePeriod(period))
// Output:
// [
//   {Date: time.Date(2000, time.October, 1, 0, 0, 0, 0, time.UTC), Amount: 15, Category: "grocieries"},
//   {Date: time.Date(2000, time.October, 11, 0, 0, 0, 0, time.UTC), Amount: 300, Category: "utility-bills"},
// ]
```

## 3. Implement filter predicate "ByCategory"
Implement the `ByCategory` filter predicate. This function accepts category and returns function of expenses record. Returned function returns true when record category is equal to a provided category.

```go
Filter(records, ByCategory("grocieries"))
// Output:
// [
//   {Date: time.Date(2000, time.October, 1, 0, 0, 0, 0, time.UTC), Amount: 15, Category: "grocieries"},
//   {Date: time.Date(2000, time.October, 12, 0, 0, 0, 0, time.UTC), Amount: 28, Category: "grocieries"},
// ]
```

## 4. Calculate the total amount of expenses in period
Implement the `Total` function to return a sum of expenses in the date period:

```go
import "time"

from, _ := time.Parse("2006-01-02", "2000-10-01")
to, _ := time.Parse("2006-01-02", "2000-10-31")
oct2000 := DatePeriod{From: from, To: to}

from, _ = time.Parse("2006-01-02", "2000-11-01")
to, _ = time.Parse("2006-01-02", "2000-11-30")
nov2000 := DatePeriod{From: from, To: to}

recordDate, _ := time.Parse("2006-01-02", "2000-10-11")
records := []Record{{Date: recordDate, Amount: 16, Category: "entertainment"}}

Total(records, oct2000)
// Output: 16

Total(records, nov2000)
// Output: 0
```

## 5. Calculate the total amount of category expenses in period
Implement the `CategoryExpenses` function to return the category's expenses in the period. The function should differentiate a case when a category is not present in the expenses records and the case when there are no category's expenses in the provided period.
In case, when the category is not present the function should return an error.

```go
from, _ := time.Parse("2006-01-02", "2000-10-01")
to, _ := time.Parse("2006-01-02", "2000-10-31")
oct2000 := DatePeriod{From: from, To: to}

from, _ = time.Parse("2006-01-02", "2000-11-01")
to, _ = time.Parse("2006-01-02", "2000-11-30")
nov2000 := DatePeriod{From: from, To: to}

records := []Record{
  // 2000-10-01
  {Date: time.Date(2000, time.October, 1, 0, 0, 0, 0, time.UTC), Amount: 15, Category: "grocieries"},
  // 2000-10-11
  {Date: time.Date(2000, time.October, 11, 0, 0, 0, 0, time.UTC), Amount: 300, Category: "utility-bills"},
  // 2000-10-12
  {Date: time.Date(2000, time.October, 12, 0, 0, 0, 0, time.UTC), Amount: 28, Category: "grocieries"},
  // 2000-10-26
  {Date: time.Date(2000, time.October, 26, 0, 0, 0, 0, time.UTC), Amount: 300, Category: "university"},
  // 2000-10-28
  {Date: time.Date(2000, time.October, 28, 0, 0, 0, 0, time.UTC), Amount: 1300, Category: "rent"},
}

CategoryExpenses(records, oct2000, "entertainment")
// Output: 0, error(unknown category entertainment)

CategoryExpenses(records, oct2000, "rent")
// Output: 1300, nil

CategoryExpenses(records, nov2000, "rent")
// Output: 0, nil
```
