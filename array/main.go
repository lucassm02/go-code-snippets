package main

import "fmt"

func main() {

	var hello [5]string

	fmt.Println(hello)
	fmt.Println(len(hello))

	hello[0] = "H"
	hello[1] = "e"
	hello[2] = "l"
	hello[3] = "l"
	hello[4] = "o"

	// hello[5] = " " // Error

	fmt.Println(hello)
	fmt.Println(len(hello))

	welcome := [7]string{"W", "e", "l", "c", "o", "m", "e"}

	fmt.Println(welcome)
	fmt.Println(len(welcome))

	// welcome[7] = ""  // Error
}
