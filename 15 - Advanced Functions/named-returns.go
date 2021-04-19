package main

import "fmt"

// Function with named return
func mathCalcs(n1, n2 int8) (sum int8, sub int8) {
	sum = n1 + n2
	sub = n1 - n2

	return
}

func main() {
	fmt.Println("Named returns")

	sum, sub := mathCalcs(6, 2)

	fmt.Println("SUM", sum)
	fmt.Println("SUB", sub)
}
