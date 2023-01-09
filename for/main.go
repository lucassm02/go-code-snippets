package main

import "fmt"

func main() {

	value := 0
	for { // while true

		if value == 2 {
			break
		}

		fmt.Println("Hello")
		value++
	}

	fmt.Print("\n")

	var arr = []string{"H", "E", "L", "L", "O"}

	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}

	value2 := 0

	fmt.Print("\n")

	for value2 < 5 {

		if value2 == 3 {
			fmt.Println("Continue")
			value2++
			continue
		}

		fmt.Println("Hello")
		value2++
	}

	fmt.Print("\n")

	for _, value := range arr {
		fmt.Println(value)
	}

}
