package main

import "fmt"

func main() {
	// ARITHMETIC OPERATORS
	var sum = 1 + 2
	var sub = 2 - 1
	var div = 10 / 5
	var mult = 5 * 2
	var divRest = 10 % 2
	fmt.Println(sum, sub, div, mult, divRest)

	// We cannot work with diferent types
	var number1 int8 = 10
	var number2 int16 = 25

	// soma := number1 + number2
	fmt.Println(number1, number2)

	// ATTRIBUTION OPERATORS
	var variable1 string = "Variable1"
	variable2 := "Variable2"

	fmt.Println(variable1, variable2)

	// RELATIONAL OPERATORS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LOGICAL OPERATORS
	trueValue, falseValue := true, false
	fmt.Println(trueValue && falseValue)
	fmt.Println(trueValue || falseValue)
	fmt.Println(!trueValue || !falseValue)

	// UNARY OPERATORS
	number := 10
	number++
	number += 15
	number--
	number *= 2
	number /= 2
	fmt.Println(number)

	// TERNARY OPERATOR NOT EXISTS ON GOLANG
}
