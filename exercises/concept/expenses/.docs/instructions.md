# Instructions

Bob is a financial adviser and helps people to manage their expenses. Bob's clients send expenses records for him to analyze. Bob has records for the previous periods so he can see changes in spending. Bob does not like calendars and uses *Bob epoch* instead of dates. Bob epoch is the number of days elapsed since Bob's client started their activity.

In this exercise, you are going to build a program to help Bob manage and analyze the expenses of his clients.

Bob works with `Record`s and `DaysPeriod`s.

A `Record` represents an expense record that contains the day in which the expense was made, the money spent, and the category of the expense.

```go
// Record represents an expense record.
type Record struct {
	Day         int
	Amount      float64
	Category    string
}
```

A `DaysPeriod` represents a range of days and includes all days from the day `From` up to the day `To`.
Both ends are included in the range.

```go
// DaysPeriod represents a period of days.
type DaysPeriod struct {
	From int
	To   int
}

p := DaysPeriod{From: 1, To: 31}
// p represents all days from the day 1 to the day 31:
//  - days 1, 20, 16 and 31 are examples of days that are included
//    in the range of time specified by p
//  - days 50 and 40 are examples of days that are not included
//    in the range of time specified by p  
```

## 1. Implement a general records filter

Bob deals with a lot of records every day, but not all of them are interesting depending on the analysis Bob is making.
Let's help Bob perform some basic filtering of records.

Implement the generic `Filter` function to filter records according to a criteria given by a function.
This filter function accepts a collection of records and a predicate function and returns only the records in the collection that satisfy the predicate.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
  return r.Day == 1
}

Filter(records, Day1Records)
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"}
// ]
```

## 2. Filter records within a period of time

Bob frequently needs to filter records that are in a given period of time.

Implement the `ByDaysPeriod` function that will help Bob create such filters.
This function accepts a `DaysPeriod` and returns function that takes a record and tells whether the record is in the period of time specified by the `DaysPeriod` given as argument.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

period := DaysPeriod{From: 1, To: 15}

Filter(records, ByDaysPeriod(period))
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 11, Amount: 300, Category: "utility-bills"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]
```

## 3. Filter records by category

Other than filtering records by a period of time, Bob also needs to filter records by its category.

Implement the `ByCategory` function that will help Bob create such filters.
This function accepts a category and returns a function that takes a record and tells whether the category of this record is the same as the category given as the argument.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

Filter(records, ByCategory("groceries"))
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]
```

## 4. Calculate the total amount of expenses in a period

Bob also wants to know the total amount of the expenses in a period of time.

Implement the `TotalByPeriod` function to return a sum of expenses in the days period.

```go
records := []Record{
  {Day: 15, Amount: 16, Category: "entertainment"},
  {Day: 32, Amount: 20, Category: "groceries"},
  {Day: 40, Amount: 30, Category: "entertainment"}
}

p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

TotalByPeriod(records, p1)
// => 16

TotalByPeriod(records, p2)
// => 50
```

## 5. Calculate the total expenses for records of a category in a period

For the most complex reports Bob makes to his clients, Bob needs to filter records by category and period of time at the same time. 
That means Bob wants to know the total expenses for records in a category in a given period of time.

Implement the `CategoryExpenses` function that returns the total amount of expenses in a category in a given period of days. The function should also differentiate the case when the given category is not present in the expenses records and the case when there are no category's expenses in the provided period.
When the category is not a category of any of the records (regardless of period of time) the function should return an error.

```go
p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

CategoryExpenses(records, p1, "entertainment")
// => 0, error(unknown category entertainment)

CategoryExpenses(records, p1, "rent")
// => 1300, nil

CategoryExpenses(records, p2, "rent")
// => 0, nil
```
