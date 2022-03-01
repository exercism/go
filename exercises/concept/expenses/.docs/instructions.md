# Instructions

Bob is a financial adviser and helps people to manage their expenses. His clients send expenses records and Bob analyses them. Bob has records for the previous periods so that he can see changes in spending.

Bob needs to build a report that contains the top 3 categories of expenses and the total amount of the expenses. Sometimes customers are interested to see expenses in a particular category.

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

## 1. Calculate the total amount of expenses in period

Implement the `Total` function to return a sum of expenses in period:

```go
oct2000 := DatePeriod{From: "2000-10-01", To: "2000-10-31"}
nov2000 := DatePeriod{From: "2000-11-01", To: "2000-11-30"}
records := []Record{{Date: "2000-10-11", Amount: 16, Category: "entertainment"}}

Total(records, oct2000)
// Output: 16

Total(records, nov2000)
// Output: 0
```

## 2. Calculate the total amount of category expenses in period

Implement the `CategoryExpenses` function to return the category's expenses in the period. The function should differentiate a case when a category is not present in the expenses records and the case when there are no category's expenses in the provided period.
In case, when the category is not present the function should return an error.

```go
oct2000 := DatePeriod{From: "2000-10-01", To: "2000-10-31"}
nov2000 := DatePeriod{From: "2000-11-01", To: "2000-11-30"}
records := []Record{
  {Date: "2000-10-01", Amount: 15, Category: "grocieries"},
  {Date: "2000-10-11", Amount: 300, Category: "utility-bills"},
  {Date: "2000-10-12", Amount: 28, Category: "grocieries"},
  {Date: "2000-10-26", Amount: 300, Category: "university"},
  {Date: "2000-10-28", Amount: 1300, Category: "rent"},
}

CategoryExpenses(records, oct2000, "entertainment")
// Output: 0, error(unknown category entertainment)

CategoryExpenses(records, oct2000, "rent")
// Output: 1300, nil

CategoryExpenses(records, nov2000, "rent")
// Output: 0, nil
```
