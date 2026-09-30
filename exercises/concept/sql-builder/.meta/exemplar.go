package sqlbuilder

import (
	"strings"
)

const (
	sqlSelect     = "SELECT "
	sqlFrom       = " FROM "
	sqlWildcard   = "*"
	sqlEnd        = ";"
	sqlInsertInto = "INSERT INTO "
	sqlValues     = " VALUES "

	estimatedColumnLength          = 15
	estimatedInsertStatementLength = 128
)

type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query from the given table name
// and a slice of column names.
func BuildSelectQuery(table string, columns []string) string {
	var sb strings.Builder

	sb.Grow(len(sqlSelect) + len(sqlFrom) + len(table) + len(sqlEnd) + len(columns)*estimatedColumnLength)

	sb.WriteString(sqlSelect)

	if len(columns) == 0 {
		sb.WriteString(sqlWildcard)
	} else {
		for i, col := range columns {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(col)
		}
	}

	sb.WriteString(sqlFrom)
	sb.WriteString(table)
	sb.WriteString(sqlEnd)

	return sb.String()
}

// BuildBatchInserts generates a slice of SQL INSERT statements for the given rows.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	if len(rows) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(rows))
	var sb strings.Builder

	sb.Grow(estimatedInsertStatementLength)
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		sb.WriteString(sqlInsertInto)
		sb.WriteString(table)
		sb.WriteString(" (")
		for i, pair := range row {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(pair.Column)
		}
		sb.WriteString(")")

		sb.WriteString(sqlValues)
		sb.WriteString("(")
		for i, pair := range row {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteByte('\'')
			sb.WriteString(pair.Value)
			sb.WriteByte('\'')
		}
		sb.WriteString(");")

		result = append(result, sb.String())
	}

	return result
}
