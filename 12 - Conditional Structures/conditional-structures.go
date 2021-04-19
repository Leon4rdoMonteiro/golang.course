package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional Strucutures")

	number := 10

	if number > 15 {
		fmt.Println("Greather than 15")
	} else {
		fmt.Println("Less or equal to 15")
	}

	// IF Init - Initialize vars into IFs
	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Number is greather than 0")
		fmt.Println("Number", anotherNumber)

		// We can use this var into IF scope
	} else if number < -10 {
		fmt.Println("Number is less than -10")
	} else {
		fmt.Println("Number is between 0 and -10")
	}

}
