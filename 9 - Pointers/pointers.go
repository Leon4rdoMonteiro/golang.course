package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variable1 int8 = 10
	var variable2 int8 = variable1

	variable1++

	fmt.Println(variable1, variable2)

	// POINTERS ARE A REFERENCE OF MEMORY ADDRESS
	var variable3 int8 = 10
	var pointer *int8

	pointer = &variable3 // &variable = Gets the memory address of variable3 value
	fmt.Println(variable3, *pointer)

	variable3++
	fmt.Println(variable3, *pointer) // *variable - Makes an dereference - gets the memory address value

	// To assign a memory address of a pointer value to a variable
	variable4 := *pointer
	fmt.Println(variable4, *pointer)
}
