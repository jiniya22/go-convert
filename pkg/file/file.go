package file

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func Create(filename string, data []string) {
	file, err := openFile(filename)
	if err != nil {
		panic(fmt.Sprintf("파일 생성중 오류 발생. filename: %s", filename))
	}
	defer file.Close()

	for _, d := range data {
		if _, err = file.WriteString(d + "\n"); err != nil {
			panic(err)
		}
	}
	fmt.Printf("파일 생성 완료 : %s\n", filename)
}

func openFile(filename string) (*os.File, error) {
	f := strings.Split(filename, ".")

	r, _ := regexp.Compile("[^0-9]")
	newFilename := f[0] + "_" + r.ReplaceAllString(time.Now().Format(time.DateTime), "") + "." + f[1]
	return os.OpenFile(newFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
}
