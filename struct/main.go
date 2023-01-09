package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (car Car) info() string {

	return fmt.Sprintf("\nName: %s\nYear: %d\nColor: %s\n", car.Name, car.Year, car.Color) // format string

}

func main() {

	car := Car{"Corola", 2017, "Black"}

	fmt.Println(car)
	fmt.Println(car.Name)

	fmt.Print(car.info())

}
