package main

import (
	"go-convert/pkg/convert"
	"go-convert/pkg/file"
	"go-convert/pkg/parser"
)

func main() {
	tableName := "user"
	filepath := "resources/data.csv"
	columNames, rows := parser.ParseCsv(filepath)
	sql := convert.CreateSql(tableName, columNames, rows)
	file.Create("result/data.sql", sql)
}
