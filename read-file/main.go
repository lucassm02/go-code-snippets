package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Read steam of bytes
func exampleOne(filePath string) {
	fmt.Println("Example One")

	file, _ := os.Open(filePath)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		fmt.Println(strings.TrimSpace(line))

		if err == io.EOF {
			break
		}
	}

	file.Close()

	// All file content
	// file, _ := ioutil.ReadFile("file.txt")

	// fmt.Println(string(file))

}

// Put all file in memory
func exampleTwo(filePath string) {

	fmt.Println("Example two")

	file, err := ioutil.ReadFile("file.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(file))

}

var filePath string = "file.txt"

func main() {
	exampleOne(filePath)
	exampleTwo(filePath)
}
