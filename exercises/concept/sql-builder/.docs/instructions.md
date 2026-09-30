# Instructions

You are building a lightweight SQL query generator utility package for a database migration tool.
To ensure the application scales well and handles large volumes of text without crushing the memory heap, you must implement efficient string construction techniques using Go's `strings.Builder`.

## 1. Generate SELECT Queries

Implement the `BuildSelectQuery` function.
It takes a table name (`string`) and a slice of column names (`[]string`).
It should return a fully formed SQL `SELECT` statement ending with a semicolon.

If the columns slice is empty or `nil`, the query should fall back to using the `*` wildcard to select all columns.

```go
columns := []string{"id", "name", "email"}
BuildSelectQuery("users", columns)
// Output: "SELECT id, name, email FROM users;"

BuildSelectQuery("orders", nil)
// Output: "SELECT * FROM orders;"
```

## 2. Generate Batch INSERT Statements

Implement the `BuildBatchInserts` function.
It takes a table name (`string`) and a two-dimensional slice representing database rows (`[][]ColumnValue`).
Each element in a row links a specific column name to its string value.

The function must return a slice of individual `INSERT` SQL strings.
The structured slice format guarantees a predictable and stable query output order.

If the input slice of rows is empty, return an empty string slice.
If any individual row is empty, skip it.
String values within the `VALUES` clause must be enclosed in single quotes `'`.

```go
rows := [][]ColumnValue{
    {
        {Column: "title", Value: "Book"},
        {Column: "price", Value: "100"},
    },
    {
        {Column: "title", Value: "Magazine"},
        {Column: "price", Value: "15"},
    },
}

BuildBatchInserts("products", rows)
// Output: [
//   "INSERT INTO products (title, price) VALUES ('Book', '100');",
//   "INSERT INTO products (title, price) VALUES ('Magazine', '15');"
// ]
```

## Performance Requirements

- Avoid using standard string concatenation (`+` or `+=`) inside loops.
- Maximize performance by taking advantage of `strings.Builder` and its memory allocation optimization patterns (`Grow` and `Reset`).
