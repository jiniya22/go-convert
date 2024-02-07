package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ParseCsv(filepath string) ([]string, [][]string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("파일이 존재하지 않습니다. filename: %s", filepath))
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var rows [][]string
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		rows = append(rows, row)
	}
	return rows[0], rows[1:]
}
