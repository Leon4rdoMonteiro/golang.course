package main

import "fmt"

func invertSignal(number *int) {
	*number *= -1
}

func main() {
	number := 20

	fmt.Println(number)

	invertSignal(&number)

	fmt.Println(number)
}
