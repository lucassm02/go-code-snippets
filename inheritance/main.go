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

type SuperCar struct {
	Car
	Name  string
	Power string
}

func (car SuperCar) info() string {
	return fmt.Sprintf("\nName: %s\nYear: %d\nColor: %s\nPower: %s\n", car.Name, car.Year, car.Color, car.Power) // format string
}

func main() {

	car := Car{"Corola", 2017, "Black"}
	superCarOne := SuperCar{Car: car, Power: "Fly"}
	// superCarTwo := SuperCar{car, "Fly"}

	fmt.Println(superCarOne)
	fmt.Println(superCarOne.Name)

	fmt.Print(superCarOne.info())

}
