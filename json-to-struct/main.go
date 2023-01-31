package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string `json:"car_color"`
}

func main() {

	var car Car

	carJson := []byte(`{"name": "Corola", "year": 2017, "car_color": "blue"}`)

	json.Unmarshal(carJson, &car)

	fmt.Println(car)

}
