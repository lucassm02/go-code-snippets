package main

import "fmt"

func main() {

	// slice := []int{}
	slice := make([]int, 0)

	slice2 := make([]int, 0, 5)

	fmt.Println("Slice 1: ", slice)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 1 len: ", len(slice))
	fmt.Println("Slice 1 cap: ", cap(slice))
	fmt.Println("Slice 2 len: ", len(slice2))
	fmt.Println("Slice 2 cap: ", cap(slice2)) // Capacity 5

	slice = append(slice, 1, 2, 3)
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6) // doubles the capacity when exceeding the set capacity (generates a new array and set the data from the old array into the new array)

	fmt.Print("\n")

	fmt.Println("Slice 1: ", slice)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 1 len: ", len(slice))
	fmt.Println("Slice 1 cap: ", cap(slice))
	fmt.Println("Slice 2 len: ", len(slice2))
	fmt.Println("Slice 2 cap: ", cap(slice2)) // Capacity 10

	fmt.Print("\n")

	slice3 := slice // Reference memory space

	fmt.Println("Slice 1: ", slice)
	fmt.Println("Slice 1: ", slice)
	fmt.Println("Slice 3: ", slice3)
	fmt.Println("Slice 3: ", slice3)

	slice[0] = 2 // modify value in slice memory

	fmt.Println("Slice 1: ", slice)
	fmt.Println("Slice 3: ", slice3)

	fmt.Println("Slice 1 len: ", len(slice))
	fmt.Println("Slice 1 cap: ", cap(slice))
	fmt.Println("Slice 3 len: ", len(slice3))
	fmt.Println("Slice 3 cap: ", cap(slice3))

	fmt.Print("\n")

	slice4 := slice2 // Reference memory space

	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 4: ", slice4)

	slice2 = append(slice2, 7, 8, 9, 10, 11) // when overflowing capacity a new array is generated and stored in a new memory location

	slice2[10] = 1 // now slice4 is no longer affected

	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 4: ", slice4)

	fmt.Println("Slice 2 len: ", len(slice2))
	fmt.Println("Slice 2 cap: ", cap(slice2))
	fmt.Println("Slice 4 len: ", len(slice4))
	fmt.Println("Slice 4 cap: ", cap(slice4))

	slice5 := []string{"Hello", "Lucas", "Santos"}

	fmt.Print("\n")
	fmt.Println(slice5[0])
	fmt.Println(slice5[:2])
	fmt.Println(slice5[2:3])
	fmt.Println(slice5[1:])

}
