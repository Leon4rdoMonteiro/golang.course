package main

import "fmt"

func main() {
	fmt.Println("Anonymous functions")

	fmt.Println()

	result := func(person string, code int) string {
		return fmt.Sprintf("Hello World by %s! Code: %d.", person, code)
	}("Leonardo", 8)

	fmt.Println(result)
}
