package sqlbuilder

import "strings"

type ColumnValue struct {
	Column string
	Value  string
}

func BuildSelectQuery(table string, columns []string) string {
	var sb strings.Builder

	sb.WriteString("SELECT ")
	if len(columns) == 0 {
		sb.WriteString("*")
	} else {
		sb.WriteString(strings.Join(columns, ", "))
	}

	sb.WriteString(" FROM ")
	sb.WriteString(table)
	sb.WriteString(";")

	return sb.String()
}

func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	if len(rows) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(rows))
	var sb strings.Builder

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		sb.Reset()

		sb.WriteString("INSERT INTO ")
		sb.WriteString(table)
		sb.WriteString(" (")
		for i, pair := range row {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(pair.Column)
		}
		sb.WriteString(")")

		sb.WriteString(" VALUES (")
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
