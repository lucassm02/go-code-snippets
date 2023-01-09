package main

import "fmt"

func scanInput() string {
	var inputValue string

	// fmt.Scanf("%d", &inputValue)
	fmt.Scan(&inputValue)

	return inputValue
}

func main() {
	fmt.Println("Digite qualquer tecla...")
	value := scanInput()

	fmt.Println(value)

}
