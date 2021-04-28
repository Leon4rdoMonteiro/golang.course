package main

import "fmt"

func first() {
	fmt.Println("First function")
}

func last() {
	fmt.Println("Last function")
}

func isStudentApproved(grade1, grade2 float32) bool {
	defer fmt.Print("The result has been calculated with success, see the result: ")

	fmt.Println("The is being calculated...")

	result := (grade1 + grade2) / 2

	if result >= 6 {
		return true
	}

	return false
}

func main() {
	defer last()

	first()

	fmt.Println(isStudentApproved(4, 8))
}
