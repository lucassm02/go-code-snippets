package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"car_name"`
	Year  int    `json:"-"`
	color string
}

func main() {

	car := Car{"Corola", 2017, "Black"}

	result, _ := json.Marshal(car)

	fmt.Println(string(result))

}
