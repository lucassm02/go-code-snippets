package main

import "fmt"

func main() {

	value := 10

	// if 10 { error, only boolean values
	// }

	if value == 10 {
		fmt.Println("True")
	} else if value < 10 {
		fmt.Println("False")
	} else {
		fmt.Println("False")
	}

	if value2 := 20; value < 10 {
		fmt.Println(value2)
	}

}
