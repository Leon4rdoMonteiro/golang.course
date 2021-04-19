package main

import "fmt"

// Variance function - accepts none or many non specified params
func sums(numbers ...int) int {
	total := 0

	// Numbers is a slice
	for _, number := range numbers {
		total += number
	}

	return total
}

// Only a variance param per function
// Using static param and variance params with different types
func concat(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println("Variance functions")

	sumTotal := sums(1, 2, 3, 4, 5)
	fmt.Println("SUM TOTAL", sumTotal)

	fmt.Println()

	concat("Copy the following number:", 1, 3, 5, 7, 9)
}
