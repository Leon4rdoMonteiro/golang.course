package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays & Slices")
	fmt.Println()

	// ARRAYS

	// Most common assignments
	var fruitsArray [3]string
	fruitsArray[0] = "Maça"
	fruitsArray[1] = "Laranja"
	fruitsArray[2] = "Banana"
	fmt.Println(fruitsArray)

	// Assignment in the declaration
	colorsArray := [5]string{"Black", "White", "Gray", "Green", "Purple"}
	fmt.Println(colorsArray)

	// Declare array size according to elements amount
	dinamicDeclarationArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println(dinamicDeclarationArray)

	// SLICES - Most dinamic and most used than arrays
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Adds a new item on slice
	slice = append(slice, 9)
	fmt.Println(slice)

	// Gets a slice from an array
	sliceFromArray := colorsArray[0:4] // Uses a pointer to assign values
	fmt.Println(sliceFromArray)
	fmt.Println()

	fmt.Println("ARRAY TYPE", reflect.TypeOf(fruitsArray))
	fmt.Println("SLICE TYPE", reflect.TypeOf(slice))

	// Internal Arrays
	slice1 := make([]float32, 10, 15)
	// An interna array is created  with 15 positions
	// After, an slice is generated from this array with 10 positions
	// If capacity is exceed, a the internal array capacity will double in size
	// We always will have capacity using slices
	fmt.Println(slice1)
	fmt.Println("Slice Length", len(slice1))
	fmt.Println("Slice Capacity", cap(slice1))
}
