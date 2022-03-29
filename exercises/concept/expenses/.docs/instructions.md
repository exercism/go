# Instructions

Bob is a financial adviser and helps people to manage their expenses. His clients send expenses records and Bob analyses them. Bob has records for the previous periods so that he can see changes in spending. Bob does not like calendars and uses *Bob epoch* instead of date - a number of days that have elapsed since Bob's client started their activity.

Bob needs to build two reports:
* total expenses in a days period
* total expenses in a category in a days period

In this exercise, you're going to build a program to help Bob.

All functions receive expenses records and a days period:

```go
// Record represents an expense record.
type Record struct {
	Day         int
	Amount      float64
	Category    string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}
```

The period includes both days, i.e:

```go
p := DaysPeriod{From: 1, To: 31}
// p includes 1
// ...
// p includes 16
// ...
// p includes 31
//
// p does not include 40
// p does not include 50
```

The program operates with variable days periods and items categories. But the logic of calculation of the items sum does not depend on these variables. Therefore, program's functionality can be split into two tasks:
* Items filtration - by days period or category
* Calculation of items sum

## 1. Implement records filter
Implement the generic `Filter` function. The filter accepts a collection of items and a predicate function. It iterates through the collection and applies a predicate to each item. When the predicate function returns true, the item is added to a new collection. Then the filter function returns the collection of filtered items.

## 2. Implement filter predicate "ByDaysPeriod"
Implement the `ByDaysPeriod` filter predicate. This function accepts days period and returns function of expenses record. Returned function returns true when record day is within provided days period.

```go
period := DaysPeriod{From: 1, To: 15}

records := []Record{
  {Day: 1, Amount: 15, Category: "grocieries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "grocieries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

Filter(records, ByDaysPeriod(period))
// Output:
// [
//   {Day: 1, Amount: 15, Category: "grocieries"},
//   {Day: 11, Amount: 300, Category: "utility-bills"},
//   {Day: 12, Amount: 28, Category: "grocieries"},
// ]
```

## 3. Implement filter predicate "ByCategory"
Implement the `ByCategory` filter predicate. This function accepts category and returns function of expenses record. Returned function returns true when record category is equal to a provided category.

```go
Filter(records, ByCategory("grocieries"))
// Output:
// [
//   {Day: 1, Amount: 15, Category: "grocieries"},
//   {Day: 12, Amount: 28, Category: "grocieries"},
// ]
```

## 4. Calculate the total amount of expenses in period
Implement the `Total` function to return a sum of expenses in the days period:

```go
p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

records := []Record{{Day: 15, Amount: 16, Category: "entertainment"}}

Total(records, p1)
// Output: 16

Total(records, p2)
// Output: 0
```

## 5. Calculate the total amount of category expenses in days period
Implement the `CategoryExpenses` function to return the category's expenses in the days period. The function should differentiate a case when a category is not present in the expenses records and the case when there are no category's expenses in the provided period.
In case, when the category is not present the function should return an error.

```go
p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

records := []Record{
  {Day: 1, Amount: 15, Category: "grocieries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "grocieries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

CategoryExpenses(records, p1, "entertainment")
// Output: 0, error(unknown category entertainment)

CategoryExpenses(records, p1, "rent")
// Output: 1300, nil

CategoryExpenses(records, p2, "rent")
// Output: 0, nil
```
