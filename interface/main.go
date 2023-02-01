package main

import "fmt"

type Vehicle interface {
	start() string
}

type Car struct {
	Name string
}

func (car Car) start() string {
	return "The car" + car.Name + "has been started!"
}

func startVehicle(vehicle Vehicle) {
	fmt.Println(vehicle.start())
}

func main() {
	car := Car{"Gol"}
	startVehicle(car)
}
