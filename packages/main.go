package main

import (
	"fmt"
	"go-code-snippets/packages/car"
)

// 1 - go mod init go-course/packages
// 2 - go mod tidy
// 3 - go get github.com/google/uuid

func main() {

	// fmt.Println(uuid.NewUUID())
	c := car.Car{Name: "Corola", Color: "Blue"}
	fmt.Println(c.GetColor())
}
