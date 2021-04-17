package main

import "fmt"

// Simple function
func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Functions with many returns
func MathCalc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

// Anonymous function
var f = func(txt string) string {
	return txt
}

func main() {
	simpleSumResult := sum(5, 6)
	fmt.Println(simpleSumResult)

	anonymousFnResult := f("Anonymous function")
	fmt.Println(anonymousFnResult)

	sumResult, subResult := MathCalc(10, 5)
	fmt.Println(sumResult, subResult)

	// Ignore a function return value
	resultSum, _ := MathCalc(12, 6)
	fmt.Println(resultSum)
}
