// Variables and Constants
package main

import "fmt"

func main() {
	// Declare a variable with explicit typing
	var explicitVar string = "ExplicitVar"

	// Declare many explicit typing variables
	var (
		explicitVar1 string = "ExplicitVar 1"
		explicitVar2 string = "ExplicitVar 2"
	)

	fmt.Println(explicitVar, explicitVar1, explicitVar2)

	// Declare a variables with implicit typing - by its value
	implicitVar := "ImplicitVar"

	// Declare many variables with implicit typing - by its value
	implicitVar1, implicitVar2 := "ImplicitVar1", "ImplicitVar2"
	fmt.Println(implicitVar, implicitVar1, implicitVar2)

	// Declare constants - We can use the same variables rules to create explicit or implicit constants
	const explicitConst string = "ExplicitConst"

	const (
		explicitConst1 string = "ExplicitConst1"
		explicitConst2 string = "ExplicitConst2"
	)

	fmt.Println(explicitConst, explicitConst1, explicitConst2)

	implicitConst := "ImplicitConst"

	implicitConst1, implicitConst2 := "ImplicitConst1", "ImplicitConst2"

	fmt.Println(implicitConst, implicitConst1, implicitConst2)

	// Swap variables values
	var x int = 1
	y := 2

	x, y = y, x

	fmt.Println(x, y)
}
