package sqlbuilder

// BuildSelectQuery constructs a basic "SELECT col1, col2 FROM table" query.
// It allows practicing strings.Builder zero-value, WriteString, and String().
func BuildSelectQuery(table string, columns []string) string {
	panic("Please implement BuildSelectQuery")
}

// BuildBatchInserts constructs multiple INSERT statements using a SINGLE reused builder.
// It allows practicing the Reset() method to clean the buffer between iterations.
func BuildBatchInserts(table string, rows []map[string]string) []string {
	panic("Please implement BuildBatchInserts")
}
