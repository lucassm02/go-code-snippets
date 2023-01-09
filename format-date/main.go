package main

import (
	"fmt"
	"time"
)

func showDate() {
	fmt.Println(time.Now().Format("02/01/2006 15:04:05"))
}

func main() {
	showDate()
}
