package main

import (
	"fmt"

	"github.com/google/uuid"
)

// 1 - go mod init go-course/packages
// 2 - go mod tidy
// 3 - go get github.com/google/uuid

func main() {

	fmt.Println(uuid.NewUUID())

}
