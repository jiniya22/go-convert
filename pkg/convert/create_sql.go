package convert

import (
	"fmt"
	"go-convert/pkg/lang"
	"strings"
)

func CreateSql(tableName string, columnNames []string, rows [][]string) []string {
	var result = make([]string, len(rows))
	var builder strings.Builder
	prefix := insertStatementPrefix(tableName, columnNames)

	for i, row := range rows {
		builder.WriteString(prefix)
		for j, col := range row {
			if j == 0 {
				builder.WriteString("(")
			}
			builder.WriteString(getData(col))
			if j < len(row)-1 {
				builder.WriteString(",")
			} else {
				builder.WriteString(")")
			}
		}
		builder.WriteString(";")
		result[i] = builder.String()
		builder.Reset()
	}

	return result
}

func insertStatementPrefix(tableName string, columnNames []string) string {
	return fmt.Sprintf("INSERT INTO `%s` (%v) values", tableName, strings.Join(columnNames, ","))
}

func getData(s string) string {
	if lang.IsInt(s) || lang.IsFloat(s) {
		return fmt.Sprintf("%s", s)
	} else if s == "NULL" {
		return "NULL"
	}
	return fmt.Sprintf("\"%s\"", s)
}
