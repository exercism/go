# Hints

## 1. Total expenses in period

- Use `Records.Filter` to select expenses records in period.
- Implement a function that takes period as a parameter and returns `func(Record) bool`. Returned function can be used in `Records.Filter`.

## 2. Top N categories of expenses in period

- First select expenses in period, reuse the filter.
- Group total expenses by category. The result of the grouping will be a collection of structures with two fields: category name and total expenses.
- Use [`sort`][sort-pkg] package to sort a collection of category expenses.
- Use [example][sort-example] as reference for working with `sort` package.

## 3. Category expenses in period

- Implement filter predicate `func byCategory(string) func(Record) bool` that takes category and returns a predicate function. The predicate function returns `true` when Record.Category matches provided category.

[sort-pkg]: https://pkg.go.dev/sort
[sort-example]: https://gobyexample.com/sorting-by-functions