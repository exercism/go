# Hints

## 1. Total expenses in period

- Use `Records.Filter` to select expenses records in period.
- Implement a function that takes period as a parameter and returns `func(Record) bool`. Returned function can be used in `Records.Filter`.

## 2. Category expenses in period

- Similarly to filtering records by period, the records can be filtered by category. Implement a new functino that takes category as a parameter and returns `func(Record) bool`.
- Use `Records.Filter` to select expenses in category.
- Reuse `Total` to calculate category expenses.
