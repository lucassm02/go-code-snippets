package main

import (
	"fmt"
	"go-code-snippets/visibility/helper"
)

func main() {

	fmt.Println(helper.VisibleValue)
	helper.FuncVisible()
	// println(helper.invisibleValue) not available

}
