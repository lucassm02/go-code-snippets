package helper

import "fmt"

const invisibleValue string = "Invisible"
const VisibleValue string = "Visible"

func FuncVisible() {
	fmt.Println(VisibleValue)
}

func funcInvisible() {
	fmt.Println(invisibleValue)
}
