package main

import "fmt"

func closure() func() {
	text := "Inside closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Inside main function"
	fmt.Println(text)

	printText := closure()

	// The function stores the text variable value of closure function
	printText()
}
