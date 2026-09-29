package sqlbuilder

import (
	"strings"
)

// SQL syntax components used to construct the query.
const (
	sqlSelect     = "SELECT "
	sqlFrom       = " FROM "
	sqlWildcard   = "*"
	sqlEnd        = ";"
	sqlInsertInto = "INSERT INTO "
	sqlValues     = " VALUES "
)

type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query from the given table name
// and a slice of column names.
//
// Performance optimizations:
//   - Uses strings.Builder to prevent unnecessary string allocations in the heap.
//   - Pre-calculates the exact buffer size in bytes to avoid dynamic re-allocations during execution.
//
// Edge cases handled:
//   - If the columns slice is empty or nil, it automatically falls back to the "*" wildcard.
func BuildSelectQuery(table string, columns []string) string {
	var builder strings.Builder

	// Calculate the exact size of the destination buffer in bytes.
	// Initial size includes base keywords ("SELECT ", " FROM ", ";") and the table name.
	totalSize := len(sqlSelect) + len(sqlFrom) + len(sqlEnd) + len(table)

	if len(columns) == 0 {
		totalSize += len(sqlWildcard)
	} else {
		for i, col := range columns {
			totalSize += len(col)
			if i > 0 {
				totalSize += 2 // Account for the separating comma and space (", ")
			}
		}
	}

	// Pre-allocate memory once to completely avoid overhead from dynamic buffer growing.
	builder.Grow(totalSize)

	// Assemble the SQL string sequentially into the buffer without intermediate allocations.
	builder.WriteString(sqlSelect)

	if len(columns) == 0 {
		builder.WriteString(sqlWildcard)
	} else {
		for i, col := range columns {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(col)
		}
	}

	builder.WriteString(sqlFrom)
	builder.WriteString(table)
	builder.WriteString(sqlEnd)

	return builder.String()
}

// BuildBatchInserts generates a slice of SQL INSERT statements for the given rows.
// It accepts a slice of rows, where each row is a structured slice of ColumnValue pairs,
// ensuring a deterministic and predictable execution order.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	if len(rows) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(rows))
	var builder strings.Builder

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		// Reset the builder buffer to start fresh for each statement.
		// This retains the underlying allocated capacity across loop iterations.
		builder.Reset()

		// Build the column definitions part: INSERT INTO table (col1, col2)
		builder.WriteString(sqlInsertInto)
		builder.WriteString(table)
		builder.WriteString(" (")
		for i, pair := range row {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(pair.Column)
		}
		builder.WriteString(")")

		// Build the values mapping part: VALUES ('val1', 'val2');
		builder.WriteString(sqlValues)
		builder.WriteString("(")
		for i, pair := range row {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteRune('\'')
			builder.WriteString(pair.Value)
			builder.WriteRune('\'')
		}
		builder.WriteString(");")

		result = append(result, builder.String())
	}

	return result
}
