package main

import "fmt"

var initVar int

func main() {
	fmt.Println("Executing main function")
	fmt.Println("InitVar: ", initVar)
}

// Function is executed before main function
func init() {
	fmt.Println("Executing init function")
	initVar = 10
}
