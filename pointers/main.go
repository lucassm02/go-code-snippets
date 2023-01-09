package main

import (
	"fmt"
)

func setValeu(value **int) int {

	**value = **value + 1

	return **value

}

func main() {

	value1 := 10

	fmt.Println(value1) // 10

	value2 := &value1

	fmt.Print("\n")
	fmt.Println(value2)  // 0xc0000180a0
	fmt.Println(*value2) // 10

	value1 = 20

	fmt.Print("\n")
	fmt.Println(value1)  // 20
	fmt.Println(*value2) // 20

	*value2 = 30

	fmt.Print("\n")
	fmt.Println(value1)  // 30
	fmt.Println(*value2) // 30

	var value3 *int = &value1

	fmt.Println(*value3) // 30

	setValeu(&value3)

	fmt.Print("\n")
	fmt.Println(&value1) // 31
	fmt.Println(*value2) // 31
	fmt.Println(*value3) // 31

}
