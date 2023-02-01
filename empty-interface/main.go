package main

import "fmt"

type Names []interface {
}

func (names *Names) load() {
	*names = Names{"Lucas", "Jean", "Davi", 0}
}

func (names Names) print() {
	fmt.Println(names)
}

func main() {
	names := Names{}
	names.load()
	names.print()
}
