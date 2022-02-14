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
type Period struct {
	DateFrom time.Time
	DateTo   time.Time
}
```

The period should include both dates, i.e:
```go
p := Period{DateFrom: "2000-10-01", DateTo: "2000-10-31"}
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
Total([]Record{{Date: "2000-10-11", Amount: 16, Category: "entertainment"}},
  Period{DateFrom: "2000-10-01", DateTo: "2000-10-31"})
// Output: 16

Total([]Record{{Date: "2000-10-11", Amount: 16, Category: "entertainment"}},
  Period{DateFrom: "2000-11-01", DateTo: "2000-11-30"})
// Output: 0
```

## 2. Find the top N expenses categories in period

Implement the `TopCategoriesN` function to return the categories that are responsible for the most spending.

The categories should be sorted by the sum of expenses in descending order, i.e the most expensive category on the top of the list, etc. Categories with the same sum of expenses should be ordered alphabetically.

```go
records := []Record{
  {Date: "2000-10-11", Amount: 300, Category: "utility-bills"},
  {Date: "2000-10-26", Amount: 300, Category: "university"},
  {Date: "2000-10-28", Amount: 1300, Category: "rent"},
}
period := Period{DateFrom: "2000-10-01", DateTo: "2000-10-31"}

TopCategoriesN(records, period, 1)
// Output: [rent]

TopCategoriesN(records, period, 2)
// Output: [rent, university]

TopCategoriesN(records, period, 3)
// Output: [rent, university, utility-bills]

TopCategoriesN(records, period, 4)
// Output: [rent, university, utility-bills]
```

If there are less categories than parameter `N` the function should output all categories. Function should ignore non-positive `N` and output `nil` in such cases.

```go
TopCategoriesN(records, period, -1)
// Output: nil

TopCategoriesN(records, period, 0)
// Output: nil
```

## 3. Calculate the total amount of category expenses in period

Implement the `CategoryExpenses` function to return the category's expenses in the period. The function should differentiate a case when a category is not present in the expenses records and the case when there are no category's expenses in the provided period.
In case, when the category is not present the function should return an error.

```go
records := []Record{
  {Date: "2000-10-01", Amount: 15, Category: "grocieries"},
  {Date: "2000-10-11", Amount: 300, Category: "utility-bills"},
  {Date: "2000-10-12", Amount: 28, Category: "grocieries"},
  {Date: "2000-10-26", Amount: 300, Category: "university"},
  {Date: "2000-10-28", Amount: 1300, Category: "rent"},
}
october := Period{DateFrom: "2000-10-01", DateTo: "2000-10-31"}
november := Period{DateFrom: "2000-11-01", DateTo: "2000-11-30"}

CategoryExpenses(records, october, "entertainment")
// Output: 0, error(unknown category entertainment)

CategoryExpenses(records, october, "rent")
// Output: 1300, nil

CategoryExpenses(records, november, "rent")
// Output: 0, nil
```
