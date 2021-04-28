package main

import "fmt"

// Recover the execution post a panic
func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution recovered with success!")
	}
}

func isApprovedStudent(grade1, grade2 float32) bool {
	defer recoverExecution()

	result := (grade1 + grade2) / 2

	if result > 6 {
		return true
	}

	if result < 6 {
		return false
	}

	// Stops the program execution
	panic("The result is 6!")
}

func main() {
	fmt.Println(isApprovedStudent(6, 6))
}
