package main

import "fmt"

func main() {
	name := "Lucas"

	switch name {
	case "Lucas":
		fmt.Println("Hello Lucas!")
	default:
		fmt.Println("Hello anonymous!")
	}
}
