package main

import (
	"fmt"
	"os"
)

func writeFile(fileName string) {
	file, err := os.OpenFile("teste.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString("Line one\n")
	file.WriteString("Line two\n")

	file.Close()
}

func main() {
	writeFile("file.txt")
}
