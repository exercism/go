# Hints

## 1. Total expenses in period

- Implement filter predicate `func byPeriod(Period) func(Record) bool` that takes period and returns a predicate function. The predicate function returns `true` when Record.Date is within the period.

## 2. Top N categories of expenses in period

- Use `sort.Interface` to sort expenses records.

## 3. Category expenses in period

- Implement filter predicate `func byCategory(string) func(Record) bool` that takes category and returns a predicate function. The predicate function returns `true` when Record.Category matches provided category.
