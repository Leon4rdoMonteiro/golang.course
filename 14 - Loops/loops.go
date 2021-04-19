package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	// INCREMENTING
	count := 0

	for count < 10 {
		fmt.Println("Incrementing:", count)
		time.Sleep(time.Second)
		count++
	}

	fmt.Println()

	// Using with "Loop init"

	for i := 0; i < 10; i += 2 {
		fmt.Println("Incrementing with loop init:", i)
		time.Sleep(time.Second)
	}

	fmt.Println()

	// ARRAYS AND SLICES ITERATION
	family := [3]string{"Leonardo", "Claudia", "Oziel"}

	fmt.Println("Code", "Name")
	for index, name := range family {
		fmt.Println(index, name)
	}

	fmt.Println()

	// STRINGS ITERATIONS
	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	fmt.Println()

	// MAPS ITERATION
	user := map[string]string{
		"first_name": "Leonardo",
		"last_name":  "Monteiro",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Example of infinity loop
	for {
		fmt.Println("Running an infinity loop...")
		time.Sleep(time.Second)
	}

	// Structs are not iterable
}
