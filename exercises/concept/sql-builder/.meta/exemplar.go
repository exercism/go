package sqlbuilder

import "strings"

const (
	sqlSelect     = "SELECT "
	sqlFrom       = " FROM "
	sqlWildcard   = "*"
	sqlEnd        = ";"
	sqlInsertInto = "INSERT INTO "
	sqlValues     = " VALUES "

	separatorLength      = 2
	quotesLength         = 2
	columnBracketsLength = 3
	valueBracketsLength  = 4
)

type ColumnValue struct {
	Column string
	Value  string
}

func BuildSelectQuery(table string, columns []string) string {
	var sb strings.Builder

	totalSize := len(sqlSelect) + len(sqlFrom) + len(table) + len(sqlEnd)
	if len(columns) == 0 {
		totalSize += len(sqlWildcard)
	} else {
		for _, col := range columns {
			totalSize += len(col)
		}
		totalSize += (len(columns) - 1) * separatorLength
	}

	sb.Grow(totalSize)

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

func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	if len(rows) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(rows))
	var sb strings.Builder

	baseInsertSize := len(sqlInsertInto) + len(table) + columnBracketsLength + len(sqlValues) + valueBracketsLength

	maxRowSize := 0
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		currentSize := baseInsertSize
		for _, pair := range row {
			currentSize += len(pair.Column) + len(pair.Value) + quotesLength
		}

		if len(row) > 1 {
			currentSize += (len(row) - 1) * separatorLength * 2
		}

		if currentSize > maxRowSize {
			maxRowSize = currentSize
		}
	}

	if maxRowSize > 0 {
		sb.Grow(maxRowSize)
	}

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
