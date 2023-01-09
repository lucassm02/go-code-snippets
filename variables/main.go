package main

import "fmt"

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'w'
	f := "Name"

	// https://pkg.go.dev/fmt
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

}
