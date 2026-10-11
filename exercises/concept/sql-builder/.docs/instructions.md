# Instructions

Andrew, a developer at his company, recently built a lightweight SQL query generator utility for a database migration tool.
His first naive implementation worked perfectly for small sets of data.
However, during large production migrations, the application started consuming excessive memory and struggled under heavy load.
After profiling the application, Andrew realized that standard string concatenation inside loops was forcing Go to constantly reallocate memory.
Now, he needs your help to refactor the utility using Go's `strings.Builder` to ensure the application scales well and handles large volumes of text without crushing the memory heap.

## 1. Generate SELECT Queries

Implement the `BuildSelectQuery` function.
It takes a table name (`string`) and a slice of column names (`[]string`).
It should return a fully formed SQL `SELECT` statement ending with a semicolon.
If the columns slice is empty or `nil`, the query should fall back to using the `*` wildcard to select all columns.

```go
columns := []string{"id", "name", "email"}
BuildSelectQuery("users", columns)
// Function return: "SELECT id, name, email FROM users;"

BuildSelectQuery("orders", nil)
// Function return: "SELECT * FROM orders;"
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
// Function return strings slice: [
//   "INSERT INTO products (title, price) VALUES ('Book', '100');",
//   "INSERT INTO products (title, price) VALUES ('Magazine', '15');"
// ]
```
