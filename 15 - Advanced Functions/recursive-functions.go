package main

import "fmt"

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1)
}

func main() {
	fmt.Println("Recursive functions")
	fmt.Println()

	pos := uint(7)

	fmt.Println("Fibonacci sequence:")

	for i := uint(0); i <= pos; i++ {
		fmt.Println(fibonacci(i))
	}
}
