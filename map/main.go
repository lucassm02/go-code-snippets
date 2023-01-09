package main

import "fmt"

func main() {

	newMap := make(map[string]string)

	newMap["name"] = "Lucas"
	newMap["surname"] = "Santos"

	fmt.Println(newMap)

	delete(newMap, "surname")

	fmt.Println(newMap)

	surname, surnameExists := newMap["surname"]
	name, nameExists := newMap["name"]

	fmt.Println(surnameExists, surname) // false
	fmt.Println(nameExists, name)       // true

	var newMap2 = map[string]int{}
	newMap3 := map[string]int{"age": 0}

	fmt.Println(newMap2)
	fmt.Println(newMap3)

	if _, exists := newMap["name"]; exists {
		fmt.Println("Existe!")
	}

}
