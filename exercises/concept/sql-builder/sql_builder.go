package sqlbuilder

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it should default to the wildcard "*".
//
// You should practice using the zero-value of strings.Builder, WriteString, and String().
func BuildSelectQuery(table string, columns []string) string {
	panic("Please implement BuildSelectQuery")
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows should be skipped, and values must be enclosed in single quotes.
//
// You should practice reusing a SINGLE strings.Builder instance and clearing its buffer
// between iterations using the Reset() method for optimal performance.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	panic("Please implement BuildBatchInserts")
}
